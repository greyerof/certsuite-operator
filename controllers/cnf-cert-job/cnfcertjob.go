package cnfcertjob

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/test-network-function/cnf-certsuite-operator/controllers/definitions"
)

const (
	// Be careful when changing this SA name.
	// 1. It must match the flag --extra-service-accounts in "make bundle".
	// 2. The prefix is "cnf-certsuite-". It should match the field namePrefix field in config/default/kustomization.yaml.
	clusterAccessServiceAccountName = "cnf-certsuite-cluster-access"
)

func New(options ...func(*corev1.Pod) error) (*corev1.Pod, error) {
	jobPod := newInitialJobPod()

	for _, o := range options {
		err := o(jobPod)
		if err != nil {
			return nil, err
		}
	}
	return jobPod, nil
}

//nolint:funlen
func newInitialJobPod() *corev1.Pod {
	return &corev1.Pod{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{},
		Spec: corev1.PodSpec{
			ServiceAccountName: clusterAccessServiceAccountName,
			RestartPolicy:      "Never",
			Containers: []corev1.Container{
				{
					Name: definitions.CnfCertSuiteSidecarContainerName,
					Env: []corev1.EnvVar{
						{
							Name: "MY_POD_NAME",
							ValueFrom: &corev1.EnvVarSource{
								FieldRef: &corev1.ObjectFieldSelector{
									FieldPath: "metadata.name",
								},
							},
						},
						{
							Name: "MY_POD_NAMESPACE",
							ValueFrom: &corev1.EnvVarSource{
								FieldRef: &corev1.ObjectFieldSelector{
									FieldPath: "metadata.namespace",
								},
							},
						},
						{
							Name:  definitions.SideCarResultsFolderEnvVar,
							Value: definitions.CnfCertSuiteResultsFolder,
						},
					},
					ImagePullPolicy: "IfNotPresent",
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "cnf-certsuite-output",
							ReadOnly:  true,
							MountPath: definitions.CnfCertSuiteResultsFolder,
						},
					},
				},
				{
					Name:    definitions.CnfCertSuiteContainerName,
					Image:   "quay.io/testnetworkfunction/cnf-certification-test:unstable",
					Command: []string{"./run-cnf-suites.sh"},
					Args:    []string{"-o", definitions.CnfCertSuiteResultsFolder},
					Env: []corev1.EnvVar{
						{
							Name:  "PFLT_DOCKERCONFIG",
							Value: definitions.PreflightDockerConfigFilePath,
						},
						{
							Name:  "TNF_CONFIGURATION_PATH",
							Value: definitions.CnfCertSuiteConfigFilePath,
						},
						{
							Name:  "TNF_NON_INTRUSIVE_ONLY",
							Value: "true",
						},
					},
					ImagePullPolicy: "Always",
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "cnf-certsuite-output",
							MountPath: definitions.CnfCertSuiteResultsFolder,
						},
						{
							Name:      "cnf-certsuite-config",
							ReadOnly:  true,
							MountPath: definitions.CnfCnfCertSuiteConfigFolder,
						},
					},
				},
			},
			Volumes: []corev1.Volume{
				{
					Name: "cnf-certsuite-output",
					VolumeSource: corev1.VolumeSource{
						EmptyDir: &corev1.EmptyDirVolumeSource{},
					},
				},
			},
		},
		Status: corev1.PodStatus{},
	}
}

func WithPodName(podName string) func(*corev1.Pod) error {
	return func(p *corev1.Pod) error {
		p.ObjectMeta.Name = podName
		return nil
	}
}

func WithNamespace(namespace string) func(*corev1.Pod) error {
	return func(p *corev1.Pod) error {
		p.ObjectMeta.Namespace = namespace
		return nil
	}
}

func WithCertSuiteConfigRunName(certSuiteConfigRunName string) func(*corev1.Pod) error {
	return func(p *corev1.Pod) error {
		envVar := corev1.EnvVar{Name: "RUN_CR_NAME", Value: certSuiteConfigRunName}
		sideCarContainer := getSideCarAppContainer(p)
		if sideCarContainer == nil {
			return fmt.Errorf("side Car app Container is not found in pod %s", p.Name)
		}
		sideCarContainer.Env = append(sideCarContainer.Env, envVar)
		return nil
	}
}

func WithLabelsFilter(labelsFilter string) func(*corev1.Pod) error {
	return func(p *corev1.Pod) error {
		cnfCertSuiteContainer := getCnfCertSuiteContainer(p)
		if cnfCertSuiteContainer == nil {
			return fmt.Errorf("cnf cert suite Container is not found in pod %s", p.Name)
		}
		cnfCertSuiteContainer.Args = append(cnfCertSuiteContainer.Args, "-l", labelsFilter)
		return nil
	}
}

func WithLogLevel(loglevel string) func(*corev1.Pod) error {
	return func(p *corev1.Pod) error {
		envVar := corev1.EnvVar{Name: "TNF_LOG_LEVEL", Value: loglevel}
		cnfCertSuiteContainer := getCnfCertSuiteContainer(p)
		if cnfCertSuiteContainer == nil {
			return fmt.Errorf("cnf cert suite Container is not found in pod %s", p.Name)
		}
		cnfCertSuiteContainer.Env = append(cnfCertSuiteContainer.Env, envVar)
		return nil
	}
}

func WithTimeOut(timeout string) func(*corev1.Pod) error {
	return func(p *corev1.Pod) error {
		envVar := corev1.EnvVar{Name: "TIMEOUT", Value: timeout}
		cnfCertSuiteContainer := getCnfCertSuiteContainer(p)
		if cnfCertSuiteContainer == nil {
			return fmt.Errorf("cnf cert suite Container is not found in pod %s", p.Name)
		}
		cnfCertSuiteContainer.Env = append(cnfCertSuiteContainer.Env, envVar)
		return nil
	}
}

func WithConfigMap(configMapName string) func(*corev1.Pod) error {
	return func(p *corev1.Pod) error {
		Volume := corev1.Volume{
			Name: "cnf-certsuite-config",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: configMapName,
					},
				},
			},
		}
		p.Spec.Volumes = append(p.Spec.Volumes, Volume)
		return nil
	}
}

func WithPreflightSecret(preflightSecretName *string) func(*corev1.Pod) error {
	return func(p *corev1.Pod) error {
		if preflightSecretName == nil {
			return nil
		}

		Volume := corev1.Volume{
			Name: "cnf-certsuite-preflight-dockerconfig",
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName: *preflightSecretName,
				},
			},
		}
		p.Spec.Volumes = append(p.Spec.Volumes, Volume)

		cnfCertCuiteContainer := getCnfCertSuiteContainer(p)
		volumeMount := corev1.VolumeMount{
			Name:      "cnf-certsuite-preflight-dockerconfig",
			ReadOnly:  true,
			MountPath: definitions.CnfPreflightConfigFolder,
		}
		cnfCertCuiteContainer.VolumeMounts = append(cnfCertCuiteContainer.VolumeMounts, volumeMount)
		return nil
	}
}

func WithSideCarApp(sideCarAppImage string) func(*corev1.Pod) error {
	return func(p *corev1.Pod) error {
		sideCarContainer := getSideCarAppContainer(p)
		if sideCarContainer == nil {
			return fmt.Errorf("side Car app Container is not found in pod %s", p.Name)
		}
		sideCarContainer.Image = sideCarAppImage
		return nil
	}
}

func WithEnableDataCollection(enableDataCollection string) func(*corev1.Pod) error {
	return func(p *corev1.Pod) error {
		envVar := corev1.EnvVar{Name: "TNF_ENABLE_DATA_COLLECTION", Value: enableDataCollection}
		cnfCertSuiteContainer := getCnfCertSuiteContainer(p)
		if cnfCertSuiteContainer == nil {
			return fmt.Errorf("cnf cert suite Container is not found in pod %s", p.Name)
		}
		cnfCertSuiteContainer.Env = append(cnfCertSuiteContainer.Env, envVar)
		return nil
	}
}

func getSideCarAppContainer(p *corev1.Pod) *corev1.Container {
	for i := range p.Spec.Containers {
		if p.Spec.Containers[i].Name == definitions.CnfCertSuiteSidecarContainerName {
			return &p.Spec.Containers[i]
		}
	}
	return nil
}

func getCnfCertSuiteContainer(p *corev1.Pod) *corev1.Container {
	for i := range p.Spec.Containers {
		if p.Spec.Containers[i].Name == definitions.CnfCertSuiteContainerName {
			return &p.Spec.Containers[i]
		}
	}
	return nil
}

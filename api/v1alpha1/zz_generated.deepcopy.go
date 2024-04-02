//go:build !ignore_autogenerated

/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnfCertificationSuiteReport) DeepCopyInto(out *CnfCertificationSuiteReport) {
	*out = *in
	in.CnfTargets.DeepCopyInto(&out.CnfTargets)
	out.Summary = in.Summary
	if in.Results != nil {
		in, out := &in.Results, &out.Results
		*out = make([]TestCaseResult, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnfCertificationSuiteReport.
func (in *CnfCertificationSuiteReport) DeepCopy() *CnfCertificationSuiteReport {
	if in == nil {
		return nil
	}
	out := new(CnfCertificationSuiteReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnfCertificationSuiteReportStatusSummary) DeepCopyInto(out *CnfCertificationSuiteReportStatusSummary) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnfCertificationSuiteReportStatusSummary.
func (in *CnfCertificationSuiteReportStatusSummary) DeepCopy() *CnfCertificationSuiteReportStatusSummary {
	if in == nil {
		return nil
	}
	out := new(CnfCertificationSuiteReportStatusSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnfCertificationSuiteRun) DeepCopyInto(out *CnfCertificationSuiteRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnfCertificationSuiteRun.
func (in *CnfCertificationSuiteRun) DeepCopy() *CnfCertificationSuiteRun {
	if in == nil {
		return nil
	}
	out := new(CnfCertificationSuiteRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CnfCertificationSuiteRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnfCertificationSuiteRunList) DeepCopyInto(out *CnfCertificationSuiteRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CnfCertificationSuiteRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnfCertificationSuiteRunList.
func (in *CnfCertificationSuiteRunList) DeepCopy() *CnfCertificationSuiteRunList {
	if in == nil {
		return nil
	}
	out := new(CnfCertificationSuiteRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CnfCertificationSuiteRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnfCertificationSuiteRunSpec) DeepCopyInto(out *CnfCertificationSuiteRunSpec) {
	*out = *in
	if in.PreflightSecretName != nil {
		in, out := &in.PreflightSecretName, &out.PreflightSecretName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnfCertificationSuiteRunSpec.
func (in *CnfCertificationSuiteRunSpec) DeepCopy() *CnfCertificationSuiteRunSpec {
	if in == nil {
		return nil
	}
	out := new(CnfCertificationSuiteRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnfCertificationSuiteRunStatus) DeepCopyInto(out *CnfCertificationSuiteRunStatus) {
	*out = *in
	if in.CnfCertSuitePodName != nil {
		in, out := &in.CnfCertSuitePodName, &out.CnfCertSuitePodName
		*out = new(string)
		**out = **in
	}
	if in.Report != nil {
		in, out := &in.Report, &out.Report
		*out = new(CnfCertificationSuiteReport)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnfCertificationSuiteRunStatus.
func (in *CnfCertificationSuiteRunStatus) DeepCopy() *CnfCertificationSuiteRunStatus {
	if in == nil {
		return nil
	}
	out := new(CnfCertificationSuiteRunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnfPod) DeepCopyInto(out *CnfPod) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnfPod.
func (in *CnfPod) DeepCopy() *CnfPod {
	if in == nil {
		return nil
	}
	out := new(CnfPod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnfResource) DeepCopyInto(out *CnfResource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnfResource.
func (in *CnfResource) DeepCopy() *CnfResource {
	if in == nil {
		return nil
	}
	out := new(CnfResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnfTargets) DeepCopyInto(out *CnfTargets) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]CnfPod, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Deployments != nil {
		in, out := &in.Deployments, &out.Deployments
		*out = make([]CnfResource, len(*in))
		copy(*out, *in)
	}
	if in.StatefulSets != nil {
		in, out := &in.StatefulSets, &out.StatefulSets
		*out = make([]CnfResource, len(*in))
		copy(*out, *in)
	}
	if in.Csvs != nil {
		in, out := &in.Csvs, &out.Csvs
		*out = make([]CnfResource, len(*in))
		copy(*out, *in)
	}
	if in.Crds != nil {
		in, out := &in.Crds, &out.Crds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]CnfResource, len(*in))
		copy(*out, *in)
	}
	if in.HelmChartReleases != nil {
		in, out := &in.HelmChartReleases, &out.HelmChartReleases
		*out = make([]CnfResource, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnfTargets.
func (in *CnfTargets) DeepCopy() *CnfTargets {
	if in == nil {
		return nil
	}
	out := new(CnfTargets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in TargetResource) DeepCopyInto(out *TargetResource) {
	{
		in := &in
		*out = make(TargetResource, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetResource.
func (in TargetResource) DeepCopy() TargetResource {
	if in == nil {
		return nil
	}
	out := new(TargetResource)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetResources) DeepCopyInto(out *TargetResources) {
	*out = *in
	if in.Compliant != nil {
		in, out := &in.Compliant, &out.Compliant
		*out = make([]TargetResource, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(TargetResource, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	if in.NonCompliant != nil {
		in, out := &in.NonCompliant, &out.NonCompliant
		*out = make([]TargetResource, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(TargetResource, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetResources.
func (in *TargetResources) DeepCopy() *TargetResources {
	if in == nil {
		return nil
	}
	out := new(TargetResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestCaseResult) DeepCopyInto(out *TestCaseResult) {
	*out = *in
	if in.TargetResources != nil {
		in, out := &in.TargetResources, &out.TargetResources
		*out = new(TargetResources)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestCaseResult.
func (in *TestCaseResult) DeepCopy() *TestCaseResult {
	if in == nil {
		return nil
	}
	out := new(TestCaseResult)
	in.DeepCopyInto(out)
	return out
}

// +build !ignore_autogenerated

/*

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
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicMetadata) DeepCopyInto(out *BasicMetadata) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicMetadata.
func (in *BasicMetadata) DeepCopy() *BasicMetadata {
	if in == nil {
		return nil
	}
	out := new(BasicMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FakePodTemplateSpec) DeepCopyInto(out *FakePodTemplateSpec) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FakePodTemplateSpec.
func (in *FakePodTemplateSpec) DeepCopy() *FakePodTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(FakePodTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinnedDeployment) DeepCopyInto(out *PinnedDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinnedDeployment.
func (in *PinnedDeployment) DeepCopy() *PinnedDeployment {
	if in == nil {
		return nil
	}
	out := new(PinnedDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PinnedDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinnedDeploymentList) DeepCopyInto(out *PinnedDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PinnedDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinnedDeploymentList.
func (in *PinnedDeploymentList) DeepCopy() *PinnedDeploymentList {
	if in == nil {
		return nil
	}
	out := new(PinnedDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PinnedDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinnedDeploymentSpec) DeepCopyInto(out *PinnedDeploymentSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Templates.DeepCopyInto(&out.Templates)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinnedDeploymentSpec.
func (in *PinnedDeploymentSpec) DeepCopy() *PinnedDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(PinnedDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinnedDeploymentStatus) DeepCopyInto(out *PinnedDeploymentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinnedDeploymentStatus.
func (in *PinnedDeploymentStatus) DeepCopy() *PinnedDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(PinnedDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreviousNextPodTemplateSpecPair) DeepCopyInto(out *PreviousNextPodTemplateSpecPair) {
	*out = *in
	in.Previous.DeepCopyInto(&out.Previous)
	in.Next.DeepCopyInto(&out.Next)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreviousNextPodTemplateSpecPair.
func (in *PreviousNextPodTemplateSpecPair) DeepCopy() *PreviousNextPodTemplateSpecPair {
	if in == nil {
		return nil
	}
	out := new(PreviousNextPodTemplateSpecPair)
	in.DeepCopyInto(out)
	return out
}
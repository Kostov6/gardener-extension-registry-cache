//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bastion) DeepCopyInto(out *Bastion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bastion.
func (in *Bastion) DeepCopy() *Bastion {
	if in == nil {
		return nil
	}
	out := new(Bastion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Bastion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BastionIngressPolicy) DeepCopyInto(out *BastionIngressPolicy) {
	*out = *in
	in.IPBlock.DeepCopyInto(&out.IPBlock)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BastionIngressPolicy.
func (in *BastionIngressPolicy) DeepCopy() *BastionIngressPolicy {
	if in == nil {
		return nil
	}
	out := new(BastionIngressPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BastionList) DeepCopyInto(out *BastionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Bastion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BastionList.
func (in *BastionList) DeepCopy() *BastionList {
	if in == nil {
		return nil
	}
	out := new(BastionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BastionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BastionSpec) DeepCopyInto(out *BastionSpec) {
	*out = *in
	out.ShootRef = in.ShootRef
	if in.SeedName != nil {
		in, out := &in.SeedName, &out.SeedName
		*out = new(string)
		**out = **in
	}
	if in.ProviderType != nil {
		in, out := &in.ProviderType, &out.ProviderType
		*out = new(string)
		**out = **in
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]BastionIngressPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BastionSpec.
func (in *BastionSpec) DeepCopy() *BastionSpec {
	if in == nil {
		return nil
	}
	out := new(BastionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BastionStatus) DeepCopyInto(out *BastionStatus) {
	*out = *in
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(v1.LoadBalancerIngress)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1beta1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastHeartbeatTimestamp != nil {
		in, out := &in.LastHeartbeatTimestamp, &out.LastHeartbeatTimestamp
		*out = (*in).DeepCopy()
	}
	if in.ExpirationTimestamp != nil {
		in, out := &in.ExpirationTimestamp, &out.ExpirationTimestamp
		*out = (*in).DeepCopy()
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BastionStatus.
func (in *BastionStatus) DeepCopy() *BastionStatus {
	if in == nil {
		return nil
	}
	out := new(BastionStatus)
	in.DeepCopyInto(out)
	return out
}

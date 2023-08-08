//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAttestor) DeepCopyInto(out *NodeAttestor) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAttestor.
func (in *NodeAttestor) DeepCopy() *NodeAttestor {
	if in == nil {
		return nil
	}
	out := new(NodeAttestor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpireAgent) DeepCopyInto(out *SpireAgent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpireAgent.
func (in *SpireAgent) DeepCopy() *SpireAgent {
	if in == nil {
		return nil
	}
	out := new(SpireAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpireAgent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpireAgentList) DeepCopyInto(out *SpireAgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpireAgent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpireAgentList.
func (in *SpireAgentList) DeepCopy() *SpireAgentList {
	if in == nil {
		return nil
	}
	out := new(SpireAgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpireAgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpireAgentSpec) DeepCopyInto(out *SpireAgentSpec) {
	*out = *in
	out.NodeAttestor = in.NodeAttestor
	if in.WorkloadAttestors != nil {
		in, out := &in.WorkloadAttestors, &out.WorkloadAttestors
		*out = make([]WorkloadAttestor, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpireAgentSpec.
func (in *SpireAgentSpec) DeepCopy() *SpireAgentSpec {
	if in == nil {
		return nil
	}
	out := new(SpireAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpireAgentStatus) DeepCopyInto(out *SpireAgentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpireAgentStatus.
func (in *SpireAgentStatus) DeepCopy() *SpireAgentStatus {
	if in == nil {
		return nil
	}
	out := new(SpireAgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpireServer) DeepCopyInto(out *SpireServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpireServer.
func (in *SpireServer) DeepCopy() *SpireServer {
	if in == nil {
		return nil
	}
	out := new(SpireServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpireServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpireServerList) DeepCopyInto(out *SpireServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpireServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpireServerList.
func (in *SpireServerList) DeepCopy() *SpireServerList {
	if in == nil {
		return nil
	}
	out := new(SpireServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpireServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpireServerSpec) DeepCopyInto(out *SpireServerSpec) {
	*out = *in
	if in.NodeAttestors != nil {
		in, out := &in.NodeAttestors, &out.NodeAttestors
		*out = make([]NodeAttestor, len(*in))
		copy(*out, *in)
	}
	if in.CertAuthorities != nil {
		in, out := &in.CertAuthorities, &out.CertAuthorities
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpireServerSpec.
func (in *SpireServerSpec) DeepCopy() *SpireServerSpec {
	if in == nil {
		return nil
	}
	out := new(SpireServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpireServerStatus) DeepCopyInto(out *SpireServerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpireServerStatus.
func (in *SpireServerStatus) DeepCopy() *SpireServerStatus {
	if in == nil {
		return nil
	}
	out := new(SpireServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadAttestor) DeepCopyInto(out *WorkloadAttestor) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadAttestor.
func (in *WorkloadAttestor) DeepCopy() *WorkloadAttestor {
	if in == nil {
		return nil
	}
	out := new(WorkloadAttestor)
	in.DeepCopyInto(out)
	return out
}

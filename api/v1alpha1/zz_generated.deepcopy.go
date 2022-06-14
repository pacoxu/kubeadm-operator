//go:build !ignore_autogenerated
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
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubeadm/operator/errors"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommandDescriptor) DeepCopyInto(out *CommandDescriptor) {
	*out = *in
	if in.KubeadmRenewCertificates != nil {
		in, out := &in.KubeadmRenewCertificates, &out.KubeadmRenewCertificates
		*out = new(KubeadmRenewCertsCommandSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeadmUpgradeApply != nil {
		in, out := &in.KubeadmUpgradeApply, &out.KubeadmUpgradeApply
		*out = new(KubeadmUpgradeApplyCommandSpec)
		**out = **in
	}
	if in.KubeadmUpgradeNode != nil {
		in, out := &in.KubeadmUpgradeNode, &out.KubeadmUpgradeNode
		*out = new(KubeadmUpgradeNodeCommandSpec)
		**out = **in
	}
	if in.Preflight != nil {
		in, out := &in.Preflight, &out.Preflight
		*out = new(PreflightCommandSpec)
		**out = **in
	}
	if in.UpgradeKubeadm != nil {
		in, out := &in.UpgradeKubeadm, &out.UpgradeKubeadm
		*out = new(UpgradeKubeadmCommandSpec)
		**out = **in
	}
	if in.UpgradeKubeletAndKubeactl != nil {
		in, out := &in.UpgradeKubeletAndKubeactl, &out.UpgradeKubeletAndKubeactl
		*out = new(UpgradeKubeletAndKubeactlCommandSpec)
		**out = **in
	}
	if in.KubectlDrain != nil {
		in, out := &in.KubectlDrain, &out.KubectlDrain
		*out = new(KubectlDrainCommandSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.KubectlUncordon != nil {
		in, out := &in.KubectlUncordon, &out.KubectlUncordon
		*out = new(KubectlUncordonCommandSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeadmUpgradeKubeProxy != nil {
		in, out := &in.KubeadmUpgradeKubeProxy, &out.KubeadmUpgradeKubeProxy
		*out = new(KubeadmUpgradeKubeProxySpec)
		**out = **in
	}
	if in.Pass != nil {
		in, out := &in.Pass, &out.Pass
		*out = new(PassCommandSpec)
		**out = **in
	}
	if in.Fail != nil {
		in, out := &in.Fail, &out.Fail
		*out = new(FailCommandSpec)
		**out = **in
	}
	if in.Wait != nil {
		in, out := &in.Wait, &out.Wait
		*out = new(WaitCommandSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommandDescriptor.
func (in *CommandDescriptor) DeepCopy() *CommandDescriptor {
	if in == nil {
		return nil
	}
	out := new(CommandDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomOperationSpec) DeepCopyInto(out *CustomOperationSpec) {
	*out = *in
	if in.Workflow != nil {
		in, out := &in.Workflow, &out.Workflow
		*out = make([]RuntimeTaskGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomOperationSpec.
func (in *CustomOperationSpec) DeepCopy() *CustomOperationSpec {
	if in == nil {
		return nil
	}
	out := new(CustomOperationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailCommandSpec) DeepCopyInto(out *FailCommandSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailCommandSpec.
func (in *FailCommandSpec) DeepCopy() *FailCommandSpec {
	if in == nil {
		return nil
	}
	out := new(FailCommandSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmRenewCertsCommandSpec) DeepCopyInto(out *KubeadmRenewCertsCommandSpec) {
	*out = *in
	if in.Commands != nil {
		in, out := &in.Commands, &out.Commands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmRenewCertsCommandSpec.
func (in *KubeadmRenewCertsCommandSpec) DeepCopy() *KubeadmRenewCertsCommandSpec {
	if in == nil {
		return nil
	}
	out := new(KubeadmRenewCertsCommandSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmUpgradeApplyCommandSpec) DeepCopyInto(out *KubeadmUpgradeApplyCommandSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmUpgradeApplyCommandSpec.
func (in *KubeadmUpgradeApplyCommandSpec) DeepCopy() *KubeadmUpgradeApplyCommandSpec {
	if in == nil {
		return nil
	}
	out := new(KubeadmUpgradeApplyCommandSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmUpgradeKubeProxySpec) DeepCopyInto(out *KubeadmUpgradeKubeProxySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmUpgradeKubeProxySpec.
func (in *KubeadmUpgradeKubeProxySpec) DeepCopy() *KubeadmUpgradeKubeProxySpec {
	if in == nil {
		return nil
	}
	out := new(KubeadmUpgradeKubeProxySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmUpgradeNodeCommandSpec) DeepCopyInto(out *KubeadmUpgradeNodeCommandSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmUpgradeNodeCommandSpec.
func (in *KubeadmUpgradeNodeCommandSpec) DeepCopy() *KubeadmUpgradeNodeCommandSpec {
	if in == nil {
		return nil
	}
	out := new(KubeadmUpgradeNodeCommandSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubectlDrainCommandSpec) DeepCopyInto(out *KubectlDrainCommandSpec) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubectlDrainCommandSpec.
func (in *KubectlDrainCommandSpec) DeepCopy() *KubectlDrainCommandSpec {
	if in == nil {
		return nil
	}
	out := new(KubectlDrainCommandSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubectlUncordonCommandSpec) DeepCopyInto(out *KubectlUncordonCommandSpec) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubectlUncordonCommandSpec.
func (in *KubectlUncordonCommandSpec) DeepCopy() *KubectlUncordonCommandSpec {
	if in == nil {
		return nil
	}
	out := new(KubectlUncordonCommandSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Operation) DeepCopyInto(out *Operation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Operation.
func (in *Operation) DeepCopy() *Operation {
	if in == nil {
		return nil
	}
	out := new(Operation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Operation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationList) DeepCopyInto(out *OperationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Operation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationList.
func (in *OperationList) DeepCopy() *OperationList {
	if in == nil {
		return nil
	}
	out := new(OperationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationSpec) DeepCopyInto(out *OperationSpec) {
	*out = *in
	in.OperatorDescriptor.DeepCopyInto(&out.OperatorDescriptor)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationSpec.
func (in *OperationSpec) DeepCopy() *OperationSpec {
	if in == nil {
		return nil
	}
	out := new(OperationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationStatus) DeepCopyInto(out *OperationStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.ErrorReason != nil {
		in, out := &in.ErrorReason, &out.ErrorReason
		*out = new(errors.OperationStatusError)
		**out = **in
	}
	if in.ErrorMessage != nil {
		in, out := &in.ErrorMessage, &out.ErrorMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationStatus.
func (in *OperationStatus) DeepCopy() *OperationStatus {
	if in == nil {
		return nil
	}
	out := new(OperationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorDescriptor) DeepCopyInto(out *OperatorDescriptor) {
	*out = *in
	if in.Upgrade != nil {
		in, out := &in.Upgrade, &out.Upgrade
		*out = new(UpgradeOperationSpec)
		**out = **in
	}
	if in.RenewCertificates != nil {
		in, out := &in.RenewCertificates, &out.RenewCertificates
		*out = new(RenewCertificatesOperationSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomOperation != nil {
		in, out := &in.CustomOperation, &out.CustomOperation
		*out = new(CustomOperationSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorDescriptor.
func (in *OperatorDescriptor) DeepCopy() *OperatorDescriptor {
	if in == nil {
		return nil
	}
	out := new(OperatorDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PassCommandSpec) DeepCopyInto(out *PassCommandSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PassCommandSpec.
func (in *PassCommandSpec) DeepCopy() *PassCommandSpec {
	if in == nil {
		return nil
	}
	out := new(PassCommandSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightCommandSpec) DeepCopyInto(out *PreflightCommandSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightCommandSpec.
func (in *PreflightCommandSpec) DeepCopy() *PreflightCommandSpec {
	if in == nil {
		return nil
	}
	out := new(PreflightCommandSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RenewCertificatesOperationSpec) DeepCopyInto(out *RenewCertificatesOperationSpec) {
	*out = *in
	if in.Commands != nil {
		in, out := &in.Commands, &out.Commands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RenewCertificatesOperationSpec.
func (in *RenewCertificatesOperationSpec) DeepCopy() *RenewCertificatesOperationSpec {
	if in == nil {
		return nil
	}
	out := new(RenewCertificatesOperationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeTask) DeepCopyInto(out *RuntimeTask) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeTask.
func (in *RuntimeTask) DeepCopy() *RuntimeTask {
	if in == nil {
		return nil
	}
	out := new(RuntimeTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RuntimeTask) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeTaskGroup) DeepCopyInto(out *RuntimeTaskGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeTaskGroup.
func (in *RuntimeTaskGroup) DeepCopy() *RuntimeTaskGroup {
	if in == nil {
		return nil
	}
	out := new(RuntimeTaskGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RuntimeTaskGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeTaskGroupList) DeepCopyInto(out *RuntimeTaskGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RuntimeTaskGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeTaskGroupList.
func (in *RuntimeTaskGroupList) DeepCopy() *RuntimeTaskGroupList {
	if in == nil {
		return nil
	}
	out := new(RuntimeTaskGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RuntimeTaskGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeTaskGroupSpec) DeepCopyInto(out *RuntimeTaskGroupSpec) {
	*out = *in
	in.NodeSelector.DeepCopyInto(&out.NodeSelector)
	in.Selector.DeepCopyInto(&out.Selector)
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeTaskGroupSpec.
func (in *RuntimeTaskGroupSpec) DeepCopy() *RuntimeTaskGroupSpec {
	if in == nil {
		return nil
	}
	out := new(RuntimeTaskGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeTaskGroupStatus) DeepCopyInto(out *RuntimeTaskGroupStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.ErrorReason != nil {
		in, out := &in.ErrorReason, &out.ErrorReason
		*out = new(errors.RuntimeTaskGroupStatusError)
		**out = **in
	}
	if in.ErrorMessage != nil {
		in, out := &in.ErrorMessage, &out.ErrorMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeTaskGroupStatus.
func (in *RuntimeTaskGroupStatus) DeepCopy() *RuntimeTaskGroupStatus {
	if in == nil {
		return nil
	}
	out := new(RuntimeTaskGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeTaskList) DeepCopyInto(out *RuntimeTaskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RuntimeTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeTaskList.
func (in *RuntimeTaskList) DeepCopy() *RuntimeTaskList {
	if in == nil {
		return nil
	}
	out := new(RuntimeTaskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RuntimeTaskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeTaskSpec) DeepCopyInto(out *RuntimeTaskSpec) {
	*out = *in
	if in.Commands != nil {
		in, out := &in.Commands, &out.Commands
		*out = make([]CommandDescriptor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeTaskSpec.
func (in *RuntimeTaskSpec) DeepCopy() *RuntimeTaskSpec {
	if in == nil {
		return nil
	}
	out := new(RuntimeTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeTaskStatus) DeepCopyInto(out *RuntimeTaskStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.ErrorReason != nil {
		in, out := &in.ErrorReason, &out.ErrorReason
		*out = new(errors.RuntimeTaskStatusError)
		**out = **in
	}
	if in.ErrorMessage != nil {
		in, out := &in.ErrorMessage, &out.ErrorMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeTaskStatus.
func (in *RuntimeTaskStatus) DeepCopy() *RuntimeTaskStatus {
	if in == nil {
		return nil
	}
	out := new(RuntimeTaskStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeTaskTemplateSpec) DeepCopyInto(out *RuntimeTaskTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeTaskTemplateSpec.
func (in *RuntimeTaskTemplateSpec) DeepCopy() *RuntimeTaskTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(RuntimeTaskTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeKubeadmCommandSpec) DeepCopyInto(out *UpgradeKubeadmCommandSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeKubeadmCommandSpec.
func (in *UpgradeKubeadmCommandSpec) DeepCopy() *UpgradeKubeadmCommandSpec {
	if in == nil {
		return nil
	}
	out := new(UpgradeKubeadmCommandSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeKubeletAndKubeactlCommandSpec) DeepCopyInto(out *UpgradeKubeletAndKubeactlCommandSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeKubeletAndKubeactlCommandSpec.
func (in *UpgradeKubeletAndKubeactlCommandSpec) DeepCopy() *UpgradeKubeletAndKubeactlCommandSpec {
	if in == nil {
		return nil
	}
	out := new(UpgradeKubeletAndKubeactlCommandSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeOperationSpec) DeepCopyInto(out *UpgradeOperationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeOperationSpec.
func (in *UpgradeOperationSpec) DeepCopy() *UpgradeOperationSpec {
	if in == nil {
		return nil
	}
	out := new(UpgradeOperationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WaitCommandSpec) DeepCopyInto(out *WaitCommandSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WaitCommandSpec.
func (in *WaitCommandSpec) DeepCopy() *WaitCommandSpec {
	if in == nil {
		return nil
	}
	out := new(WaitCommandSpec)
	in.DeepCopyInto(out)
	return out
}

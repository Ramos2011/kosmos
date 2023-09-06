//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by defaulter-gen. DO NOT EDIT.

package v1beta2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1beta2 "k8s.io/kube-scheduler/config/v1beta2"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&v1beta2.DefaultPreemptionArgs{}, func(obj interface{}) { SetObjectDefaults_DefaultPreemptionArgs(obj.(*v1beta2.DefaultPreemptionArgs)) })
	scheme.AddTypeDefaultingFunc(&v1beta2.InterPodAffinityArgs{}, func(obj interface{}) { SetObjectDefaults_InterPodAffinityArgs(obj.(*v1beta2.InterPodAffinityArgs)) })
	scheme.AddTypeDefaultingFunc(&v1beta2.KubeSchedulerConfiguration{}, func(obj interface{}) {
		SetObjectDefaults_KubeSchedulerConfiguration(obj.(*v1beta2.KubeSchedulerConfiguration))
	})
	scheme.AddTypeDefaultingFunc(&v1beta2.NodeResourcesBalancedAllocationArgs{}, func(obj interface{}) {
		SetObjectDefaults_NodeResourcesBalancedAllocationArgs(obj.(*v1beta2.NodeResourcesBalancedAllocationArgs))
	})
	scheme.AddTypeDefaultingFunc(&v1beta2.NodeResourcesFitArgs{}, func(obj interface{}) { SetObjectDefaults_NodeResourcesFitArgs(obj.(*v1beta2.NodeResourcesFitArgs)) })
	scheme.AddTypeDefaultingFunc(&v1beta2.PodTopologySpreadArgs{}, func(obj interface{}) { SetObjectDefaults_PodTopologySpreadArgs(obj.(*v1beta2.PodTopologySpreadArgs)) })
	scheme.AddTypeDefaultingFunc(&v1beta2.VolumeBindingArgs{}, func(obj interface{}) { SetObjectDefaults_VolumeBindingArgs(obj.(*v1beta2.VolumeBindingArgs)) })
	return nil
}

func SetObjectDefaults_DefaultPreemptionArgs(in *v1beta2.DefaultPreemptionArgs) {
	SetDefaults_DefaultPreemptionArgs(in)
}

func SetObjectDefaults_InterPodAffinityArgs(in *v1beta2.InterPodAffinityArgs) {
	SetDefaults_InterPodAffinityArgs(in)
}

func SetObjectDefaults_KubeSchedulerConfiguration(in *v1beta2.KubeSchedulerConfiguration) {
	SetDefaults_KubeSchedulerConfiguration(in)
}

func SetObjectDefaults_NodeResourcesBalancedAllocationArgs(in *v1beta2.NodeResourcesBalancedAllocationArgs) {
	SetDefaults_NodeResourcesBalancedAllocationArgs(in)
}

func SetObjectDefaults_NodeResourcesFitArgs(in *v1beta2.NodeResourcesFitArgs) {
	SetDefaults_NodeResourcesFitArgs(in)
}

func SetObjectDefaults_PodTopologySpreadArgs(in *v1beta2.PodTopologySpreadArgs) {
	SetDefaults_PodTopologySpreadArgs(in)
}

func SetObjectDefaults_VolumeBindingArgs(in *v1beta2.VolumeBindingArgs) {
	SetDefaults_VolumeBindingArgs(in)
}

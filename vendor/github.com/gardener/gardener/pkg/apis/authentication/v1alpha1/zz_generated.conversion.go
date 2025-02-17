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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	authentication "github.com/gardener/gardener/pkg/apis/authentication"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddConversionFunc((*authentication.KubeconfigRequest)(nil), (*AdminKubeconfigRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authentication_KubeconfigRequest_To_v1alpha1_AdminKubeconfigRequest(a.(*authentication.KubeconfigRequest), b.(*AdminKubeconfigRequest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*authentication.KubeconfigRequest)(nil), (*ViewerKubeconfigRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authentication_KubeconfigRequest_To_v1alpha1_ViewerKubeconfigRequest(a.(*authentication.KubeconfigRequest), b.(*ViewerKubeconfigRequest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*AdminKubeconfigRequest)(nil), (*authentication.KubeconfigRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AdminKubeconfigRequest_To_authentication_KubeconfigRequest(a.(*AdminKubeconfigRequest), b.(*authentication.KubeconfigRequest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*ViewerKubeconfigRequest)(nil), (*authentication.KubeconfigRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ViewerKubeconfigRequest_To_authentication_KubeconfigRequest(a.(*ViewerKubeconfigRequest), b.(*authentication.KubeconfigRequest), scope)
	}); err != nil {
		return err
	}
	return nil
}

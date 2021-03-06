/*
Copyright 2017 The Kubernetes Authors.

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

package builders

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// Deprecated
var _ SchemeFns = &DefaultSchemeFns{}

// Deprecated
var SchemeFnsSingleton = DefaultSchemeFns{}

// Deprecated
type DefaultSchemeFns struct {
}

func (DefaultSchemeFns) DefaultingFunction(interface{}) {}

func (DefaultSchemeFns) GetConversionFunctions() []interface{} { return []interface{}{} }

func (DefaultSchemeFns) Register(scheme *runtime.Scheme) error { return nil }

func (DefaultSchemeFns) FieldSelectorConversion(label, value string) (string, string, error) {
	return runtime.DefaultMetaV1FieldSelectorConversion(label, value)
}

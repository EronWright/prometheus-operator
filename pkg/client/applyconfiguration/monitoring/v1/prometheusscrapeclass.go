// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// PrometheusScrapeClassApplyConfiguration represents an declarative configuration of the PrometheusScrapeClass type for use
// with apply.
type PrometheusScrapeClassApplyConfiguration struct {
	CommonScrapeClassFieldsApplyConfiguration `json:",inline"`
}

// PrometheusScrapeClassApplyConfiguration constructs an declarative configuration of the PrometheusScrapeClass type for use with
// apply.
func PrometheusScrapeClass() *PrometheusScrapeClassApplyConfiguration {
	return &PrometheusScrapeClassApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *PrometheusScrapeClassApplyConfiguration) WithName(value string) *PrometheusScrapeClassApplyConfiguration {
	b.Name = &value
	return b
}

// WithDefault sets the Default field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Default field is set to the value of the last call.
func (b *PrometheusScrapeClassApplyConfiguration) WithDefault(value bool) *PrometheusScrapeClassApplyConfiguration {
	b.Default = &value
	return b
}

// WithTLSConfig sets the TLSConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLSConfig field is set to the value of the last call.
func (b *PrometheusScrapeClassApplyConfiguration) WithTLSConfig(value *TLSConfigApplyConfiguration) *PrometheusScrapeClassApplyConfiguration {
	b.TLSConfig = value
	return b
}

// WithAuthorization sets the Authorization field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Authorization field is set to the value of the last call.
func (b *PrometheusScrapeClassApplyConfiguration) WithAuthorization(value *AuthorizationApplyConfiguration) *PrometheusScrapeClassApplyConfiguration {
	b.Authorization = value
	return b
}

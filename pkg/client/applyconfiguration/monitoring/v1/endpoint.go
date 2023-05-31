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

import (
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	corev1 "k8s.io/api/core/v1"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// EndpointApplyConfiguration represents an declarative configuration of the Endpoint type for use
// with apply.
type EndpointApplyConfiguration struct {
	Port                 *string                              `json:"port,omitempty"`
	TargetPort           *intstr.IntOrString                  `json:"targetPort,omitempty"`
	Path                 *string                              `json:"path,omitempty"`
	Scheme               *string                              `json:"scheme,omitempty"`
	Params               map[string][]string                  `json:"params,omitempty"`
	Interval             *v1.Duration                         `json:"interval,omitempty"`
	ScrapeTimeout        *v1.Duration                         `json:"scrapeTimeout,omitempty"`
	TLSConfig            *TLSConfigApplyConfiguration         `json:"tlsConfig,omitempty"`
	BearerTokenFile      *string                              `json:"bearerTokenFile,omitempty"`
	BearerTokenSecret    *corev1.SecretKeySelector            `json:"bearerTokenSecret,omitempty"`
	Authorization        *SafeAuthorizationApplyConfiguration `json:"authorization,omitempty"`
	HonorLabels          *bool                                `json:"honorLabels,omitempty"`
	HonorTimestamps      *bool                                `json:"honorTimestamps,omitempty"`
	BasicAuth            *BasicAuthApplyConfiguration         `json:"basicAuth,omitempty"`
	OAuth2               *OAuth2ApplyConfiguration            `json:"oauth2,omitempty"`
	MetricRelabelConfigs []*v1.RelabelConfig                  `json:"metricRelabelings,omitempty"`
	RelabelConfigs       []*v1.RelabelConfig                  `json:"relabelings,omitempty"`
	ProxyURL             *string                              `json:"proxyUrl,omitempty"`
	FollowRedirects      *bool                                `json:"followRedirects,omitempty"`
	EnableHttp2          *bool                                `json:"enableHttp2,omitempty"`
	FilterRunning        *bool                                `json:"filterRunning,omitempty"`
	ScrapeClass          *string                              `json:"scrapeClass,omitempty"`
}

// EndpointApplyConfiguration constructs an declarative configuration of the Endpoint type for use with
// apply.
func Endpoint() *EndpointApplyConfiguration {
	return &EndpointApplyConfiguration{}
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithPort(value string) *EndpointApplyConfiguration {
	b.Port = &value
	return b
}

// WithTargetPort sets the TargetPort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetPort field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithTargetPort(value intstr.IntOrString) *EndpointApplyConfiguration {
	b.TargetPort = &value
	return b
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithPath(value string) *EndpointApplyConfiguration {
	b.Path = &value
	return b
}

// WithScheme sets the Scheme field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Scheme field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithScheme(value string) *EndpointApplyConfiguration {
	b.Scheme = &value
	return b
}

// WithParams puts the entries into the Params field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Params field,
// overwriting an existing map entries in Params field with the same key.
func (b *EndpointApplyConfiguration) WithParams(entries map[string][]string) *EndpointApplyConfiguration {
	if b.Params == nil && len(entries) > 0 {
		b.Params = make(map[string][]string, len(entries))
	}
	for k, v := range entries {
		b.Params[k] = v
	}
	return b
}

// WithInterval sets the Interval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Interval field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithInterval(value v1.Duration) *EndpointApplyConfiguration {
	b.Interval = &value
	return b
}

// WithScrapeTimeout sets the ScrapeTimeout field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ScrapeTimeout field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithScrapeTimeout(value v1.Duration) *EndpointApplyConfiguration {
	b.ScrapeTimeout = &value
	return b
}

// WithTLSConfig sets the TLSConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLSConfig field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithTLSConfig(value *TLSConfigApplyConfiguration) *EndpointApplyConfiguration {
	b.TLSConfig = value
	return b
}

// WithBearerTokenFile sets the BearerTokenFile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BearerTokenFile field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithBearerTokenFile(value string) *EndpointApplyConfiguration {
	b.BearerTokenFile = &value
	return b
}

// WithBearerTokenSecret sets the BearerTokenSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BearerTokenSecret field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithBearerTokenSecret(value corev1.SecretKeySelector) *EndpointApplyConfiguration {
	b.BearerTokenSecret = &value
	return b
}

// WithAuthorization sets the Authorization field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Authorization field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithAuthorization(value *SafeAuthorizationApplyConfiguration) *EndpointApplyConfiguration {
	b.Authorization = value
	return b
}

// WithHonorLabels sets the HonorLabels field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HonorLabels field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithHonorLabels(value bool) *EndpointApplyConfiguration {
	b.HonorLabels = &value
	return b
}

// WithHonorTimestamps sets the HonorTimestamps field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HonorTimestamps field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithHonorTimestamps(value bool) *EndpointApplyConfiguration {
	b.HonorTimestamps = &value
	return b
}

// WithBasicAuth sets the BasicAuth field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BasicAuth field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithBasicAuth(value *BasicAuthApplyConfiguration) *EndpointApplyConfiguration {
	b.BasicAuth = value
	return b
}

// WithOAuth2 sets the OAuth2 field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OAuth2 field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithOAuth2(value *OAuth2ApplyConfiguration) *EndpointApplyConfiguration {
	b.OAuth2 = value
	return b
}

// WithMetricRelabelConfigs adds the given value to the MetricRelabelConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the MetricRelabelConfigs field.
func (b *EndpointApplyConfiguration) WithMetricRelabelConfigs(values ...**v1.RelabelConfig) *EndpointApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithMetricRelabelConfigs")
		}
		b.MetricRelabelConfigs = append(b.MetricRelabelConfigs, *values[i])
	}
	return b
}

// WithRelabelConfigs adds the given value to the RelabelConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the RelabelConfigs field.
func (b *EndpointApplyConfiguration) WithRelabelConfigs(values ...**v1.RelabelConfig) *EndpointApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRelabelConfigs")
		}
		b.RelabelConfigs = append(b.RelabelConfigs, *values[i])
	}
	return b
}

// WithProxyURL sets the ProxyURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProxyURL field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithProxyURL(value string) *EndpointApplyConfiguration {
	b.ProxyURL = &value
	return b
}

// WithFollowRedirects sets the FollowRedirects field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FollowRedirects field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithFollowRedirects(value bool) *EndpointApplyConfiguration {
	b.FollowRedirects = &value
	return b
}

// WithEnableHttp2 sets the EnableHttp2 field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableHttp2 field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithEnableHttp2(value bool) *EndpointApplyConfiguration {
	b.EnableHttp2 = &value
	return b
}

// WithFilterRunning sets the FilterRunning field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FilterRunning field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithFilterRunning(value bool) *EndpointApplyConfiguration {
	b.FilterRunning = &value
	return b
}

// WithScrapeClass sets the ScrapeClass field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ScrapeClass field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithScrapeClass(value string) *EndpointApplyConfiguration {
	b.ScrapeClass = &value
	return b
}

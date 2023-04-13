/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// NetworkLoadBalancerForwardingRuleHealthCheck struct for NetworkLoadBalancerForwardingRuleHealthCheck
type NetworkLoadBalancerForwardingRuleHealthCheck struct {
	// The maximum time in milliseconds to wait for the client to acknowledge or send data; default is 50,000 (50 seconds).
	ClientTimeout *int32 `json:"clientTimeout,omitempty"`
	// The maximum time in milliseconds to wait for a connection attempt to a target to succeed; default is 5000 (five seconds).
	ConnectTimeout *int32 `json:"connectTimeout,omitempty"`
	// The maximum number of attempts to reconnect to a target after a connection failure. Valid range is 0 to 65535 and default is three reconnection attempts.
	Retries *int32 `json:"retries,omitempty"`
	// The maximum time in milliseconds that a target can remain inactive; default is 50,000 (50 seconds).
	TargetTimeout *int32 `json:"targetTimeout,omitempty"`
}

// NewNetworkLoadBalancerForwardingRuleHealthCheck instantiates a new NetworkLoadBalancerForwardingRuleHealthCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkLoadBalancerForwardingRuleHealthCheck() *NetworkLoadBalancerForwardingRuleHealthCheck {
	this := NetworkLoadBalancerForwardingRuleHealthCheck{}

	return &this
}

// NewNetworkLoadBalancerForwardingRuleHealthCheckWithDefaults instantiates a new NetworkLoadBalancerForwardingRuleHealthCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkLoadBalancerForwardingRuleHealthCheckWithDefaults() *NetworkLoadBalancerForwardingRuleHealthCheck {
	this := NetworkLoadBalancerForwardingRuleHealthCheck{}
	return &this
}

// GetClientTimeout returns the ClientTimeout field value
// If the value is explicit nil, nil is returned
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetClientTimeout() *int32 {
	if o == nil {
		return nil
	}

	return o.ClientTimeout

}

// GetClientTimeoutOk returns a tuple with the ClientTimeout field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetClientTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.ClientTimeout, true
}

// SetClientTimeout sets field value
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) SetClientTimeout(v int32) {

	o.ClientTimeout = &v

}

// HasClientTimeout returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) HasClientTimeout() bool {
	if o != nil && o.ClientTimeout != nil {
		return true
	}

	return false
}

// GetConnectTimeout returns the ConnectTimeout field value
// If the value is explicit nil, nil is returned
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetConnectTimeout() *int32 {
	if o == nil {
		return nil
	}

	return o.ConnectTimeout

}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetConnectTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.ConnectTimeout, true
}

// SetConnectTimeout sets field value
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) SetConnectTimeout(v int32) {

	o.ConnectTimeout = &v

}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) HasConnectTimeout() bool {
	if o != nil && o.ConnectTimeout != nil {
		return true
	}

	return false
}

// GetRetries returns the Retries field value
// If the value is explicit nil, nil is returned
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetRetries() *int32 {
	if o == nil {
		return nil
	}

	return o.Retries

}

// GetRetriesOk returns a tuple with the Retries field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetRetriesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Retries, true
}

// SetRetries sets field value
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) SetRetries(v int32) {

	o.Retries = &v

}

// HasRetries returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) HasRetries() bool {
	if o != nil && o.Retries != nil {
		return true
	}

	return false
}

// GetTargetTimeout returns the TargetTimeout field value
// If the value is explicit nil, nil is returned
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetTargetTimeout() *int32 {
	if o == nil {
		return nil
	}

	return o.TargetTimeout

}

// GetTargetTimeoutOk returns a tuple with the TargetTimeout field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetTargetTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.TargetTimeout, true
}

// SetTargetTimeout sets field value
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) SetTargetTimeout(v int32) {

	o.TargetTimeout = &v

}

// HasTargetTimeout returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) HasTargetTimeout() bool {
	if o != nil && o.TargetTimeout != nil {
		return true
	}

	return false
}

func (o NetworkLoadBalancerForwardingRuleHealthCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientTimeout != nil {
		toSerialize["clientTimeout"] = o.ClientTimeout
	}

	if o.ConnectTimeout != nil {
		toSerialize["connectTimeout"] = o.ConnectTimeout
	}

	if o.Retries != nil {
		toSerialize["retries"] = o.Retries
	}

	if o.TargetTimeout != nil {
		toSerialize["targetTimeout"] = o.TargetTimeout
	}

	return json.Marshal(toSerialize)
}

type NullableNetworkLoadBalancerForwardingRuleHealthCheck struct {
	value *NetworkLoadBalancerForwardingRuleHealthCheck
	isSet bool
}

func (v NullableNetworkLoadBalancerForwardingRuleHealthCheck) Get() *NetworkLoadBalancerForwardingRuleHealthCheck {
	return v.value
}

func (v *NullableNetworkLoadBalancerForwardingRuleHealthCheck) Set(val *NetworkLoadBalancerForwardingRuleHealthCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkLoadBalancerForwardingRuleHealthCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkLoadBalancerForwardingRuleHealthCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkLoadBalancerForwardingRuleHealthCheck(val *NetworkLoadBalancerForwardingRuleHealthCheck) *NullableNetworkLoadBalancerForwardingRuleHealthCheck {
	return &NullableNetworkLoadBalancerForwardingRuleHealthCheck{value: val, isSet: true}
}

func (v NullableNetworkLoadBalancerForwardingRuleHealthCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkLoadBalancerForwardingRuleHealthCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

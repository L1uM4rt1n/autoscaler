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

// IpBlocks struct for IpBlocks
type IpBlocks struct {
	Links *PaginationLinks `json:"_links,omitempty"`
	// URL to the object representation (absolute path).
	Href *string `json:"href,omitempty"`
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// Array of items in the collection.
	Items *[]IpBlock `json:"items,omitempty"`
	// The limit, specified in the request (if not specified, the endpoint's default pagination limit is used).
	Limit *float32 `json:"limit,omitempty"`
	// The offset, specified in the request (if not is specified, 0 is used by default).
	Offset *float32 `json:"offset,omitempty"`
	// The type of object that has been created.
	Type *Type `json:"type,omitempty"`
}

// NewIpBlocks instantiates a new IpBlocks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpBlocks() *IpBlocks {
	this := IpBlocks{}

	return &this
}

// NewIpBlocksWithDefaults instantiates a new IpBlocks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpBlocksWithDefaults() *IpBlocks {
	this := IpBlocks{}
	return &this
}

// GetLinks returns the Links field value
// If the value is explicit nil, nil is returned
func (o *IpBlocks) GetLinks() *PaginationLinks {
	if o == nil {
		return nil
	}

	return o.Links

}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpBlocks) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil {
		return nil, false
	}

	return o.Links, true
}

// SetLinks sets field value
func (o *IpBlocks) SetLinks(v PaginationLinks) {

	o.Links = &v

}

// HasLinks returns a boolean if a field has been set.
func (o *IpBlocks) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, nil is returned
func (o *IpBlocks) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpBlocks) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *IpBlocks) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *IpBlocks) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetId returns the Id field value
// If the value is explicit nil, nil is returned
func (o *IpBlocks) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpBlocks) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *IpBlocks) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *IpBlocks) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, nil is returned
func (o *IpBlocks) GetItems() *[]IpBlock {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpBlocks) GetItemsOk() (*[]IpBlock, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *IpBlocks) SetItems(v []IpBlock) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *IpBlocks) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// GetLimit returns the Limit field value
// If the value is explicit nil, nil is returned
func (o *IpBlocks) GetLimit() *float32 {
	if o == nil {
		return nil
	}

	return o.Limit

}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpBlocks) GetLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Limit, true
}

// SetLimit sets field value
func (o *IpBlocks) SetLimit(v float32) {

	o.Limit = &v

}

// HasLimit returns a boolean if a field has been set.
func (o *IpBlocks) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// GetOffset returns the Offset field value
// If the value is explicit nil, nil is returned
func (o *IpBlocks) GetOffset() *float32 {
	if o == nil {
		return nil
	}

	return o.Offset

}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpBlocks) GetOffsetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Offset, true
}

// SetOffset sets field value
func (o *IpBlocks) SetOffset(v float32) {

	o.Offset = &v

}

// HasOffset returns a boolean if a field has been set.
func (o *IpBlocks) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, nil is returned
func (o *IpBlocks) GetType() *Type {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpBlocks) GetTypeOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *IpBlocks) SetType(v Type) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *IpBlocks) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

func (o IpBlocks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}

	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	return json.Marshal(toSerialize)
}

type NullableIpBlocks struct {
	value *IpBlocks
	isSet bool
}

func (v NullableIpBlocks) Get() *IpBlocks {
	return v.value
}

func (v *NullableIpBlocks) Set(val *IpBlocks) {
	v.value = val
	v.isSet = true
}

func (v NullableIpBlocks) IsSet() bool {
	return v.isSet
}

func (v *NullableIpBlocks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpBlocks(val *IpBlocks) *NullableIpBlocks {
	return &NullableIpBlocks{value: val, isSet: true}
}

func (v NullableIpBlocks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpBlocks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

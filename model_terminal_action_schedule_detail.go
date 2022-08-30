/*
Management API

Configure and manage your Adyen company and merchant accounts, stores, and payment terminals. ## Authentication Each request to the Management API must be signed with an API key. [Generate your API key](https://docs.adyen.com/development-resources/api-credentials#generate-api-key) in the Customer Area and then set this key to the `X-API-Key` header value.  To access the live endpoints, you need to generate a new API key in your live Customer Area. ## Versioning  Management API handles versioning as part of the endpoint URL. For example, to send a request to version 1 of the `/companies/{companyId}/webhooks` endpoint, use:  ```text https://management-test.adyen.com/v1/companies/{companyId}/webhooks ```

API version: 1
Contact: developer-experience@adyen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TerminalActionScheduleDetail struct for TerminalActionScheduleDetail
type TerminalActionScheduleDetail struct {
	Id *string `json:"id,omitempty"`
	TerminalId *string `json:"terminalId,omitempty"`
}

// NewTerminalActionScheduleDetail instantiates a new TerminalActionScheduleDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalActionScheduleDetail() *TerminalActionScheduleDetail {
	this := TerminalActionScheduleDetail{}
	return &this
}

// NewTerminalActionScheduleDetailWithDefaults instantiates a new TerminalActionScheduleDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalActionScheduleDetailWithDefaults() *TerminalActionScheduleDetail {
	this := TerminalActionScheduleDetail{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TerminalActionScheduleDetail) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalActionScheduleDetail) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TerminalActionScheduleDetail) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TerminalActionScheduleDetail) SetId(v string) {
	o.Id = &v
}

// GetTerminalId returns the TerminalId field value if set, zero value otherwise.
func (o *TerminalActionScheduleDetail) GetTerminalId() string {
	if o == nil || o.TerminalId == nil {
		var ret string
		return ret
	}
	return *o.TerminalId
}

// GetTerminalIdOk returns a tuple with the TerminalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalActionScheduleDetail) GetTerminalIdOk() (*string, bool) {
	if o == nil || o.TerminalId == nil {
		return nil, false
	}
	return o.TerminalId, true
}

// HasTerminalId returns a boolean if a field has been set.
func (o *TerminalActionScheduleDetail) HasTerminalId() bool {
	if o != nil && o.TerminalId != nil {
		return true
	}

	return false
}

// SetTerminalId gets a reference to the given string and assigns it to the TerminalId field.
func (o *TerminalActionScheduleDetail) SetTerminalId(v string) {
	o.TerminalId = &v
}

func (o TerminalActionScheduleDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TerminalId != nil {
		toSerialize["terminalId"] = o.TerminalId
	}
	return json.Marshal(toSerialize)
}

type NullableTerminalActionScheduleDetail struct {
	value *TerminalActionScheduleDetail
	isSet bool
}

func (v NullableTerminalActionScheduleDetail) Get() *TerminalActionScheduleDetail {
	return v.value
}

func (v *NullableTerminalActionScheduleDetail) Set(val *TerminalActionScheduleDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalActionScheduleDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalActionScheduleDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalActionScheduleDetail(val *TerminalActionScheduleDetail) *NullableTerminalActionScheduleDetail {
	return &NullableTerminalActionScheduleDetail{value: val, isSet: true}
}

func (v NullableTerminalActionScheduleDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalActionScheduleDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



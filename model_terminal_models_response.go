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

// TerminalModelsResponse struct for TerminalModelsResponse
type TerminalModelsResponse struct {
	// The terminal models that the API credential has access to.
	Data []IdName `json:"data,omitempty"`
}

// NewTerminalModelsResponse instantiates a new TerminalModelsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalModelsResponse() *TerminalModelsResponse {
	this := TerminalModelsResponse{}
	return &this
}

// NewTerminalModelsResponseWithDefaults instantiates a new TerminalModelsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalModelsResponseWithDefaults() *TerminalModelsResponse {
	this := TerminalModelsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TerminalModelsResponse) GetData() []IdName {
	if o == nil || o.Data == nil {
		var ret []IdName
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalModelsResponse) GetDataOk() ([]IdName, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TerminalModelsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []IdName and assigns it to the Data field.
func (o *TerminalModelsResponse) SetData(v []IdName) {
	o.Data = v
}

func (o TerminalModelsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableTerminalModelsResponse struct {
	value *TerminalModelsResponse
	isSet bool
}

func (v NullableTerminalModelsResponse) Get() *TerminalModelsResponse {
	return v.value
}

func (v *NullableTerminalModelsResponse) Set(val *TerminalModelsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalModelsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalModelsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalModelsResponse(val *TerminalModelsResponse) *NullableTerminalModelsResponse {
	return &NullableTerminalModelsResponse{value: val, isSet: true}
}

func (v NullableTerminalModelsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalModelsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



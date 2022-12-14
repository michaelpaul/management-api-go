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

// AdditionalSettingsResponse struct for AdditionalSettingsResponse
type AdditionalSettingsResponse struct {
	// Object containing list of event codes for which the notifcation will not be sent. 
	ExcludeEventCodes []string `json:"excludeEventCodes,omitempty"`
	// Object containing list of event codes for which the notifcation will be sent. 
	IncludeEventCodes []string `json:"includeEventCodes,omitempty"`
	// Object containing boolean key-value pairs. The key can be any [standard webhook additional setting](https://docs.adyen.com/development-resources/webhooks/additional-settings), and the value indicates if the setting is enabled. For example, `captureDelayHours`: **true** means the standard notifications you get will contain the number of hours remaining until the payment will be captured.
	Properties *map[string]bool `json:"properties,omitempty"`
}

// NewAdditionalSettingsResponse instantiates a new AdditionalSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalSettingsResponse() *AdditionalSettingsResponse {
	this := AdditionalSettingsResponse{}
	return &this
}

// NewAdditionalSettingsResponseWithDefaults instantiates a new AdditionalSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalSettingsResponseWithDefaults() *AdditionalSettingsResponse {
	this := AdditionalSettingsResponse{}
	return &this
}

// GetExcludeEventCodes returns the ExcludeEventCodes field value if set, zero value otherwise.
func (o *AdditionalSettingsResponse) GetExcludeEventCodes() []string {
	if o == nil || o.ExcludeEventCodes == nil {
		var ret []string
		return ret
	}
	return o.ExcludeEventCodes
}

// GetExcludeEventCodesOk returns a tuple with the ExcludeEventCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalSettingsResponse) GetExcludeEventCodesOk() ([]string, bool) {
	if o == nil || o.ExcludeEventCodes == nil {
		return nil, false
	}
	return o.ExcludeEventCodes, true
}

// HasExcludeEventCodes returns a boolean if a field has been set.
func (o *AdditionalSettingsResponse) HasExcludeEventCodes() bool {
	if o != nil && o.ExcludeEventCodes != nil {
		return true
	}

	return false
}

// SetExcludeEventCodes gets a reference to the given []string and assigns it to the ExcludeEventCodes field.
func (o *AdditionalSettingsResponse) SetExcludeEventCodes(v []string) {
	o.ExcludeEventCodes = v
}

// GetIncludeEventCodes returns the IncludeEventCodes field value if set, zero value otherwise.
func (o *AdditionalSettingsResponse) GetIncludeEventCodes() []string {
	if o == nil || o.IncludeEventCodes == nil {
		var ret []string
		return ret
	}
	return o.IncludeEventCodes
}

// GetIncludeEventCodesOk returns a tuple with the IncludeEventCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalSettingsResponse) GetIncludeEventCodesOk() ([]string, bool) {
	if o == nil || o.IncludeEventCodes == nil {
		return nil, false
	}
	return o.IncludeEventCodes, true
}

// HasIncludeEventCodes returns a boolean if a field has been set.
func (o *AdditionalSettingsResponse) HasIncludeEventCodes() bool {
	if o != nil && o.IncludeEventCodes != nil {
		return true
	}

	return false
}

// SetIncludeEventCodes gets a reference to the given []string and assigns it to the IncludeEventCodes field.
func (o *AdditionalSettingsResponse) SetIncludeEventCodes(v []string) {
	o.IncludeEventCodes = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *AdditionalSettingsResponse) GetProperties() map[string]bool {
	if o == nil || o.Properties == nil {
		var ret map[string]bool
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalSettingsResponse) GetPropertiesOk() (*map[string]bool, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *AdditionalSettingsResponse) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]bool and assigns it to the Properties field.
func (o *AdditionalSettingsResponse) SetProperties(v map[string]bool) {
	o.Properties = &v
}

func (o AdditionalSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExcludeEventCodes != nil {
		toSerialize["excludeEventCodes"] = o.ExcludeEventCodes
	}
	if o.IncludeEventCodes != nil {
		toSerialize["includeEventCodes"] = o.IncludeEventCodes
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableAdditionalSettingsResponse struct {
	value *AdditionalSettingsResponse
	isSet bool
}

func (v NullableAdditionalSettingsResponse) Get() *AdditionalSettingsResponse {
	return v.value
}

func (v *NullableAdditionalSettingsResponse) Set(val *AdditionalSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalSettingsResponse(val *AdditionalSettingsResponse) *NullableAdditionalSettingsResponse {
	return &NullableAdditionalSettingsResponse{value: val, isSet: true}
}

func (v NullableAdditionalSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



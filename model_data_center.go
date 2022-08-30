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

// DataCenter struct for DataCenter
type DataCenter struct {
	// The unique [live URL prefix](https://docs.adyen.com/development-resources/live-endpoints#live-url-prefix) for your live endpoint. Each data center has its own live URL prefix.  This field is empty for requests made in the test environment.
	LivePrefix *string `json:"livePrefix,omitempty"`
	// The name assigned to a data center, for example **EU** for the European data center. Possible values are:  * **default**: the European data center. This value is always returned in the test environment.  * **AU** * **EU** * **US**
	Name *string `json:"name,omitempty"`
}

// NewDataCenter instantiates a new DataCenter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataCenter() *DataCenter {
	this := DataCenter{}
	return &this
}

// NewDataCenterWithDefaults instantiates a new DataCenter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataCenterWithDefaults() *DataCenter {
	this := DataCenter{}
	return &this
}

// GetLivePrefix returns the LivePrefix field value if set, zero value otherwise.
func (o *DataCenter) GetLivePrefix() string {
	if o == nil || o.LivePrefix == nil {
		var ret string
		return ret
	}
	return *o.LivePrefix
}

// GetLivePrefixOk returns a tuple with the LivePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCenter) GetLivePrefixOk() (*string, bool) {
	if o == nil || o.LivePrefix == nil {
		return nil, false
	}
	return o.LivePrefix, true
}

// HasLivePrefix returns a boolean if a field has been set.
func (o *DataCenter) HasLivePrefix() bool {
	if o != nil && o.LivePrefix != nil {
		return true
	}

	return false
}

// SetLivePrefix gets a reference to the given string and assigns it to the LivePrefix field.
func (o *DataCenter) SetLivePrefix(v string) {
	o.LivePrefix = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DataCenter) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCenter) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DataCenter) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DataCenter) SetName(v string) {
	o.Name = &v
}

func (o DataCenter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LivePrefix != nil {
		toSerialize["livePrefix"] = o.LivePrefix
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDataCenter struct {
	value *DataCenter
	isSet bool
}

func (v NullableDataCenter) Get() *DataCenter {
	return v.value
}

func (v *NullableDataCenter) Set(val *DataCenter) {
	v.value = val
	v.isSet = true
}

func (v NullableDataCenter) IsSet() bool {
	return v.isSet
}

func (v *NullableDataCenter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataCenter(val *DataCenter) *NullableDataCenter {
	return &NullableDataCenter{value: val, isSet: true}
}

func (v NullableDataCenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataCenter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


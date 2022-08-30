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

// JSONObject struct for JSONObject
type JSONObject struct {
	Paths []JSONPath `json:"paths,omitempty"`
	RootPath *JSONPath `json:"rootPath,omitempty"`
}

// NewJSONObject instantiates a new JSONObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSONObject() *JSONObject {
	this := JSONObject{}
	return &this
}

// NewJSONObjectWithDefaults instantiates a new JSONObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSONObjectWithDefaults() *JSONObject {
	this := JSONObject{}
	return &this
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *JSONObject) GetPaths() []JSONPath {
	if o == nil || o.Paths == nil {
		var ret []JSONPath
		return ret
	}
	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONObject) GetPathsOk() ([]JSONPath, bool) {
	if o == nil || o.Paths == nil {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *JSONObject) HasPaths() bool {
	if o != nil && o.Paths != nil {
		return true
	}

	return false
}

// SetPaths gets a reference to the given []JSONPath and assigns it to the Paths field.
func (o *JSONObject) SetPaths(v []JSONPath) {
	o.Paths = v
}

// GetRootPath returns the RootPath field value if set, zero value otherwise.
func (o *JSONObject) GetRootPath() JSONPath {
	if o == nil || o.RootPath == nil {
		var ret JSONPath
		return ret
	}
	return *o.RootPath
}

// GetRootPathOk returns a tuple with the RootPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONObject) GetRootPathOk() (*JSONPath, bool) {
	if o == nil || o.RootPath == nil {
		return nil, false
	}
	return o.RootPath, true
}

// HasRootPath returns a boolean if a field has been set.
func (o *JSONObject) HasRootPath() bool {
	if o != nil && o.RootPath != nil {
		return true
	}

	return false
}

// SetRootPath gets a reference to the given JSONPath and assigns it to the RootPath field.
func (o *JSONObject) SetRootPath(v JSONPath) {
	o.RootPath = &v
}

func (o JSONObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Paths != nil {
		toSerialize["paths"] = o.Paths
	}
	if o.RootPath != nil {
		toSerialize["rootPath"] = o.RootPath
	}
	return json.Marshal(toSerialize)
}

type NullableJSONObject struct {
	value *JSONObject
	isSet bool
}

func (v NullableJSONObject) Get() *JSONObject {
	return v.value
}

func (v *NullableJSONObject) Set(val *JSONObject) {
	v.value = val
	v.isSet = true
}

func (v NullableJSONObject) IsSet() bool {
	return v.isSet
}

func (v *NullableJSONObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSONObject(val *JSONObject) *NullableJSONObject {
	return &NullableJSONObject{value: val, isSet: true}
}

func (v NullableJSONObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSONObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



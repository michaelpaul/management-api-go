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

// Address2 struct for Address2
type Address2 struct {
	// The name of the city.
	City *string `json:"city,omitempty"`
	// The two-letter country code in [ISO_3166-1_alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format.
	Country string `json:"country"`
	// The street address.
	Line1 *string `json:"line1,omitempty"`
	// Second address line.
	Line2 *string `json:"line2,omitempty"`
	// Third address line.
	Line3 *string `json:"line3,omitempty"`
	// The postal code.
	PostalCode *string `json:"postalCode,omitempty"`
	// The state or province code as defined in [ISO 3166-2](https://www.iso.org/standard/72483.html). For example, **ON** for Ontario, Canada.  Required for the following countries:  - Australia - Brazil - Canada - India - Mexico - New Zealand - United States
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
}

// NewAddress2 instantiates a new Address2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress2(country string) *Address2 {
	this := Address2{}
	this.Country = country
	return &this
}

// NewAddress2WithDefaults instantiates a new Address2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddress2WithDefaults() *Address2 {
	this := Address2{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *Address2) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address2) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *Address2) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *Address2) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value
func (o *Address2) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *Address2) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *Address2) SetCountry(v string) {
	o.Country = v
}

// GetLine1 returns the Line1 field value if set, zero value otherwise.
func (o *Address2) GetLine1() string {
	if o == nil || o.Line1 == nil {
		var ret string
		return ret
	}
	return *o.Line1
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address2) GetLine1Ok() (*string, bool) {
	if o == nil || o.Line1 == nil {
		return nil, false
	}
	return o.Line1, true
}

// HasLine1 returns a boolean if a field has been set.
func (o *Address2) HasLine1() bool {
	if o != nil && o.Line1 != nil {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given string and assigns it to the Line1 field.
func (o *Address2) SetLine1(v string) {
	o.Line1 = &v
}

// GetLine2 returns the Line2 field value if set, zero value otherwise.
func (o *Address2) GetLine2() string {
	if o == nil || o.Line2 == nil {
		var ret string
		return ret
	}
	return *o.Line2
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address2) GetLine2Ok() (*string, bool) {
	if o == nil || o.Line2 == nil {
		return nil, false
	}
	return o.Line2, true
}

// HasLine2 returns a boolean if a field has been set.
func (o *Address2) HasLine2() bool {
	if o != nil && o.Line2 != nil {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given string and assigns it to the Line2 field.
func (o *Address2) SetLine2(v string) {
	o.Line2 = &v
}

// GetLine3 returns the Line3 field value if set, zero value otherwise.
func (o *Address2) GetLine3() string {
	if o == nil || o.Line3 == nil {
		var ret string
		return ret
	}
	return *o.Line3
}

// GetLine3Ok returns a tuple with the Line3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address2) GetLine3Ok() (*string, bool) {
	if o == nil || o.Line3 == nil {
		return nil, false
	}
	return o.Line3, true
}

// HasLine3 returns a boolean if a field has been set.
func (o *Address2) HasLine3() bool {
	if o != nil && o.Line3 != nil {
		return true
	}

	return false
}

// SetLine3 gets a reference to the given string and assigns it to the Line3 field.
func (o *Address2) SetLine3(v string) {
	o.Line3 = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *Address2) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address2) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *Address2) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *Address2) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetStateOrProvince returns the StateOrProvince field value if set, zero value otherwise.
func (o *Address2) GetStateOrProvince() string {
	if o == nil || o.StateOrProvince == nil {
		var ret string
		return ret
	}
	return *o.StateOrProvince
}

// GetStateOrProvinceOk returns a tuple with the StateOrProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address2) GetStateOrProvinceOk() (*string, bool) {
	if o == nil || o.StateOrProvince == nil {
		return nil, false
	}
	return o.StateOrProvince, true
}

// HasStateOrProvince returns a boolean if a field has been set.
func (o *Address2) HasStateOrProvince() bool {
	if o != nil && o.StateOrProvince != nil {
		return true
	}

	return false
}

// SetStateOrProvince gets a reference to the given string and assigns it to the StateOrProvince field.
func (o *Address2) SetStateOrProvince(v string) {
	o.StateOrProvince = &v
}

func (o Address2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if true {
		toSerialize["country"] = o.Country
	}
	if o.Line1 != nil {
		toSerialize["line1"] = o.Line1
	}
	if o.Line2 != nil {
		toSerialize["line2"] = o.Line2
	}
	if o.Line3 != nil {
		toSerialize["line3"] = o.Line3
	}
	if o.PostalCode != nil {
		toSerialize["postalCode"] = o.PostalCode
	}
	if o.StateOrProvince != nil {
		toSerialize["stateOrProvince"] = o.StateOrProvince
	}
	return json.Marshal(toSerialize)
}

type NullableAddress2 struct {
	value *Address2
	isSet bool
}

func (v NullableAddress2) Get() *Address2 {
	return v.value
}

func (v *NullableAddress2) Set(val *Address2) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress2) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress2(val *Address2) *NullableAddress2 {
	return &NullableAddress2{value: val, isSet: true}
}

func (v NullableAddress2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// CreateMerchantUserRequest struct for CreateMerchantUserRequest
type CreateMerchantUserRequest struct {
	// The list of [account groups](https://docs.adyen.com/account/account-structure#account-groups) associated with this user.
	AccountGroups []string `json:"accountGroups,omitempty"`
	// The email address of the user.
	Email string `json:"email"`
	Name Name `json:"name"`
	// The list of [roles](https://docs.adyen.com/account/user-roles) for this user.
	Roles []string `json:"roles,omitempty"`
	// The [tz database name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) of the time zone of the user. For example, **Europe/Amsterdam**.
	TimeZoneCode *string `json:"timeZoneCode,omitempty"`
	// The username for this user. Allowed length: 255 alphanumeric characters.
	Username string `json:"username"`
}

// NewCreateMerchantUserRequest instantiates a new CreateMerchantUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMerchantUserRequest(email string, name Name, username string) *CreateMerchantUserRequest {
	this := CreateMerchantUserRequest{}
	this.Email = email
	this.Name = name
	this.Username = username
	return &this
}

// NewCreateMerchantUserRequestWithDefaults instantiates a new CreateMerchantUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMerchantUserRequestWithDefaults() *CreateMerchantUserRequest {
	this := CreateMerchantUserRequest{}
	return &this
}

// GetAccountGroups returns the AccountGroups field value if set, zero value otherwise.
func (o *CreateMerchantUserRequest) GetAccountGroups() []string {
	if o == nil || o.AccountGroups == nil {
		var ret []string
		return ret
	}
	return o.AccountGroups
}

// GetAccountGroupsOk returns a tuple with the AccountGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantUserRequest) GetAccountGroupsOk() ([]string, bool) {
	if o == nil || o.AccountGroups == nil {
		return nil, false
	}
	return o.AccountGroups, true
}

// HasAccountGroups returns a boolean if a field has been set.
func (o *CreateMerchantUserRequest) HasAccountGroups() bool {
	if o != nil && o.AccountGroups != nil {
		return true
	}

	return false
}

// SetAccountGroups gets a reference to the given []string and assigns it to the AccountGroups field.
func (o *CreateMerchantUserRequest) SetAccountGroups(v []string) {
	o.AccountGroups = v
}

// GetEmail returns the Email field value
func (o *CreateMerchantUserRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateMerchantUserRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateMerchantUserRequest) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value
func (o *CreateMerchantUserRequest) GetName() Name {
	if o == nil {
		var ret Name
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateMerchantUserRequest) GetNameOk() (*Name, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateMerchantUserRequest) SetName(v Name) {
	o.Name = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *CreateMerchantUserRequest) GetRoles() []string {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantUserRequest) GetRolesOk() ([]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *CreateMerchantUserRequest) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *CreateMerchantUserRequest) SetRoles(v []string) {
	o.Roles = v
}

// GetTimeZoneCode returns the TimeZoneCode field value if set, zero value otherwise.
func (o *CreateMerchantUserRequest) GetTimeZoneCode() string {
	if o == nil || o.TimeZoneCode == nil {
		var ret string
		return ret
	}
	return *o.TimeZoneCode
}

// GetTimeZoneCodeOk returns a tuple with the TimeZoneCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantUserRequest) GetTimeZoneCodeOk() (*string, bool) {
	if o == nil || o.TimeZoneCode == nil {
		return nil, false
	}
	return o.TimeZoneCode, true
}

// HasTimeZoneCode returns a boolean if a field has been set.
func (o *CreateMerchantUserRequest) HasTimeZoneCode() bool {
	if o != nil && o.TimeZoneCode != nil {
		return true
	}

	return false
}

// SetTimeZoneCode gets a reference to the given string and assigns it to the TimeZoneCode field.
func (o *CreateMerchantUserRequest) SetTimeZoneCode(v string) {
	o.TimeZoneCode = &v
}

// GetUsername returns the Username field value
func (o *CreateMerchantUserRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CreateMerchantUserRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CreateMerchantUserRequest) SetUsername(v string) {
	o.Username = v
}

func (o CreateMerchantUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountGroups != nil {
		toSerialize["accountGroups"] = o.AccountGroups
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.TimeZoneCode != nil {
		toSerialize["timeZoneCode"] = o.TimeZoneCode
	}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableCreateMerchantUserRequest struct {
	value *CreateMerchantUserRequest
	isSet bool
}

func (v NullableCreateMerchantUserRequest) Get() *CreateMerchantUserRequest {
	return v.value
}

func (v *NullableCreateMerchantUserRequest) Set(val *CreateMerchantUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMerchantUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMerchantUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMerchantUserRequest(val *CreateMerchantUserRequest) *NullableCreateMerchantUserRequest {
	return &NullableCreateMerchantUserRequest{value: val, isSet: true}
}

func (v NullableCreateMerchantUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMerchantUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



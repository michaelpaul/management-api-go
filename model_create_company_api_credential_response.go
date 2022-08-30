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

// CreateCompanyApiCredentialResponse struct for CreateCompanyApiCredentialResponse
type CreateCompanyApiCredentialResponse struct {
	Links *ApiCredentialLinks `json:"_links,omitempty"`
	// Indicates if the API credential is enabled. Must be set to **true** to use the credential in your integration.
	Active bool `json:"active"`
	// List of IP addresses from which your client can make requests.  If the list is empty, we allow requests from any IP. If the list is not empty and we get a request from an IP which is not on the list, you get a security error.
	AllowedIpAddresses []string `json:"allowedIpAddresses"`
	// List containing the [allowed origins](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) linked to the API credential.
	AllowedOrigins []AllowedOrigin `json:"allowedOrigins,omitempty"`
	// The API key for the API credential that was created.
	ApiKey string `json:"apiKey"`
	// List of merchant accounts that the API credential has access to.
	AssociatedMerchantAccounts []string `json:"associatedMerchantAccounts"`
	// Public key used for [client-side authentication](https://docs.adyen.com/development-resources/client-side-authentication). The client key is required for Drop-in and Components integrations.
	ClientKey string `json:"clientKey"`
	// Description of the API credential.
	Description *string `json:"description,omitempty"`
	// Unique identifier of the API credential.
	Id string `json:"id"`
	// The password for the API credential that was created.
	Password string `json:"password"`
	// List of [roles](https://docs.adyen.com/development-resources/api-credentials#roles-1) for the API credential.
	Roles []string `json:"roles"`
	// The name of the [API credential](https://docs.adyen.com/development-resources/api-credentials), for example **ws@Company.TestCompany**.
	Username string `json:"username"`
}

// NewCreateCompanyApiCredentialResponse instantiates a new CreateCompanyApiCredentialResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCompanyApiCredentialResponse(active bool, allowedIpAddresses []string, apiKey string, associatedMerchantAccounts []string, clientKey string, id string, password string, roles []string, username string) *CreateCompanyApiCredentialResponse {
	this := CreateCompanyApiCredentialResponse{}
	this.Active = active
	this.AllowedIpAddresses = allowedIpAddresses
	this.ApiKey = apiKey
	this.AssociatedMerchantAccounts = associatedMerchantAccounts
	this.ClientKey = clientKey
	this.Id = id
	this.Password = password
	this.Roles = roles
	this.Username = username
	return &this
}

// NewCreateCompanyApiCredentialResponseWithDefaults instantiates a new CreateCompanyApiCredentialResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCompanyApiCredentialResponseWithDefaults() *CreateCompanyApiCredentialResponse {
	this := CreateCompanyApiCredentialResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CreateCompanyApiCredentialResponse) GetLinks() ApiCredentialLinks {
	if o == nil || o.Links == nil {
		var ret ApiCredentialLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCompanyApiCredentialResponse) GetLinksOk() (*ApiCredentialLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CreateCompanyApiCredentialResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ApiCredentialLinks and assigns it to the Links field.
func (o *CreateCompanyApiCredentialResponse) SetLinks(v ApiCredentialLinks) {
	o.Links = &v
}

// GetActive returns the Active field value
func (o *CreateCompanyApiCredentialResponse) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *CreateCompanyApiCredentialResponse) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *CreateCompanyApiCredentialResponse) SetActive(v bool) {
	o.Active = v
}

// GetAllowedIpAddresses returns the AllowedIpAddresses field value
func (o *CreateCompanyApiCredentialResponse) GetAllowedIpAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllowedIpAddresses
}

// GetAllowedIpAddressesOk returns a tuple with the AllowedIpAddresses field value
// and a boolean to check if the value has been set.
func (o *CreateCompanyApiCredentialResponse) GetAllowedIpAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedIpAddresses, true
}

// SetAllowedIpAddresses sets field value
func (o *CreateCompanyApiCredentialResponse) SetAllowedIpAddresses(v []string) {
	o.AllowedIpAddresses = v
}

// GetAllowedOrigins returns the AllowedOrigins field value if set, zero value otherwise.
func (o *CreateCompanyApiCredentialResponse) GetAllowedOrigins() []AllowedOrigin {
	if o == nil || o.AllowedOrigins == nil {
		var ret []AllowedOrigin
		return ret
	}
	return o.AllowedOrigins
}

// GetAllowedOriginsOk returns a tuple with the AllowedOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCompanyApiCredentialResponse) GetAllowedOriginsOk() ([]AllowedOrigin, bool) {
	if o == nil || o.AllowedOrigins == nil {
		return nil, false
	}
	return o.AllowedOrigins, true
}

// HasAllowedOrigins returns a boolean if a field has been set.
func (o *CreateCompanyApiCredentialResponse) HasAllowedOrigins() bool {
	if o != nil && o.AllowedOrigins != nil {
		return true
	}

	return false
}

// SetAllowedOrigins gets a reference to the given []AllowedOrigin and assigns it to the AllowedOrigins field.
func (o *CreateCompanyApiCredentialResponse) SetAllowedOrigins(v []AllowedOrigin) {
	o.AllowedOrigins = v
}

// GetApiKey returns the ApiKey field value
func (o *CreateCompanyApiCredentialResponse) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *CreateCompanyApiCredentialResponse) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *CreateCompanyApiCredentialResponse) SetApiKey(v string) {
	o.ApiKey = v
}

// GetAssociatedMerchantAccounts returns the AssociatedMerchantAccounts field value
func (o *CreateCompanyApiCredentialResponse) GetAssociatedMerchantAccounts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AssociatedMerchantAccounts
}

// GetAssociatedMerchantAccountsOk returns a tuple with the AssociatedMerchantAccounts field value
// and a boolean to check if the value has been set.
func (o *CreateCompanyApiCredentialResponse) GetAssociatedMerchantAccountsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssociatedMerchantAccounts, true
}

// SetAssociatedMerchantAccounts sets field value
func (o *CreateCompanyApiCredentialResponse) SetAssociatedMerchantAccounts(v []string) {
	o.AssociatedMerchantAccounts = v
}

// GetClientKey returns the ClientKey field value
func (o *CreateCompanyApiCredentialResponse) GetClientKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientKey
}

// GetClientKeyOk returns a tuple with the ClientKey field value
// and a boolean to check if the value has been set.
func (o *CreateCompanyApiCredentialResponse) GetClientKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientKey, true
}

// SetClientKey sets field value
func (o *CreateCompanyApiCredentialResponse) SetClientKey(v string) {
	o.ClientKey = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateCompanyApiCredentialResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCompanyApiCredentialResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateCompanyApiCredentialResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateCompanyApiCredentialResponse) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *CreateCompanyApiCredentialResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateCompanyApiCredentialResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateCompanyApiCredentialResponse) SetId(v string) {
	o.Id = v
}

// GetPassword returns the Password field value
func (o *CreateCompanyApiCredentialResponse) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CreateCompanyApiCredentialResponse) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CreateCompanyApiCredentialResponse) SetPassword(v string) {
	o.Password = v
}

// GetRoles returns the Roles field value
func (o *CreateCompanyApiCredentialResponse) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *CreateCompanyApiCredentialResponse) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *CreateCompanyApiCredentialResponse) SetRoles(v []string) {
	o.Roles = v
}

// GetUsername returns the Username field value
func (o *CreateCompanyApiCredentialResponse) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CreateCompanyApiCredentialResponse) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CreateCompanyApiCredentialResponse) SetUsername(v string) {
	o.Username = v
}

func (o CreateCompanyApiCredentialResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if true {
		toSerialize["active"] = o.Active
	}
	if true {
		toSerialize["allowedIpAddresses"] = o.AllowedIpAddresses
	}
	if o.AllowedOrigins != nil {
		toSerialize["allowedOrigins"] = o.AllowedOrigins
	}
	if true {
		toSerialize["apiKey"] = o.ApiKey
	}
	if true {
		toSerialize["associatedMerchantAccounts"] = o.AssociatedMerchantAccounts
	}
	if true {
		toSerialize["clientKey"] = o.ClientKey
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["roles"] = o.Roles
	}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCompanyApiCredentialResponse struct {
	value *CreateCompanyApiCredentialResponse
	isSet bool
}

func (v NullableCreateCompanyApiCredentialResponse) Get() *CreateCompanyApiCredentialResponse {
	return v.value
}

func (v *NullableCreateCompanyApiCredentialResponse) Set(val *CreateCompanyApiCredentialResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCompanyApiCredentialResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCompanyApiCredentialResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCompanyApiCredentialResponse(val *CreateCompanyApiCredentialResponse) *NullableCreateCompanyApiCredentialResponse {
	return &NullableCreateCompanyApiCredentialResponse{value: val, isSet: true}
}

func (v NullableCreateCompanyApiCredentialResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCompanyApiCredentialResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



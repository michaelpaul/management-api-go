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

// StoreCreationWithMerchantCodeRequest struct for StoreCreationWithMerchantCodeRequest
type StoreCreationWithMerchantCodeRequest struct {
	Address Address2 `json:"address"`
	// The unique identifiers of the [business lines](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/businesslines__resParam_id) that the store is associated with. If not specified, the business line of the merchant account is used. Required when there are multiple business lines under the merchant account.
	BusinessLineIds []string `json:"businessLineIds,omitempty"`
	// Your description of the store.
	Description string `json:"description"`
	// When using the Zip payment method: The location ID that Zip has assigned to your store.
	ExternalReferenceId *string `json:"externalReferenceId,omitempty"`
	// The unique identifier of the merchant account that the store belongs to.
	MerchantId string `json:"merchantId"`
	// The phone number of the store, including '+' and country code.
	PhoneNumber string `json:"phoneNumber"`
	// Your reference to recognize the store by. Also known as the store code.  Allowed characters: Lowercase and uppercase letters without diacritics, numbers 0 through 9, hyphen (-), and underscore (_).
	Reference *string `json:"reference,omitempty"`
	// The store name to be shown on the shopper's bank or credit card statement and on the shopper receipt. Maximum length: 22 characters; can't be all numbers.
	ShopperStatement string `json:"shopperStatement"`
	SplitConfiguration *StoreSplitConfiguration `json:"splitConfiguration,omitempty"`
}

// NewStoreCreationWithMerchantCodeRequest instantiates a new StoreCreationWithMerchantCodeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreCreationWithMerchantCodeRequest(address Address2, description string, merchantId string, phoneNumber string, shopperStatement string) *StoreCreationWithMerchantCodeRequest {
	this := StoreCreationWithMerchantCodeRequest{}
	this.Address = address
	this.Description = description
	this.MerchantId = merchantId
	this.PhoneNumber = phoneNumber
	this.ShopperStatement = shopperStatement
	return &this
}

// NewStoreCreationWithMerchantCodeRequestWithDefaults instantiates a new StoreCreationWithMerchantCodeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreCreationWithMerchantCodeRequestWithDefaults() *StoreCreationWithMerchantCodeRequest {
	this := StoreCreationWithMerchantCodeRequest{}
	return &this
}

// GetAddress returns the Address field value
func (o *StoreCreationWithMerchantCodeRequest) GetAddress() Address2 {
	if o == nil {
		var ret Address2
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *StoreCreationWithMerchantCodeRequest) GetAddressOk() (*Address2, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *StoreCreationWithMerchantCodeRequest) SetAddress(v Address2) {
	o.Address = v
}

// GetBusinessLineIds returns the BusinessLineIds field value if set, zero value otherwise.
func (o *StoreCreationWithMerchantCodeRequest) GetBusinessLineIds() []string {
	if o == nil || o.BusinessLineIds == nil {
		var ret []string
		return ret
	}
	return o.BusinessLineIds
}

// GetBusinessLineIdsOk returns a tuple with the BusinessLineIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreCreationWithMerchantCodeRequest) GetBusinessLineIdsOk() ([]string, bool) {
	if o == nil || o.BusinessLineIds == nil {
		return nil, false
	}
	return o.BusinessLineIds, true
}

// HasBusinessLineIds returns a boolean if a field has been set.
func (o *StoreCreationWithMerchantCodeRequest) HasBusinessLineIds() bool {
	if o != nil && o.BusinessLineIds != nil {
		return true
	}

	return false
}

// SetBusinessLineIds gets a reference to the given []string and assigns it to the BusinessLineIds field.
func (o *StoreCreationWithMerchantCodeRequest) SetBusinessLineIds(v []string) {
	o.BusinessLineIds = v
}

// GetDescription returns the Description field value
func (o *StoreCreationWithMerchantCodeRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *StoreCreationWithMerchantCodeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *StoreCreationWithMerchantCodeRequest) SetDescription(v string) {
	o.Description = v
}

// GetExternalReferenceId returns the ExternalReferenceId field value if set, zero value otherwise.
func (o *StoreCreationWithMerchantCodeRequest) GetExternalReferenceId() string {
	if o == nil || o.ExternalReferenceId == nil {
		var ret string
		return ret
	}
	return *o.ExternalReferenceId
}

// GetExternalReferenceIdOk returns a tuple with the ExternalReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreCreationWithMerchantCodeRequest) GetExternalReferenceIdOk() (*string, bool) {
	if o == nil || o.ExternalReferenceId == nil {
		return nil, false
	}
	return o.ExternalReferenceId, true
}

// HasExternalReferenceId returns a boolean if a field has been set.
func (o *StoreCreationWithMerchantCodeRequest) HasExternalReferenceId() bool {
	if o != nil && o.ExternalReferenceId != nil {
		return true
	}

	return false
}

// SetExternalReferenceId gets a reference to the given string and assigns it to the ExternalReferenceId field.
func (o *StoreCreationWithMerchantCodeRequest) SetExternalReferenceId(v string) {
	o.ExternalReferenceId = &v
}

// GetMerchantId returns the MerchantId field value
func (o *StoreCreationWithMerchantCodeRequest) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *StoreCreationWithMerchantCodeRequest) GetMerchantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *StoreCreationWithMerchantCodeRequest) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *StoreCreationWithMerchantCodeRequest) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *StoreCreationWithMerchantCodeRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *StoreCreationWithMerchantCodeRequest) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *StoreCreationWithMerchantCodeRequest) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreCreationWithMerchantCodeRequest) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *StoreCreationWithMerchantCodeRequest) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *StoreCreationWithMerchantCodeRequest) SetReference(v string) {
	o.Reference = &v
}

// GetShopperStatement returns the ShopperStatement field value
func (o *StoreCreationWithMerchantCodeRequest) GetShopperStatement() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShopperStatement
}

// GetShopperStatementOk returns a tuple with the ShopperStatement field value
// and a boolean to check if the value has been set.
func (o *StoreCreationWithMerchantCodeRequest) GetShopperStatementOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ShopperStatement, true
}

// SetShopperStatement sets field value
func (o *StoreCreationWithMerchantCodeRequest) SetShopperStatement(v string) {
	o.ShopperStatement = v
}

// GetSplitConfiguration returns the SplitConfiguration field value if set, zero value otherwise.
func (o *StoreCreationWithMerchantCodeRequest) GetSplitConfiguration() StoreSplitConfiguration {
	if o == nil || o.SplitConfiguration == nil {
		var ret StoreSplitConfiguration
		return ret
	}
	return *o.SplitConfiguration
}

// GetSplitConfigurationOk returns a tuple with the SplitConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreCreationWithMerchantCodeRequest) GetSplitConfigurationOk() (*StoreSplitConfiguration, bool) {
	if o == nil || o.SplitConfiguration == nil {
		return nil, false
	}
	return o.SplitConfiguration, true
}

// HasSplitConfiguration returns a boolean if a field has been set.
func (o *StoreCreationWithMerchantCodeRequest) HasSplitConfiguration() bool {
	if o != nil && o.SplitConfiguration != nil {
		return true
	}

	return false
}

// SetSplitConfiguration gets a reference to the given StoreSplitConfiguration and assigns it to the SplitConfiguration field.
func (o *StoreCreationWithMerchantCodeRequest) SetSplitConfiguration(v StoreSplitConfiguration) {
	o.SplitConfiguration = &v
}

func (o StoreCreationWithMerchantCodeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if o.BusinessLineIds != nil {
		toSerialize["businessLineIds"] = o.BusinessLineIds
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.ExternalReferenceId != nil {
		toSerialize["externalReferenceId"] = o.ExternalReferenceId
	}
	if true {
		toSerialize["merchantId"] = o.MerchantId
	}
	if true {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if true {
		toSerialize["shopperStatement"] = o.ShopperStatement
	}
	if o.SplitConfiguration != nil {
		toSerialize["splitConfiguration"] = o.SplitConfiguration
	}
	return json.Marshal(toSerialize)
}

type NullableStoreCreationWithMerchantCodeRequest struct {
	value *StoreCreationWithMerchantCodeRequest
	isSet bool
}

func (v NullableStoreCreationWithMerchantCodeRequest) Get() *StoreCreationWithMerchantCodeRequest {
	return v.value
}

func (v *NullableStoreCreationWithMerchantCodeRequest) Set(val *StoreCreationWithMerchantCodeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreCreationWithMerchantCodeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreCreationWithMerchantCodeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreCreationWithMerchantCodeRequest(val *StoreCreationWithMerchantCodeRequest) *NullableStoreCreationWithMerchantCodeRequest {
	return &NullableStoreCreationWithMerchantCodeRequest{value: val, isSet: true}
}

func (v NullableStoreCreationWithMerchantCodeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreCreationWithMerchantCodeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



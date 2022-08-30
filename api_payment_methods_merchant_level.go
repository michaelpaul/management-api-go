/*
Management API

Configure and manage your Adyen company and merchant accounts, stores, and payment terminals. ## Authentication Each request to the Management API must be signed with an API key. [Generate your API key](https://docs.adyen.com/development-resources/api-credentials#generate-api-key) in the Customer Area and then set this key to the `X-API-Key` header value.  To access the live endpoints, you need to generate a new API key in your live Customer Area. ## Versioning  Management API handles versioning as part of the endpoint URL. For example, to send a request to version 1 of the `/companies/{companyId}/webhooks` endpoint, use:  ```text https://management-test.adyen.com/v1/companies/{companyId}/webhooks ```

API version: 1
Contact: developer-experience@adyen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// PaymentMethodsMerchantLevelApiService PaymentMethodsMerchantLevelApi service
type PaymentMethodsMerchantLevelApiService service

type ApiGetMerchantsMerchantIdPaymentMethodSettingsRequest struct {
	ctx context.Context
	ApiService *PaymentMethodsMerchantLevelApiService
	merchantId string
	storeId *string
	businessLineId *string
	pageSize *int32
	pageNumber *int32
}

// The unique identifier of the store for which to return the payment methods.
func (r ApiGetMerchantsMerchantIdPaymentMethodSettingsRequest) StoreId(storeId string) ApiGetMerchantsMerchantIdPaymentMethodSettingsRequest {
	r.storeId = &storeId
	return r
}

// The unique identifier of the Business Line for which to return the payment methods.
func (r ApiGetMerchantsMerchantIdPaymentMethodSettingsRequest) BusinessLineId(businessLineId string) ApiGetMerchantsMerchantIdPaymentMethodSettingsRequest {
	r.businessLineId = &businessLineId
	return r
}

// The number of items to have on a page, maximum 100. The default is 10 items on a page.
func (r ApiGetMerchantsMerchantIdPaymentMethodSettingsRequest) PageSize(pageSize int32) ApiGetMerchantsMerchantIdPaymentMethodSettingsRequest {
	r.pageSize = &pageSize
	return r
}

// The number of the page to fetch.
func (r ApiGetMerchantsMerchantIdPaymentMethodSettingsRequest) PageNumber(pageNumber int32) ApiGetMerchantsMerchantIdPaymentMethodSettingsRequest {
	r.pageNumber = &pageNumber
	return r
}

func (r ApiGetMerchantsMerchantIdPaymentMethodSettingsRequest) Execute() (*PaymentMethodResponse, *http.Response, error) {
	return r.ApiService.GetMerchantsMerchantIdPaymentMethodSettingsExecute(r)
}

/*
GetMerchantsMerchantIdPaymentMethodSettings Get all payment methods

Returns details for all payment methods of the merchant account identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Payment methods read


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @return ApiGetMerchantsMerchantIdPaymentMethodSettingsRequest
*/
func (a *PaymentMethodsMerchantLevelApiService) GetMerchantsMerchantIdPaymentMethodSettings(ctx context.Context, merchantId string) ApiGetMerchantsMerchantIdPaymentMethodSettingsRequest {
	return ApiGetMerchantsMerchantIdPaymentMethodSettingsRequest{
		ApiService: a,
		ctx: ctx,
		merchantId: merchantId,
	}
}

// Execute executes the request
//  @return PaymentMethodResponse
func (a *PaymentMethodsMerchantLevelApiService) GetMerchantsMerchantIdPaymentMethodSettingsExecute(r ApiGetMerchantsMerchantIdPaymentMethodSettingsRequest) (*PaymentMethodResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentMethodResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentMethodsMerchantLevelApiService.GetMerchantsMerchantIdPaymentMethodSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchants/{merchantId}/paymentMethodSettings"
	localVarPath = strings.Replace(localVarPath, "{"+"merchantId"+"}", url.PathEscape(parameterToString(r.merchantId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.storeId != nil {
		localVarQueryParams.Add("storeId", parameterToString(*r.storeId, ""))
	}
	if r.businessLineId != nil {
		localVarQueryParams.Add("businessLineId", parameterToString(*r.businessLineId, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.pageNumber != nil {
		localVarQueryParams.Add("pageNumber", parameterToString(*r.pageNumber, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdRequest struct {
	ctx context.Context
	ApiService *PaymentMethodsMerchantLevelApiService
	merchantId string
	paymentMethodId string
}

func (r ApiGetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdRequest) Execute() (*PaymentMethod, *http.Response, error) {
	return r.ApiService.GetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdExecute(r)
}

/*
GetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId Get payment method details

Returns details for the merchant account and the payment method identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Payment methods read


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @param paymentMethodId The unique identifier of the payment method.
 @return ApiGetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdRequest
*/
func (a *PaymentMethodsMerchantLevelApiService) GetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId(ctx context.Context, merchantId string, paymentMethodId string) ApiGetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdRequest {
	return ApiGetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdRequest{
		ApiService: a,
		ctx: ctx,
		merchantId: merchantId,
		paymentMethodId: paymentMethodId,
	}
}

// Execute executes the request
//  @return PaymentMethod
func (a *PaymentMethodsMerchantLevelApiService) GetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdExecute(r ApiGetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdRequest) (*PaymentMethod, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentMethod
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentMethodsMerchantLevelApiService.GetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchants/{merchantId}/paymentMethodSettings/{paymentMethodId}"
	localVarPath = strings.Replace(localVarPath, "{"+"merchantId"+"}", url.PathEscape(parameterToString(r.merchantId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"paymentMethodId"+"}", url.PathEscape(parameterToString(r.paymentMethodId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdRequest struct {
	ctx context.Context
	ApiService *PaymentMethodsMerchantLevelApiService
	merchantId string
	paymentMethodId string
	updatePaymentMethodInfo *UpdatePaymentMethodInfo
}

func (r ApiPatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdRequest) UpdatePaymentMethodInfo(updatePaymentMethodInfo UpdatePaymentMethodInfo) ApiPatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdRequest {
	r.updatePaymentMethodInfo = &updatePaymentMethodInfo
	return r
}

func (r ApiPatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdRequest) Execute() (*PaymentMethod, *http.Response, error) {
	return r.ApiService.PatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdExecute(r)
}

/*
PatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId Update a payment method

Updates payment method details for the merchant account and the payment method identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Payment methods read and write


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @param paymentMethodId The unique identifier of the payment method.
 @return ApiPatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdRequest
*/
func (a *PaymentMethodsMerchantLevelApiService) PatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId(ctx context.Context, merchantId string, paymentMethodId string) ApiPatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdRequest {
	return ApiPatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdRequest{
		ApiService: a,
		ctx: ctx,
		merchantId: merchantId,
		paymentMethodId: paymentMethodId,
	}
}

// Execute executes the request
//  @return PaymentMethod
func (a *PaymentMethodsMerchantLevelApiService) PatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdExecute(r ApiPatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdRequest) (*PaymentMethod, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentMethod
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentMethodsMerchantLevelApiService.PatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchants/{merchantId}/paymentMethodSettings/{paymentMethodId}"
	localVarPath = strings.Replace(localVarPath, "{"+"merchantId"+"}", url.PathEscape(parameterToString(r.merchantId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"paymentMethodId"+"}", url.PathEscape(parameterToString(r.paymentMethodId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updatePaymentMethodInfo
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostMerchantsMerchantIdPaymentMethodSettingsRequest struct {
	ctx context.Context
	ApiService *PaymentMethodsMerchantLevelApiService
	merchantId string
	paymentMethodSetupInfo *PaymentMethodSetupInfo
}

func (r ApiPostMerchantsMerchantIdPaymentMethodSettingsRequest) PaymentMethodSetupInfo(paymentMethodSetupInfo PaymentMethodSetupInfo) ApiPostMerchantsMerchantIdPaymentMethodSettingsRequest {
	r.paymentMethodSetupInfo = &paymentMethodSetupInfo
	return r
}

func (r ApiPostMerchantsMerchantIdPaymentMethodSettingsRequest) Execute() (*PaymentMethod, *http.Response, error) {
	return r.ApiService.PostMerchantsMerchantIdPaymentMethodSettingsExecute(r)
}

/*
PostMerchantsMerchantIdPaymentMethodSettings Request a payment method

Sends a request to add a new payment method to the merchant account identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Payment methods read and write


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @return ApiPostMerchantsMerchantIdPaymentMethodSettingsRequest
*/
func (a *PaymentMethodsMerchantLevelApiService) PostMerchantsMerchantIdPaymentMethodSettings(ctx context.Context, merchantId string) ApiPostMerchantsMerchantIdPaymentMethodSettingsRequest {
	return ApiPostMerchantsMerchantIdPaymentMethodSettingsRequest{
		ApiService: a,
		ctx: ctx,
		merchantId: merchantId,
	}
}

// Execute executes the request
//  @return PaymentMethod
func (a *PaymentMethodsMerchantLevelApiService) PostMerchantsMerchantIdPaymentMethodSettingsExecute(r ApiPostMerchantsMerchantIdPaymentMethodSettingsRequest) (*PaymentMethod, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentMethod
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentMethodsMerchantLevelApiService.PostMerchantsMerchantIdPaymentMethodSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchants/{merchantId}/paymentMethodSettings"
	localVarPath = strings.Replace(localVarPath, "{"+"merchantId"+"}", url.PathEscape(parameterToString(r.merchantId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.paymentMethodSetupInfo
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

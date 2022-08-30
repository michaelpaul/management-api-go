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

// Linger please
var (
	_ context.Context
)

// APICredentialsCompanyLevelApiService APICredentialsCompanyLevelApi service
type APICredentialsCompanyLevelApiService service

type ApiGetCompaniesCompanyIdApiCredentialsRequest struct {
	ctx context.Context
	ApiService *APICredentialsCompanyLevelApiService
	companyId string
	pageNumber *int32
	pageSize *int32
}

// The number of the page to fetch.
func (r ApiGetCompaniesCompanyIdApiCredentialsRequest) PageNumber(pageNumber int32) ApiGetCompaniesCompanyIdApiCredentialsRequest {
	r.pageNumber = &pageNumber
	return r
}
// The number of items to have on a page, maximum 100. The default is 10 items on a page.
func (r ApiGetCompaniesCompanyIdApiCredentialsRequest) PageSize(pageSize int32) ApiGetCompaniesCompanyIdApiCredentialsRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetCompaniesCompanyIdApiCredentialsRequest) Execute() (*ListCompanyApiCredentialsResponse, *http.Response, error) {
	return r.ApiService.GetCompaniesCompanyIdApiCredentialsExecute(r)
}

/*
GetCompaniesCompanyIdApiCredentials Get a list of API credentials

Returns the list of [API credentials](https://docs.adyen.com/development-resources/api-credentials) for the company account. The list is grouped into pages as defined by the query parameters.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The unique identifier of the company account.
 @return ApiGetCompaniesCompanyIdApiCredentialsRequest
*/
func (a *APICredentialsCompanyLevelApiService) GetCompaniesCompanyIdApiCredentials(ctx context.Context, companyId string) ApiGetCompaniesCompanyIdApiCredentialsRequest {
	return ApiGetCompaniesCompanyIdApiCredentialsRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
	}
}

// Execute executes the request
//  @return ListCompanyApiCredentialsResponse
func (a *APICredentialsCompanyLevelApiService) GetCompaniesCompanyIdApiCredentialsExecute(r ApiGetCompaniesCompanyIdApiCredentialsRequest) (*ListCompanyApiCredentialsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListCompanyApiCredentialsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APICredentialsCompanyLevelApiService.GetCompaniesCompanyIdApiCredentials")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/companies/{companyId}/apiCredentials"
	localVarPath = strings.Replace(localVarPath, "{"+"companyId"+"}", url.PathEscape(parameterToString(r.companyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageNumber != nil {
		localVarQueryParams.Add("pageNumber", parameterToString(*r.pageNumber, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
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

type ApiGetCompaniesCompanyIdApiCredentialsApiCredentialIdRequest struct {
	ctx context.Context
	ApiService *APICredentialsCompanyLevelApiService
	companyId string
	apiCredentialId string
}


func (r ApiGetCompaniesCompanyIdApiCredentialsApiCredentialIdRequest) Execute() (*CompanyApiCredential, *http.Response, error) {
	return r.ApiService.GetCompaniesCompanyIdApiCredentialsApiCredentialIdExecute(r)
}

/*
GetCompaniesCompanyIdApiCredentialsApiCredentialId Get an API credential

Returns the [API credential](https://docs.adyen.com/development-resources/api-credentials) identified in the path.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The unique identifier of the company account.
 @param apiCredentialId Unique identifier of the API credential.
 @return ApiGetCompaniesCompanyIdApiCredentialsApiCredentialIdRequest
*/
func (a *APICredentialsCompanyLevelApiService) GetCompaniesCompanyIdApiCredentialsApiCredentialId(ctx context.Context, companyId string, apiCredentialId string) ApiGetCompaniesCompanyIdApiCredentialsApiCredentialIdRequest {
	return ApiGetCompaniesCompanyIdApiCredentialsApiCredentialIdRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
		apiCredentialId: apiCredentialId,
	}
}

// Execute executes the request
//  @return CompanyApiCredential
func (a *APICredentialsCompanyLevelApiService) GetCompaniesCompanyIdApiCredentialsApiCredentialIdExecute(r ApiGetCompaniesCompanyIdApiCredentialsApiCredentialIdRequest) (*CompanyApiCredential, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CompanyApiCredential
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APICredentialsCompanyLevelApiService.GetCompaniesCompanyIdApiCredentialsApiCredentialId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/companies/{companyId}/apiCredentials/{apiCredentialId}"
	localVarPath = strings.Replace(localVarPath, "{"+"companyId"+"}", url.PathEscape(parameterToString(r.companyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"apiCredentialId"+"}", url.PathEscape(parameterToString(r.apiCredentialId, "")), -1)

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

type ApiPatchCompaniesCompanyIdApiCredentialsApiCredentialIdRequest struct {
	ctx context.Context
	ApiService *APICredentialsCompanyLevelApiService
	companyId string
	apiCredentialId string
	updateCompanyApiCredentialRequest *UpdateCompanyApiCredentialRequest
}

func (r ApiPatchCompaniesCompanyIdApiCredentialsApiCredentialIdRequest) UpdateCompanyApiCredentialRequest(updateCompanyApiCredentialRequest UpdateCompanyApiCredentialRequest) ApiPatchCompaniesCompanyIdApiCredentialsApiCredentialIdRequest {
	r.updateCompanyApiCredentialRequest = &updateCompanyApiCredentialRequest
	return r
}

func (r ApiPatchCompaniesCompanyIdApiCredentialsApiCredentialIdRequest) Execute() (*CompanyApiCredential, *http.Response, error) {
	return r.ApiService.PatchCompaniesCompanyIdApiCredentialsApiCredentialIdExecute(r)
}

/*
PatchCompaniesCompanyIdApiCredentialsApiCredentialId Update an API credential.

Changes the API credential's roles, merchant account access, or allowed origins. The request has the new values for the fields you want to change. The response contains the full updated API credential, including the new values from the request. 

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The unique identifier of the company account.
 @param apiCredentialId Unique identifier of the API credential.
 @return ApiPatchCompaniesCompanyIdApiCredentialsApiCredentialIdRequest
*/
func (a *APICredentialsCompanyLevelApiService) PatchCompaniesCompanyIdApiCredentialsApiCredentialId(ctx context.Context, companyId string, apiCredentialId string) ApiPatchCompaniesCompanyIdApiCredentialsApiCredentialIdRequest {
	return ApiPatchCompaniesCompanyIdApiCredentialsApiCredentialIdRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
		apiCredentialId: apiCredentialId,
	}
}

// Execute executes the request
//  @return CompanyApiCredential
func (a *APICredentialsCompanyLevelApiService) PatchCompaniesCompanyIdApiCredentialsApiCredentialIdExecute(r ApiPatchCompaniesCompanyIdApiCredentialsApiCredentialIdRequest) (*CompanyApiCredential, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CompanyApiCredential
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APICredentialsCompanyLevelApiService.PatchCompaniesCompanyIdApiCredentialsApiCredentialId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/companies/{companyId}/apiCredentials/{apiCredentialId}"
	localVarPath = strings.Replace(localVarPath, "{"+"companyId"+"}", url.PathEscape(parameterToString(r.companyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"apiCredentialId"+"}", url.PathEscape(parameterToString(r.apiCredentialId, "")), -1)

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
	localVarPostBody = r.updateCompanyApiCredentialRequest
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

type ApiPostCompaniesCompanyIdApiCredentialsRequest struct {
	ctx context.Context
	ApiService *APICredentialsCompanyLevelApiService
	companyId string
	createCompanyApiCredentialRequest *CreateCompanyApiCredentialRequest
}

func (r ApiPostCompaniesCompanyIdApiCredentialsRequest) CreateCompanyApiCredentialRequest(createCompanyApiCredentialRequest CreateCompanyApiCredentialRequest) ApiPostCompaniesCompanyIdApiCredentialsRequest {
	r.createCompanyApiCredentialRequest = &createCompanyApiCredentialRequest
	return r
}

func (r ApiPostCompaniesCompanyIdApiCredentialsRequest) Execute() (*CreateCompanyApiCredentialResponse, *http.Response, error) {
	return r.ApiService.PostCompaniesCompanyIdApiCredentialsExecute(r)
}

/*
PostCompaniesCompanyIdApiCredentials Create an API credential.

Creates an [API credential](https://docs.adyen.com/development-resources/api-credentials) for the company account identified in the path. In the request, you can specify which merchant accounts the new API credential will have access to, as well as its roles and allowed origins.

The response includes several types of authentication details:
* [API key](https://docs.adyen.com/development-resources/api-authentication#api-key-authentication): used for API request authentication.
* [Client key](https://docs.adyen.com/development-resources/client-side-authentication#how-it-works): public key used for client-side authentication.
* [Username and password](https://docs.adyen.com/development-resources/api-authentication#using-basic-authentication): used for basic authentication.

> Make sure you store the API key securely in your system. You won't be able to retrieve it later.

If your API key is lost or compromised, you need to [generate a new API key](https://docs.adyen.com/api-explorer/#/ManagementService/v1/post/companies/{companyId}/apiCredentials/{apiCredentialId}/generateApiKey).

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The unique identifier of the company account.
 @return ApiPostCompaniesCompanyIdApiCredentialsRequest
*/
func (a *APICredentialsCompanyLevelApiService) PostCompaniesCompanyIdApiCredentials(ctx context.Context, companyId string) ApiPostCompaniesCompanyIdApiCredentialsRequest {
	return ApiPostCompaniesCompanyIdApiCredentialsRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
	}
}

// Execute executes the request
//  @return CreateCompanyApiCredentialResponse
func (a *APICredentialsCompanyLevelApiService) PostCompaniesCompanyIdApiCredentialsExecute(r ApiPostCompaniesCompanyIdApiCredentialsRequest) (*CreateCompanyApiCredentialResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateCompanyApiCredentialResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APICredentialsCompanyLevelApiService.PostCompaniesCompanyIdApiCredentials")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/companies/{companyId}/apiCredentials"
	localVarPath = strings.Replace(localVarPath, "{"+"companyId"+"}", url.PathEscape(parameterToString(r.companyId, "")), -1)

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
	localVarPostBody = r.createCompanyApiCredentialRequest
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
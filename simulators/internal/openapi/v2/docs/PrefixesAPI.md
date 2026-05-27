# \PrefixesAPI

All URIs are relative to *https://eric-pc-routing-information-base:443/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeregisterPrefix**](PrefixesAPI.md#DeregisterPrefix) | **Delete** /network-instances/{networkInstanceName}/producers/{producerId}/prefixes | 
[**ReceivePrefixRegistrationEOF**](PrefixesAPI.md#ReceivePrefixRegistrationEOF) | **Post** /network-instances/{networkInstanceName}/producers/{producerId}/prefix-reg-eof | 
[**RegisterPrefix**](PrefixesAPI.md#RegisterPrefix) | **Post** /network-instances/{networkInstanceName}/producers/{producerId}/prefixes | 



## DeregisterPrefix

> DeregisterPrefix(ctx, networkInstanceName, producerId).PrefixRegister(prefixRegister).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkInstanceName := "networkInstanceName_example" // string | networkInstance to which the prefix belongs to.
	producerId := "producerId_example" // string | producer name, please refer to the clientIdentifier provided by producer when it register to RIB.
	prefixRegister := *openapiclient.NewPrefixRegister() // PrefixRegister | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PrefixesAPI.DeregisterPrefix(context.Background(), networkInstanceName, producerId).PrefixRegister(prefixRegister).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrefixesAPI.DeregisterPrefix``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInstanceName** | **string** | networkInstance to which the prefix belongs to. | 
**producerId** | **string** | producer name, please refer to the clientIdentifier provided by producer when it register to RIB. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeregisterPrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **prefixRegister** | [**PrefixRegister**](PrefixRegister.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReceivePrefixRegistrationEOF

> ReceivePrefixRegistrationEOF(ctx, networkInstanceName, producerId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkInstanceName := "networkInstanceName_example" // string | networkInstance to which the EOF belongs to.
	producerId := "producerId_example" // string | producer name, please refer to the clientIdentifier provided by producer when it register to RIB.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PrefixesAPI.ReceivePrefixRegistrationEOF(context.Background(), networkInstanceName, producerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrefixesAPI.ReceivePrefixRegistrationEOF``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInstanceName** | **string** | networkInstance to which the EOF belongs to. | 
**producerId** | **string** | producer name, please refer to the clientIdentifier provided by producer when it register to RIB. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReceivePrefixRegistrationEOFRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterPrefix

> RegisterPrefix(ctx, networkInstanceName, producerId).PrefixRegister(prefixRegister).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkInstanceName := "networkInstanceName_example" // string | networkInstance to which the prefix belongs to.
	producerId := "producerId_example" // string | producer name, please refer to the clientIdentifier provided by producer when it register to RIB.
	prefixRegister := *openapiclient.NewPrefixRegister() // PrefixRegister | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PrefixesAPI.RegisterPrefix(context.Background(), networkInstanceName, producerId).PrefixRegister(prefixRegister).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrefixesAPI.RegisterPrefix``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInstanceName** | **string** | networkInstance to which the prefix belongs to. | 
**producerId** | **string** | producer name, please refer to the clientIdentifier provided by producer when it register to RIB. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterPrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **prefixRegister** | [**PrefixRegister**](PrefixRegister.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


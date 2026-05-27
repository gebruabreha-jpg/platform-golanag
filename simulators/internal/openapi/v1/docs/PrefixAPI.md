# \PrefixAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeregisterPrefix**](PrefixAPI.md#DeregisterPrefix) | **Delete** /v1/network-instances/{network-instance-name}/prefixes | 
[**RegisterPrefix**](PrefixAPI.md#RegisterPrefix) | **Post** /v1/network-instances/{network-instance-name}/prefixes | 



## DeregisterPrefix

> DeregisterPrefix(ctx, networkInstanceName).PrefixRegister(prefixRegister).Execute()





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
	networkInstanceName := "networkInstanceName_example" // string | network-instance to which the prefix belongs to.
	prefixRegister := *openapiclient.NewPrefixRegister() // PrefixRegister | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PrefixAPI.DeregisterPrefix(context.Background(), networkInstanceName).PrefixRegister(prefixRegister).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrefixAPI.DeregisterPrefix``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInstanceName** | **string** | network-instance to which the prefix belongs to. | 

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


## RegisterPrefix

> RegisterPrefix(ctx, networkInstanceName).PrefixRegister(prefixRegister).Execute()





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
	networkInstanceName := "networkInstanceName_example" // string | network-instance to which the prefix belongs to.
	prefixRegister := *openapiclient.NewPrefixRegister() // PrefixRegister | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PrefixAPI.RegisterPrefix(context.Background(), networkInstanceName).PrefixRegister(prefixRegister).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrefixAPI.RegisterPrefix``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInstanceName** | **string** | network-instance to which the prefix belongs to. | 

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


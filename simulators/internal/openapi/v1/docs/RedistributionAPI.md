# \RedistributionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRouteRedistribution**](RedistributionAPI.md#AddRouteRedistribution) | **Post** /v1/network-instances/{network-instance-name}/redistribution | 
[**DeleteRouteRedistribution**](RedistributionAPI.md#DeleteRouteRedistribution) | **Delete** /v1/network-instances/{network-instance-name}/redistribution | 
[**UpdateRouteRedistribution**](RedistributionAPI.md#UpdateRouteRedistribution) | **Patch** /v1/network-instances/{network-instance-name}/redistribution | 



## AddRouteRedistribution

> AddRouteRedistribution(ctx, networkInstanceName).Redistribution(redistribution).Execute()





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
	networkInstanceName := "networkInstanceName_example" // string | network-instance to which the configuration belongs to.
	redistribution := *openapiclient.NewRedistribution() // Redistribution | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RedistributionAPI.AddRouteRedistribution(context.Background(), networkInstanceName).Redistribution(redistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedistributionAPI.AddRouteRedistribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInstanceName** | **string** | network-instance to which the configuration belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRouteRedistributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **redistribution** | [**Redistribution**](Redistribution.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRouteRedistribution

> DeleteRouteRedistribution(ctx, networkInstanceName).Redistribution(redistribution).Execute()





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
	networkInstanceName := "networkInstanceName_example" // string | network-instance to which the configuration belongs to.
	redistribution := *openapiclient.NewRedistribution() // Redistribution | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RedistributionAPI.DeleteRouteRedistribution(context.Background(), networkInstanceName).Redistribution(redistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedistributionAPI.DeleteRouteRedistribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInstanceName** | **string** | network-instance to which the configuration belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteRedistributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **redistribution** | [**Redistribution**](Redistribution.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRouteRedistribution

> UpdateRouteRedistribution(ctx, networkInstanceName).Redistribution(redistribution).Execute()





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
	networkInstanceName := "networkInstanceName_example" // string | network-instance to which the configuration belongs to.
	redistribution := *openapiclient.NewRedistribution() // Redistribution | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RedistributionAPI.UpdateRouteRedistribution(context.Background(), networkInstanceName).Redistribution(redistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedistributionAPI.UpdateRouteRedistribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInstanceName** | **string** | network-instance to which the configuration belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouteRedistributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **redistribution** | [**Redistribution**](Redistribution.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


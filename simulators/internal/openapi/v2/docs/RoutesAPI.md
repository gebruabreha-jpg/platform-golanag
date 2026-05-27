# \RoutesAPI

All URIs are relative to *https://eric-pc-routing-information-base:443/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteRoutes**](RoutesAPI.md#BulkDeleteRoutes) | **Delete** /network-instances/{networkInstanceName}/producers/{producerId}/routes | 
[**ReceiveRouteEOF**](RoutesAPI.md#ReceiveRouteEOF) | **Post** /network-instances/{networkInstanceName}/producers/{producerId}/route-eof | 
[**RouteOperation**](RoutesAPI.md#RouteOperation) | **Post** /network-instances/{networkInstanceName}/producers/{producerId}/routes | 



## BulkDeleteRoutes

> BulkDeleteRoutes(ctx, networkInstanceName, producerId).Execute()





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
	networkInstanceName := "networkInstanceName_example" // string | name of networkInstance that the routes belongs to.
	producerId := "producerId_example" // string | producer name, please refer to the clientIdentifier provided by producer when it register to RIB.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoutesAPI.BulkDeleteRoutes(context.Background(), networkInstanceName, producerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.BulkDeleteRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInstanceName** | **string** | name of networkInstance that the routes belongs to. | 
**producerId** | **string** | producer name, please refer to the clientIdentifier provided by producer when it register to RIB. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteRoutesRequest struct via the builder pattern


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


## ReceiveRouteEOF

> ReceiveRouteEOF(ctx, networkInstanceName, producerId).RouteEoFFromProducer(routeEoFFromProducer).Execute()





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
	routeEoFFromProducer := *openapiclient.NewRouteEoFFromProducer() // RouteEoFFromProducer | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoutesAPI.ReceiveRouteEOF(context.Background(), networkInstanceName, producerId).RouteEoFFromProducer(routeEoFFromProducer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.ReceiveRouteEOF``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReceiveRouteEOFRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **routeEoFFromProducer** | [**RouteEoFFromProducer**](RouteEoFFromProducer.md) |  | 

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


## RouteOperation

> RouteOperation(ctx, networkInstanceName, producerId).Routes(routes).Execute()





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
	networkInstanceName := "networkInstanceName_example" // string | name of network instance that the routes belongs to.
	producerId := "producerId_example" // string | producer name, please refer to the clientIdentifier provided by producer when it register to RIB.
	routes := *openapiclient.NewRoutes() // Routes | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoutesAPI.RouteOperation(context.Background(), networkInstanceName, producerId).Routes(routes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.RouteOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInstanceName** | **string** | name of network instance that the routes belongs to. | 
**producerId** | **string** | producer name, please refer to the clientIdentifier provided by producer when it register to RIB. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRouteOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **routes** | [**Routes**](Routes.md) |  | 

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


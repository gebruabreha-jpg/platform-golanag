# \RouteEofAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReceiveRouteEof**](RouteEofAPI.md#ReceiveRouteEof) | **Post** /v1/network-instances/{network-instance-name}/clients/{client-name}/route-eof | 



## ReceiveRouteEof

> ReceiveRouteEof(ctx, networkInstanceName, clientName).RouteEoFFromProducer(routeEoFFromProducer).Execute()





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
	networkInstanceName := "networkInstanceName_example" // string | network-instance to which the eof belongs to.
	clientName := "clientName_example" // string | client name, please refer to the client_identifier provided by client when it register to RIB.
	routeEoFFromProducer := *openapiclient.NewRouteEoFFromProducer() // RouteEoFFromProducer | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RouteEofAPI.ReceiveRouteEof(context.Background(), networkInstanceName, clientName).RouteEoFFromProducer(routeEoFFromProducer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteEofAPI.ReceiveRouteEof``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInstanceName** | **string** | network-instance to which the eof belongs to. | 
**clientName** | **string** | client name, please refer to the client_identifier provided by client when it register to RIB. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReceiveRouteEofRequest struct via the builder pattern


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


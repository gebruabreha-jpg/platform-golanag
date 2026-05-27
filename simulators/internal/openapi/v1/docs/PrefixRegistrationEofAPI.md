# \PrefixRegistrationEofAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReceivePrefixRegistrationEof**](PrefixRegistrationEofAPI.md#ReceivePrefixRegistrationEof) | **Post** /v1/network-instances/{network-instance-name}/clients/{client-name}/prefix-reg-eof | 



## ReceivePrefixRegistrationEof

> ReceivePrefixRegistrationEof(ctx, networkInstanceName, clientName).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PrefixRegistrationEofAPI.ReceivePrefixRegistrationEof(context.Background(), networkInstanceName, clientName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrefixRegistrationEofAPI.ReceivePrefixRegistrationEof``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReceivePrefixRegistrationEofRequest struct via the builder pattern


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


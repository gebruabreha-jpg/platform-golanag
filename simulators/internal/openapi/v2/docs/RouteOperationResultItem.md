# RouteOperationResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInstance** | **string** | Network instance. | 
**Prefix** | **string** | A valid IP prefix. | 
**AddressFamily** | [**AddressFamily**](AddressFamily.md) |  | 
**Action** | [**ConsumerRoutesNotificationItemAction**](ConsumerRoutesNotificationItemAction.md) |  | 
**Lsp** | Pointer to [**RouteLspType**](RouteLspType.md) | Route lsp flag. | [optional] 
**Status** | **int32** | http status code, such as 2xx/3xx/4xx/5xx etc.. | 

## Methods

### NewRouteOperationResultItem

`func NewRouteOperationResultItem(networkInstance string, prefix string, addressFamily AddressFamily, action ConsumerRoutesNotificationItemAction, status int32, ) *RouteOperationResultItem`

NewRouteOperationResultItem instantiates a new RouteOperationResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteOperationResultItemWithDefaults

`func NewRouteOperationResultItemWithDefaults() *RouteOperationResultItem`

NewRouteOperationResultItemWithDefaults instantiates a new RouteOperationResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkInstance

`func (o *RouteOperationResultItem) GetNetworkInstance() string`

GetNetworkInstance returns the NetworkInstance field if non-nil, zero value otherwise.

### GetNetworkInstanceOk

`func (o *RouteOperationResultItem) GetNetworkInstanceOk() (*string, bool)`

GetNetworkInstanceOk returns a tuple with the NetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstance

`func (o *RouteOperationResultItem) SetNetworkInstance(v string)`

SetNetworkInstance sets NetworkInstance field to given value.


### GetPrefix

`func (o *RouteOperationResultItem) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *RouteOperationResultItem) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *RouteOperationResultItem) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetAddressFamily

`func (o *RouteOperationResultItem) GetAddressFamily() AddressFamily`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *RouteOperationResultItem) GetAddressFamilyOk() (*AddressFamily, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *RouteOperationResultItem) SetAddressFamily(v AddressFamily)`

SetAddressFamily sets AddressFamily field to given value.


### GetAction

`func (o *RouteOperationResultItem) GetAction() ConsumerRoutesNotificationItemAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RouteOperationResultItem) GetActionOk() (*ConsumerRoutesNotificationItemAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RouteOperationResultItem) SetAction(v ConsumerRoutesNotificationItemAction)`

SetAction sets Action field to given value.


### GetLsp

`func (o *RouteOperationResultItem) GetLsp() RouteLspType`

GetLsp returns the Lsp field if non-nil, zero value otherwise.

### GetLspOk

`func (o *RouteOperationResultItem) GetLspOk() (*RouteLspType, bool)`

GetLspOk returns a tuple with the Lsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLsp

`func (o *RouteOperationResultItem) SetLsp(v RouteLspType)`

SetLsp sets Lsp field to given value.

### HasLsp

`func (o *RouteOperationResultItem) HasLsp() bool`

HasLsp returns a boolean if a field has been set.

### GetStatus

`func (o *RouteOperationResultItem) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RouteOperationResultItem) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RouteOperationResultItem) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



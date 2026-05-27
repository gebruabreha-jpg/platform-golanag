# RouteOperationResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInstance** | **string** | Network instance. | 
**Prefix** | [**IpPrefix**](IpPrefix.md) |  | 
**Afi** | **string** | Address family. | 
**Action** | **string** | Operation action. | 
**Lsp** | Pointer to [**RouteLspType**](RouteLspType.md) |  | [optional] 
**Status** | **int32** | http status code, such as 2xx/3xx/4xx/5xx etc.. | 

## Methods

### NewRouteOperationResultItem

`func NewRouteOperationResultItem(networkInstance string, prefix IpPrefix, afi string, action string, status int32, ) *RouteOperationResultItem`

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

`func (o *RouteOperationResultItem) GetPrefix() IpPrefix`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *RouteOperationResultItem) GetPrefixOk() (*IpPrefix, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *RouteOperationResultItem) SetPrefix(v IpPrefix)`

SetPrefix sets Prefix field to given value.


### GetAfi

`func (o *RouteOperationResultItem) GetAfi() string`

GetAfi returns the Afi field if non-nil, zero value otherwise.

### GetAfiOk

`func (o *RouteOperationResultItem) GetAfiOk() (*string, bool)`

GetAfiOk returns a tuple with the Afi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfi

`func (o *RouteOperationResultItem) SetAfi(v string)`

SetAfi sets Afi field to given value.


### GetAction

`func (o *RouteOperationResultItem) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RouteOperationResultItem) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RouteOperationResultItem) SetAction(v string)`

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



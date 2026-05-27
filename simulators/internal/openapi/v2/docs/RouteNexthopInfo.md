# RouteNexthopInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInstance** | Pointer to **string** | Network instance. | [optional] 
**Attr** | Pointer to [**[]RouteNexthopInfoAttrInner**](RouteNexthopInfoAttrInner.md) | Attributes of the nexthop of the route. Some attributes makes others mandatory:   - vpn_label: the vpn_label of labels field must be filled   - label: the label of labels field must be filled   - service_local: the addressInfo must be filled with a service type | [optional] 
**AddressInfo** | [**NexthopAddressInfo**](NexthopAddressInfo.md) |  | 
**InterfaceInfo** | Pointer to [**InterfaceInfo**](InterfaceInfo.md) | Interface info. | [optional] 
**Labels** | Pointer to [**NexthopLabels**](NexthopLabels.md) | label info. | [optional] 
**Priority** | Pointer to **int32** | Nexthop priority. The higher the value, the lower the priority | [optional] 

## Methods

### NewRouteNexthopInfo

`func NewRouteNexthopInfo(addressInfo NexthopAddressInfo, ) *RouteNexthopInfo`

NewRouteNexthopInfo instantiates a new RouteNexthopInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteNexthopInfoWithDefaults

`func NewRouteNexthopInfoWithDefaults() *RouteNexthopInfo`

NewRouteNexthopInfoWithDefaults instantiates a new RouteNexthopInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkInstance

`func (o *RouteNexthopInfo) GetNetworkInstance() string`

GetNetworkInstance returns the NetworkInstance field if non-nil, zero value otherwise.

### GetNetworkInstanceOk

`func (o *RouteNexthopInfo) GetNetworkInstanceOk() (*string, bool)`

GetNetworkInstanceOk returns a tuple with the NetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstance

`func (o *RouteNexthopInfo) SetNetworkInstance(v string)`

SetNetworkInstance sets NetworkInstance field to given value.

### HasNetworkInstance

`func (o *RouteNexthopInfo) HasNetworkInstance() bool`

HasNetworkInstance returns a boolean if a field has been set.

### GetAttr

`func (o *RouteNexthopInfo) GetAttr() []RouteNexthopInfoAttrInner`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *RouteNexthopInfo) GetAttrOk() (*[]RouteNexthopInfoAttrInner, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *RouteNexthopInfo) SetAttr(v []RouteNexthopInfoAttrInner)`

SetAttr sets Attr field to given value.

### HasAttr

`func (o *RouteNexthopInfo) HasAttr() bool`

HasAttr returns a boolean if a field has been set.

### GetAddressInfo

`func (o *RouteNexthopInfo) GetAddressInfo() NexthopAddressInfo`

GetAddressInfo returns the AddressInfo field if non-nil, zero value otherwise.

### GetAddressInfoOk

`func (o *RouteNexthopInfo) GetAddressInfoOk() (*NexthopAddressInfo, bool)`

GetAddressInfoOk returns a tuple with the AddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressInfo

`func (o *RouteNexthopInfo) SetAddressInfo(v NexthopAddressInfo)`

SetAddressInfo sets AddressInfo field to given value.


### GetInterfaceInfo

`func (o *RouteNexthopInfo) GetInterfaceInfo() InterfaceInfo`

GetInterfaceInfo returns the InterfaceInfo field if non-nil, zero value otherwise.

### GetInterfaceInfoOk

`func (o *RouteNexthopInfo) GetInterfaceInfoOk() (*InterfaceInfo, bool)`

GetInterfaceInfoOk returns a tuple with the InterfaceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceInfo

`func (o *RouteNexthopInfo) SetInterfaceInfo(v InterfaceInfo)`

SetInterfaceInfo sets InterfaceInfo field to given value.

### HasInterfaceInfo

`func (o *RouteNexthopInfo) HasInterfaceInfo() bool`

HasInterfaceInfo returns a boolean if a field has been set.

### GetLabels

`func (o *RouteNexthopInfo) GetLabels() NexthopLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RouteNexthopInfo) GetLabelsOk() (*NexthopLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RouteNexthopInfo) SetLabels(v NexthopLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RouteNexthopInfo) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetPriority

`func (o *RouteNexthopInfo) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RouteNexthopInfo) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RouteNexthopInfo) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RouteNexthopInfo) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# QueryNexthopNotificationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | A valid IP address. | 
**AddressFamily** | [**AddressFamily**](AddressFamily.md) |  | 
**NetworkInstance** | Pointer to **string** |  | [optional] 
**ProtocolType** | Pointer to [**RouteType**](RouteType.md) |  | [optional] 
**Flags** | Pointer to [**[]QueryNexthopNotificationItemAllOfFlags**](QueryNexthopNotificationItemAllOfFlags.md) |  | [optional] 
**ViaAddressInfo** | Pointer to [**NexthopAddressInfo**](NexthopAddressInfo.md) | The address of the route nexthop to the queried address | [optional] 
**Metric** | Pointer to **int32** |  | [optional] 
**ViaNetworkInstance** | Pointer to **string** |  | [optional] 
**NexthopTimestamp** | Pointer to **time.Time** |  | [optional] 
**Priority** | Pointer to **int32** | Nexthop priority. The higher the value, the lower the priority | [optional] 

## Methods

### NewQueryNexthopNotificationItem

`func NewQueryNexthopNotificationItem(address string, addressFamily AddressFamily, ) *QueryNexthopNotificationItem`

NewQueryNexthopNotificationItem instantiates a new QueryNexthopNotificationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryNexthopNotificationItemWithDefaults

`func NewQueryNexthopNotificationItemWithDefaults() *QueryNexthopNotificationItem`

NewQueryNexthopNotificationItemWithDefaults instantiates a new QueryNexthopNotificationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *QueryNexthopNotificationItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *QueryNexthopNotificationItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *QueryNexthopNotificationItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAddressFamily

`func (o *QueryNexthopNotificationItem) GetAddressFamily() AddressFamily`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *QueryNexthopNotificationItem) GetAddressFamilyOk() (*AddressFamily, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *QueryNexthopNotificationItem) SetAddressFamily(v AddressFamily)`

SetAddressFamily sets AddressFamily field to given value.


### GetNetworkInstance

`func (o *QueryNexthopNotificationItem) GetNetworkInstance() string`

GetNetworkInstance returns the NetworkInstance field if non-nil, zero value otherwise.

### GetNetworkInstanceOk

`func (o *QueryNexthopNotificationItem) GetNetworkInstanceOk() (*string, bool)`

GetNetworkInstanceOk returns a tuple with the NetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstance

`func (o *QueryNexthopNotificationItem) SetNetworkInstance(v string)`

SetNetworkInstance sets NetworkInstance field to given value.

### HasNetworkInstance

`func (o *QueryNexthopNotificationItem) HasNetworkInstance() bool`

HasNetworkInstance returns a boolean if a field has been set.

### GetProtocolType

`func (o *QueryNexthopNotificationItem) GetProtocolType() RouteType`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *QueryNexthopNotificationItem) GetProtocolTypeOk() (*RouteType, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *QueryNexthopNotificationItem) SetProtocolType(v RouteType)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *QueryNexthopNotificationItem) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetFlags

`func (o *QueryNexthopNotificationItem) GetFlags() []QueryNexthopNotificationItemAllOfFlags`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *QueryNexthopNotificationItem) GetFlagsOk() (*[]QueryNexthopNotificationItemAllOfFlags, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *QueryNexthopNotificationItem) SetFlags(v []QueryNexthopNotificationItemAllOfFlags)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *QueryNexthopNotificationItem) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetViaAddressInfo

`func (o *QueryNexthopNotificationItem) GetViaAddressInfo() NexthopAddressInfo`

GetViaAddressInfo returns the ViaAddressInfo field if non-nil, zero value otherwise.

### GetViaAddressInfoOk

`func (o *QueryNexthopNotificationItem) GetViaAddressInfoOk() (*NexthopAddressInfo, bool)`

GetViaAddressInfoOk returns a tuple with the ViaAddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaAddressInfo

`func (o *QueryNexthopNotificationItem) SetViaAddressInfo(v NexthopAddressInfo)`

SetViaAddressInfo sets ViaAddressInfo field to given value.

### HasViaAddressInfo

`func (o *QueryNexthopNotificationItem) HasViaAddressInfo() bool`

HasViaAddressInfo returns a boolean if a field has been set.

### GetMetric

`func (o *QueryNexthopNotificationItem) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *QueryNexthopNotificationItem) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *QueryNexthopNotificationItem) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *QueryNexthopNotificationItem) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetViaNetworkInstance

`func (o *QueryNexthopNotificationItem) GetViaNetworkInstance() string`

GetViaNetworkInstance returns the ViaNetworkInstance field if non-nil, zero value otherwise.

### GetViaNetworkInstanceOk

`func (o *QueryNexthopNotificationItem) GetViaNetworkInstanceOk() (*string, bool)`

GetViaNetworkInstanceOk returns a tuple with the ViaNetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaNetworkInstance

`func (o *QueryNexthopNotificationItem) SetViaNetworkInstance(v string)`

SetViaNetworkInstance sets ViaNetworkInstance field to given value.

### HasViaNetworkInstance

`func (o *QueryNexthopNotificationItem) HasViaNetworkInstance() bool`

HasViaNetworkInstance returns a boolean if a field has been set.

### GetNexthopTimestamp

`func (o *QueryNexthopNotificationItem) GetNexthopTimestamp() time.Time`

GetNexthopTimestamp returns the NexthopTimestamp field if non-nil, zero value otherwise.

### GetNexthopTimestampOk

`func (o *QueryNexthopNotificationItem) GetNexthopTimestampOk() (*time.Time, bool)`

GetNexthopTimestampOk returns a tuple with the NexthopTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthopTimestamp

`func (o *QueryNexthopNotificationItem) SetNexthopTimestamp(v time.Time)`

SetNexthopTimestamp sets NexthopTimestamp field to given value.

### HasNexthopTimestamp

`func (o *QueryNexthopNotificationItem) HasNexthopTimestamp() bool`

HasNexthopTimestamp returns a boolean if a field has been set.

### GetPriority

`func (o *QueryNexthopNotificationItem) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *QueryNexthopNotificationItem) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *QueryNexthopNotificationItem) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *QueryNexthopNotificationItem) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



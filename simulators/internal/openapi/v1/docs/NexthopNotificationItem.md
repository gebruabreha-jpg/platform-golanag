# NexthopNotificationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInstance** | Pointer to **string** |  | [optional] 
**AddressFamily** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**IpAddress**](IpAddress.md) |  | [optional] 
**ProtocolType** | Pointer to [**RouteType**](RouteType.md) |  | [optional] 
**Flags** | Pointer to **[]string** |  | [optional] 
**ViaAddress** | Pointer to [**IpAddress**](IpAddress.md) |  | [optional] 
**Metric** | Pointer to **int32** |  | [optional] 
**ViaNetworkInstance** | Pointer to **string** |  | [optional] 
**NexthopTimestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewNexthopNotificationItem

`func NewNexthopNotificationItem() *NexthopNotificationItem`

NewNexthopNotificationItem instantiates a new NexthopNotificationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexthopNotificationItemWithDefaults

`func NewNexthopNotificationItemWithDefaults() *NexthopNotificationItem`

NewNexthopNotificationItemWithDefaults instantiates a new NexthopNotificationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkInstance

`func (o *NexthopNotificationItem) GetNetworkInstance() string`

GetNetworkInstance returns the NetworkInstance field if non-nil, zero value otherwise.

### GetNetworkInstanceOk

`func (o *NexthopNotificationItem) GetNetworkInstanceOk() (*string, bool)`

GetNetworkInstanceOk returns a tuple with the NetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstance

`func (o *NexthopNotificationItem) SetNetworkInstance(v string)`

SetNetworkInstance sets NetworkInstance field to given value.

### HasNetworkInstance

`func (o *NexthopNotificationItem) HasNetworkInstance() bool`

HasNetworkInstance returns a boolean if a field has been set.

### GetAddressFamily

`func (o *NexthopNotificationItem) GetAddressFamily() string`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *NexthopNotificationItem) GetAddressFamilyOk() (*string, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *NexthopNotificationItem) SetAddressFamily(v string)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *NexthopNotificationItem) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetAddress

`func (o *NexthopNotificationItem) GetAddress() IpAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NexthopNotificationItem) GetAddressOk() (*IpAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NexthopNotificationItem) SetAddress(v IpAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NexthopNotificationItem) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetProtocolType

`func (o *NexthopNotificationItem) GetProtocolType() RouteType`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *NexthopNotificationItem) GetProtocolTypeOk() (*RouteType, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *NexthopNotificationItem) SetProtocolType(v RouteType)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *NexthopNotificationItem) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetFlags

`func (o *NexthopNotificationItem) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *NexthopNotificationItem) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *NexthopNotificationItem) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *NexthopNotificationItem) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetViaAddress

`func (o *NexthopNotificationItem) GetViaAddress() IpAddress`

GetViaAddress returns the ViaAddress field if non-nil, zero value otherwise.

### GetViaAddressOk

`func (o *NexthopNotificationItem) GetViaAddressOk() (*IpAddress, bool)`

GetViaAddressOk returns a tuple with the ViaAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaAddress

`func (o *NexthopNotificationItem) SetViaAddress(v IpAddress)`

SetViaAddress sets ViaAddress field to given value.

### HasViaAddress

`func (o *NexthopNotificationItem) HasViaAddress() bool`

HasViaAddress returns a boolean if a field has been set.

### GetMetric

`func (o *NexthopNotificationItem) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *NexthopNotificationItem) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *NexthopNotificationItem) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *NexthopNotificationItem) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetViaNetworkInstance

`func (o *NexthopNotificationItem) GetViaNetworkInstance() string`

GetViaNetworkInstance returns the ViaNetworkInstance field if non-nil, zero value otherwise.

### GetViaNetworkInstanceOk

`func (o *NexthopNotificationItem) GetViaNetworkInstanceOk() (*string, bool)`

GetViaNetworkInstanceOk returns a tuple with the ViaNetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaNetworkInstance

`func (o *NexthopNotificationItem) SetViaNetworkInstance(v string)`

SetViaNetworkInstance sets ViaNetworkInstance field to given value.

### HasViaNetworkInstance

`func (o *NexthopNotificationItem) HasViaNetworkInstance() bool`

HasViaNetworkInstance returns a boolean if a field has been set.

### GetNexthopTimestamp

`func (o *NexthopNotificationItem) GetNexthopTimestamp() time.Time`

GetNexthopTimestamp returns the NexthopTimestamp field if non-nil, zero value otherwise.

### GetNexthopTimestampOk

`func (o *NexthopNotificationItem) GetNexthopTimestampOk() (*time.Time, bool)`

GetNexthopTimestampOk returns a tuple with the NexthopTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthopTimestamp

`func (o *NexthopNotificationItem) SetNexthopTimestamp(v time.Time)`

SetNexthopTimestamp sets NexthopTimestamp field to given value.

### HasNexthopTimestamp

`func (o *NexthopNotificationItem) HasNexthopTimestamp() bool`

HasNexthopTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



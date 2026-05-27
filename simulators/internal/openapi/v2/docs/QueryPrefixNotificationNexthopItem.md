# QueryPrefixNotificationNexthopItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressInfo** | Pointer to [**NexthopAddressInfo**](NexthopAddressInfo.md) | The queried address info | [optional] 
**Metric** | Pointer to **int32** |  | [optional] 
**ResolvedNetworkInstanceName** | Pointer to **string** |  | [optional] 
**ProtocolType** | Pointer to [**RouteType**](RouteType.md) |  | [optional] 
**Priority** | Pointer to **int32** | Nexthop priority. The higher the value, the lower the priority | [optional] 

## Methods

### NewQueryPrefixNotificationNexthopItem

`func NewQueryPrefixNotificationNexthopItem() *QueryPrefixNotificationNexthopItem`

NewQueryPrefixNotificationNexthopItem instantiates a new QueryPrefixNotificationNexthopItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryPrefixNotificationNexthopItemWithDefaults

`func NewQueryPrefixNotificationNexthopItemWithDefaults() *QueryPrefixNotificationNexthopItem`

NewQueryPrefixNotificationNexthopItemWithDefaults instantiates a new QueryPrefixNotificationNexthopItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressInfo

`func (o *QueryPrefixNotificationNexthopItem) GetAddressInfo() NexthopAddressInfo`

GetAddressInfo returns the AddressInfo field if non-nil, zero value otherwise.

### GetAddressInfoOk

`func (o *QueryPrefixNotificationNexthopItem) GetAddressInfoOk() (*NexthopAddressInfo, bool)`

GetAddressInfoOk returns a tuple with the AddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressInfo

`func (o *QueryPrefixNotificationNexthopItem) SetAddressInfo(v NexthopAddressInfo)`

SetAddressInfo sets AddressInfo field to given value.

### HasAddressInfo

`func (o *QueryPrefixNotificationNexthopItem) HasAddressInfo() bool`

HasAddressInfo returns a boolean if a field has been set.

### GetMetric

`func (o *QueryPrefixNotificationNexthopItem) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *QueryPrefixNotificationNexthopItem) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *QueryPrefixNotificationNexthopItem) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *QueryPrefixNotificationNexthopItem) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetResolvedNetworkInstanceName

`func (o *QueryPrefixNotificationNexthopItem) GetResolvedNetworkInstanceName() string`

GetResolvedNetworkInstanceName returns the ResolvedNetworkInstanceName field if non-nil, zero value otherwise.

### GetResolvedNetworkInstanceNameOk

`func (o *QueryPrefixNotificationNexthopItem) GetResolvedNetworkInstanceNameOk() (*string, bool)`

GetResolvedNetworkInstanceNameOk returns a tuple with the ResolvedNetworkInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedNetworkInstanceName

`func (o *QueryPrefixNotificationNexthopItem) SetResolvedNetworkInstanceName(v string)`

SetResolvedNetworkInstanceName sets ResolvedNetworkInstanceName field to given value.

### HasResolvedNetworkInstanceName

`func (o *QueryPrefixNotificationNexthopItem) HasResolvedNetworkInstanceName() bool`

HasResolvedNetworkInstanceName returns a boolean if a field has been set.

### GetProtocolType

`func (o *QueryPrefixNotificationNexthopItem) GetProtocolType() RouteType`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *QueryPrefixNotificationNexthopItem) GetProtocolTypeOk() (*RouteType, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *QueryPrefixNotificationNexthopItem) SetProtocolType(v RouteType)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *QueryPrefixNotificationNexthopItem) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetPriority

`func (o *QueryPrefixNotificationNexthopItem) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *QueryPrefixNotificationNexthopItem) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *QueryPrefixNotificationNexthopItem) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *QueryPrefixNotificationNexthopItem) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



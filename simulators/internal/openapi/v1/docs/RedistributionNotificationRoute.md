# RedistributionNotificationRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | Pointer to **string** | Indicates route operation type | [optional] 
**Metric** | Pointer to **int32** |  | [optional] 
**AddressFamily** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to [**IpPrefix**](IpPrefix.md) |  | [optional] 
**NexthopsNumber** | Pointer to **int32** |  | [optional] 
**ProtocolType** | Pointer to [**RouteType**](RouteType.md) |  | [optional] 
**Tag** | Pointer to **int32** |  | [optional] 
**AsNum** | Pointer to **int32** |  | [optional] 
**LspLabel** | Pointer to **int32** |  | [optional] 
**Distance** | Pointer to **int32** |  | [optional] 
**Nexthops** | Pointer to [**[]RedistributionRouteNexthop**](RedistributionRouteNexthop.md) |  | [optional] 

## Methods

### NewRedistributionNotificationRoute

`func NewRedistributionNotificationRoute() *RedistributionNotificationRoute`

NewRedistributionNotificationRoute instantiates a new RedistributionNotificationRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedistributionNotificationRouteWithDefaults

`func NewRedistributionNotificationRouteWithDefaults() *RedistributionNotificationRoute`

NewRedistributionNotificationRouteWithDefaults instantiates a new RedistributionNotificationRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *RedistributionNotificationRoute) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *RedistributionNotificationRoute) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *RedistributionNotificationRoute) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *RedistributionNotificationRoute) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetMetric

`func (o *RedistributionNotificationRoute) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *RedistributionNotificationRoute) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *RedistributionNotificationRoute) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *RedistributionNotificationRoute) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetAddressFamily

`func (o *RedistributionNotificationRoute) GetAddressFamily() string`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *RedistributionNotificationRoute) GetAddressFamilyOk() (*string, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *RedistributionNotificationRoute) SetAddressFamily(v string)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *RedistributionNotificationRoute) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetPrefix

`func (o *RedistributionNotificationRoute) GetPrefix() IpPrefix`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *RedistributionNotificationRoute) GetPrefixOk() (*IpPrefix, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *RedistributionNotificationRoute) SetPrefix(v IpPrefix)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *RedistributionNotificationRoute) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetNexthopsNumber

`func (o *RedistributionNotificationRoute) GetNexthopsNumber() int32`

GetNexthopsNumber returns the NexthopsNumber field if non-nil, zero value otherwise.

### GetNexthopsNumberOk

`func (o *RedistributionNotificationRoute) GetNexthopsNumberOk() (*int32, bool)`

GetNexthopsNumberOk returns a tuple with the NexthopsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthopsNumber

`func (o *RedistributionNotificationRoute) SetNexthopsNumber(v int32)`

SetNexthopsNumber sets NexthopsNumber field to given value.

### HasNexthopsNumber

`func (o *RedistributionNotificationRoute) HasNexthopsNumber() bool`

HasNexthopsNumber returns a boolean if a field has been set.

### GetProtocolType

`func (o *RedistributionNotificationRoute) GetProtocolType() RouteType`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *RedistributionNotificationRoute) GetProtocolTypeOk() (*RouteType, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *RedistributionNotificationRoute) SetProtocolType(v RouteType)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *RedistributionNotificationRoute) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetTag

`func (o *RedistributionNotificationRoute) GetTag() int32`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *RedistributionNotificationRoute) GetTagOk() (*int32, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *RedistributionNotificationRoute) SetTag(v int32)`

SetTag sets Tag field to given value.

### HasTag

`func (o *RedistributionNotificationRoute) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetAsNum

`func (o *RedistributionNotificationRoute) GetAsNum() int32`

GetAsNum returns the AsNum field if non-nil, zero value otherwise.

### GetAsNumOk

`func (o *RedistributionNotificationRoute) GetAsNumOk() (*int32, bool)`

GetAsNumOk returns a tuple with the AsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNum

`func (o *RedistributionNotificationRoute) SetAsNum(v int32)`

SetAsNum sets AsNum field to given value.

### HasAsNum

`func (o *RedistributionNotificationRoute) HasAsNum() bool`

HasAsNum returns a boolean if a field has been set.

### GetLspLabel

`func (o *RedistributionNotificationRoute) GetLspLabel() int32`

GetLspLabel returns the LspLabel field if non-nil, zero value otherwise.

### GetLspLabelOk

`func (o *RedistributionNotificationRoute) GetLspLabelOk() (*int32, bool)`

GetLspLabelOk returns a tuple with the LspLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLspLabel

`func (o *RedistributionNotificationRoute) SetLspLabel(v int32)`

SetLspLabel sets LspLabel field to given value.

### HasLspLabel

`func (o *RedistributionNotificationRoute) HasLspLabel() bool`

HasLspLabel returns a boolean if a field has been set.

### GetDistance

`func (o *RedistributionNotificationRoute) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *RedistributionNotificationRoute) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *RedistributionNotificationRoute) SetDistance(v int32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *RedistributionNotificationRoute) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetNexthops

`func (o *RedistributionNotificationRoute) GetNexthops() []RedistributionRouteNexthop`

GetNexthops returns the Nexthops field if non-nil, zero value otherwise.

### GetNexthopsOk

`func (o *RedistributionNotificationRoute) GetNexthopsOk() (*[]RedistributionRouteNexthop, bool)`

GetNexthopsOk returns a tuple with the Nexthops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthops

`func (o *RedistributionNotificationRoute) SetNexthops(v []RedistributionRouteNexthop)`

SetNexthops sets Nexthops field to given value.

### HasNexthops

`func (o *RedistributionNotificationRoute) HasNexthops() bool`

HasNexthops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



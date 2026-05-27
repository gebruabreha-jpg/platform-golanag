# PrefixNotificationNexthop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamily** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**IpAddress**](IpAddress.md) |  | [optional] 
**Metric** | Pointer to **int32** |  | [optional] 
**ResolvedNetworkInstanceName** | Pointer to **string** |  | [optional] 
**ProtocolType** | Pointer to [**RouteType**](RouteType.md) |  | [optional] 

## Methods

### NewPrefixNotificationNexthop

`func NewPrefixNotificationNexthop() *PrefixNotificationNexthop`

NewPrefixNotificationNexthop instantiates a new PrefixNotificationNexthop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefixNotificationNexthopWithDefaults

`func NewPrefixNotificationNexthopWithDefaults() *PrefixNotificationNexthop`

NewPrefixNotificationNexthopWithDefaults instantiates a new PrefixNotificationNexthop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamily

`func (o *PrefixNotificationNexthop) GetAddressFamily() string`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *PrefixNotificationNexthop) GetAddressFamilyOk() (*string, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *PrefixNotificationNexthop) SetAddressFamily(v string)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *PrefixNotificationNexthop) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetAddress

`func (o *PrefixNotificationNexthop) GetAddress() IpAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PrefixNotificationNexthop) GetAddressOk() (*IpAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PrefixNotificationNexthop) SetAddress(v IpAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PrefixNotificationNexthop) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMetric

`func (o *PrefixNotificationNexthop) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *PrefixNotificationNexthop) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *PrefixNotificationNexthop) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *PrefixNotificationNexthop) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetResolvedNetworkInstanceName

`func (o *PrefixNotificationNexthop) GetResolvedNetworkInstanceName() string`

GetResolvedNetworkInstanceName returns the ResolvedNetworkInstanceName field if non-nil, zero value otherwise.

### GetResolvedNetworkInstanceNameOk

`func (o *PrefixNotificationNexthop) GetResolvedNetworkInstanceNameOk() (*string, bool)`

GetResolvedNetworkInstanceNameOk returns a tuple with the ResolvedNetworkInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedNetworkInstanceName

`func (o *PrefixNotificationNexthop) SetResolvedNetworkInstanceName(v string)`

SetResolvedNetworkInstanceName sets ResolvedNetworkInstanceName field to given value.

### HasResolvedNetworkInstanceName

`func (o *PrefixNotificationNexthop) HasResolvedNetworkInstanceName() bool`

HasResolvedNetworkInstanceName returns a boolean if a field has been set.

### GetProtocolType

`func (o *PrefixNotificationNexthop) GetProtocolType() RouteType`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *PrefixNotificationNexthop) GetProtocolTypeOk() (*RouteType, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *PrefixNotificationNexthop) SetProtocolType(v RouteType)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *PrefixNotificationNexthop) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



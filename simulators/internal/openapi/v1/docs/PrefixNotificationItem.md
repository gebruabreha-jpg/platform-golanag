# PrefixNotificationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsFound** | Pointer to **string** |  | [optional] 
**NetworkInstance** | Pointer to **string** |  | [optional] 
**AddressFamily** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to [**IpPrefix**](IpPrefix.md) |  | [optional] 
**RetNexthopType** | Pointer to **string** |  | [optional] 
**NumIpNexthop** | Pointer to **int32** | number of ip type&#39;s nexthop | [optional] 
**NumLspNexthop** | Pointer to **int32** | number of lsp type&#39;s nexthop | [optional] 
**InterNetworkInstanceFlag** | Pointer to **bool** | if across network-instance | [optional] 
**MatchedPrefix** | Pointer to [**IpPrefix**](IpPrefix.md) |  | [optional] 
**Nexthops** | Pointer to [**[]PrefixNotificationNexthop**](PrefixNotificationNexthop.md) | nexthops of matching prefixes found | [optional] 

## Methods

### NewPrefixNotificationItem

`func NewPrefixNotificationItem() *PrefixNotificationItem`

NewPrefixNotificationItem instantiates a new PrefixNotificationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefixNotificationItemWithDefaults

`func NewPrefixNotificationItemWithDefaults() *PrefixNotificationItem`

NewPrefixNotificationItemWithDefaults instantiates a new PrefixNotificationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsFound

`func (o *PrefixNotificationItem) GetIsFound() string`

GetIsFound returns the IsFound field if non-nil, zero value otherwise.

### GetIsFoundOk

`func (o *PrefixNotificationItem) GetIsFoundOk() (*string, bool)`

GetIsFoundOk returns a tuple with the IsFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFound

`func (o *PrefixNotificationItem) SetIsFound(v string)`

SetIsFound sets IsFound field to given value.

### HasIsFound

`func (o *PrefixNotificationItem) HasIsFound() bool`

HasIsFound returns a boolean if a field has been set.

### GetNetworkInstance

`func (o *PrefixNotificationItem) GetNetworkInstance() string`

GetNetworkInstance returns the NetworkInstance field if non-nil, zero value otherwise.

### GetNetworkInstanceOk

`func (o *PrefixNotificationItem) GetNetworkInstanceOk() (*string, bool)`

GetNetworkInstanceOk returns a tuple with the NetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstance

`func (o *PrefixNotificationItem) SetNetworkInstance(v string)`

SetNetworkInstance sets NetworkInstance field to given value.

### HasNetworkInstance

`func (o *PrefixNotificationItem) HasNetworkInstance() bool`

HasNetworkInstance returns a boolean if a field has been set.

### GetAddressFamily

`func (o *PrefixNotificationItem) GetAddressFamily() string`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *PrefixNotificationItem) GetAddressFamilyOk() (*string, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *PrefixNotificationItem) SetAddressFamily(v string)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *PrefixNotificationItem) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetPrefix

`func (o *PrefixNotificationItem) GetPrefix() IpPrefix`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PrefixNotificationItem) GetPrefixOk() (*IpPrefix, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PrefixNotificationItem) SetPrefix(v IpPrefix)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PrefixNotificationItem) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRetNexthopType

`func (o *PrefixNotificationItem) GetRetNexthopType() string`

GetRetNexthopType returns the RetNexthopType field if non-nil, zero value otherwise.

### GetRetNexthopTypeOk

`func (o *PrefixNotificationItem) GetRetNexthopTypeOk() (*string, bool)`

GetRetNexthopTypeOk returns a tuple with the RetNexthopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetNexthopType

`func (o *PrefixNotificationItem) SetRetNexthopType(v string)`

SetRetNexthopType sets RetNexthopType field to given value.

### HasRetNexthopType

`func (o *PrefixNotificationItem) HasRetNexthopType() bool`

HasRetNexthopType returns a boolean if a field has been set.

### GetNumIpNexthop

`func (o *PrefixNotificationItem) GetNumIpNexthop() int32`

GetNumIpNexthop returns the NumIpNexthop field if non-nil, zero value otherwise.

### GetNumIpNexthopOk

`func (o *PrefixNotificationItem) GetNumIpNexthopOk() (*int32, bool)`

GetNumIpNexthopOk returns a tuple with the NumIpNexthop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumIpNexthop

`func (o *PrefixNotificationItem) SetNumIpNexthop(v int32)`

SetNumIpNexthop sets NumIpNexthop field to given value.

### HasNumIpNexthop

`func (o *PrefixNotificationItem) HasNumIpNexthop() bool`

HasNumIpNexthop returns a boolean if a field has been set.

### GetNumLspNexthop

`func (o *PrefixNotificationItem) GetNumLspNexthop() int32`

GetNumLspNexthop returns the NumLspNexthop field if non-nil, zero value otherwise.

### GetNumLspNexthopOk

`func (o *PrefixNotificationItem) GetNumLspNexthopOk() (*int32, bool)`

GetNumLspNexthopOk returns a tuple with the NumLspNexthop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLspNexthop

`func (o *PrefixNotificationItem) SetNumLspNexthop(v int32)`

SetNumLspNexthop sets NumLspNexthop field to given value.

### HasNumLspNexthop

`func (o *PrefixNotificationItem) HasNumLspNexthop() bool`

HasNumLspNexthop returns a boolean if a field has been set.

### GetInterNetworkInstanceFlag

`func (o *PrefixNotificationItem) GetInterNetworkInstanceFlag() bool`

GetInterNetworkInstanceFlag returns the InterNetworkInstanceFlag field if non-nil, zero value otherwise.

### GetInterNetworkInstanceFlagOk

`func (o *PrefixNotificationItem) GetInterNetworkInstanceFlagOk() (*bool, bool)`

GetInterNetworkInstanceFlagOk returns a tuple with the InterNetworkInstanceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterNetworkInstanceFlag

`func (o *PrefixNotificationItem) SetInterNetworkInstanceFlag(v bool)`

SetInterNetworkInstanceFlag sets InterNetworkInstanceFlag field to given value.

### HasInterNetworkInstanceFlag

`func (o *PrefixNotificationItem) HasInterNetworkInstanceFlag() bool`

HasInterNetworkInstanceFlag returns a boolean if a field has been set.

### GetMatchedPrefix

`func (o *PrefixNotificationItem) GetMatchedPrefix() IpPrefix`

GetMatchedPrefix returns the MatchedPrefix field if non-nil, zero value otherwise.

### GetMatchedPrefixOk

`func (o *PrefixNotificationItem) GetMatchedPrefixOk() (*IpPrefix, bool)`

GetMatchedPrefixOk returns a tuple with the MatchedPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedPrefix

`func (o *PrefixNotificationItem) SetMatchedPrefix(v IpPrefix)`

SetMatchedPrefix sets MatchedPrefix field to given value.

### HasMatchedPrefix

`func (o *PrefixNotificationItem) HasMatchedPrefix() bool`

HasMatchedPrefix returns a boolean if a field has been set.

### GetNexthops

`func (o *PrefixNotificationItem) GetNexthops() []PrefixNotificationNexthop`

GetNexthops returns the Nexthops field if non-nil, zero value otherwise.

### GetNexthopsOk

`func (o *PrefixNotificationItem) GetNexthopsOk() (*[]PrefixNotificationNexthop, bool)`

GetNexthopsOk returns a tuple with the Nexthops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthops

`func (o *PrefixNotificationItem) SetNexthops(v []PrefixNotificationNexthop)`

SetNexthops sets Nexthops field to given value.

### HasNexthops

`func (o *PrefixNotificationItem) HasNexthops() bool`

HasNexthops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



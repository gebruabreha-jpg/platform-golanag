# QueryPrefixNotificationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsFound** | Pointer to [**QueryPrefixNotificationItemIsFound**](QueryPrefixNotificationItemIsFound.md) |  | [optional] 
**NetworkInstance** | Pointer to **string** |  | [optional] 
**AddressFamily** | Pointer to [**AddressFamily**](AddressFamily.md) |  | [optional] 
**Prefix** | Pointer to **string** | A valid IP prefix. | [optional] 
**RetNexthopType** | Pointer to [**PrefixNexthopType**](PrefixNexthopType.md) |  | [optional] 
**NumServiceNexthop** | Pointer to **int32** |  | [optional] 
**NumIpNexthop** | Pointer to **int32** | number of ip type\&quot;s nexthop | [optional] 
**NumLspNexthop** | Pointer to **int32** | number of lsp type\&quot;s nexthop | [optional] 
**InterNetworkInstanceFlag** | Pointer to **bool** | if across networkInstance | [optional] 
**MatchedPrefix** | Pointer to **string** | A valid IP prefix. | [optional] 
**Nexthops** | Pointer to [**[]QueryPrefixNotificationNexthopItem**](QueryPrefixNotificationNexthopItem.md) | nexthops of matching prefixes found | [optional] 

## Methods

### NewQueryPrefixNotificationItem

`func NewQueryPrefixNotificationItem() *QueryPrefixNotificationItem`

NewQueryPrefixNotificationItem instantiates a new QueryPrefixNotificationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryPrefixNotificationItemWithDefaults

`func NewQueryPrefixNotificationItemWithDefaults() *QueryPrefixNotificationItem`

NewQueryPrefixNotificationItemWithDefaults instantiates a new QueryPrefixNotificationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsFound

`func (o *QueryPrefixNotificationItem) GetIsFound() QueryPrefixNotificationItemIsFound`

GetIsFound returns the IsFound field if non-nil, zero value otherwise.

### GetIsFoundOk

`func (o *QueryPrefixNotificationItem) GetIsFoundOk() (*QueryPrefixNotificationItemIsFound, bool)`

GetIsFoundOk returns a tuple with the IsFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFound

`func (o *QueryPrefixNotificationItem) SetIsFound(v QueryPrefixNotificationItemIsFound)`

SetIsFound sets IsFound field to given value.

### HasIsFound

`func (o *QueryPrefixNotificationItem) HasIsFound() bool`

HasIsFound returns a boolean if a field has been set.

### GetNetworkInstance

`func (o *QueryPrefixNotificationItem) GetNetworkInstance() string`

GetNetworkInstance returns the NetworkInstance field if non-nil, zero value otherwise.

### GetNetworkInstanceOk

`func (o *QueryPrefixNotificationItem) GetNetworkInstanceOk() (*string, bool)`

GetNetworkInstanceOk returns a tuple with the NetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstance

`func (o *QueryPrefixNotificationItem) SetNetworkInstance(v string)`

SetNetworkInstance sets NetworkInstance field to given value.

### HasNetworkInstance

`func (o *QueryPrefixNotificationItem) HasNetworkInstance() bool`

HasNetworkInstance returns a boolean if a field has been set.

### GetAddressFamily

`func (o *QueryPrefixNotificationItem) GetAddressFamily() AddressFamily`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *QueryPrefixNotificationItem) GetAddressFamilyOk() (*AddressFamily, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *QueryPrefixNotificationItem) SetAddressFamily(v AddressFamily)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *QueryPrefixNotificationItem) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetPrefix

`func (o *QueryPrefixNotificationItem) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *QueryPrefixNotificationItem) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *QueryPrefixNotificationItem) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *QueryPrefixNotificationItem) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRetNexthopType

`func (o *QueryPrefixNotificationItem) GetRetNexthopType() PrefixNexthopType`

GetRetNexthopType returns the RetNexthopType field if non-nil, zero value otherwise.

### GetRetNexthopTypeOk

`func (o *QueryPrefixNotificationItem) GetRetNexthopTypeOk() (*PrefixNexthopType, bool)`

GetRetNexthopTypeOk returns a tuple with the RetNexthopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetNexthopType

`func (o *QueryPrefixNotificationItem) SetRetNexthopType(v PrefixNexthopType)`

SetRetNexthopType sets RetNexthopType field to given value.

### HasRetNexthopType

`func (o *QueryPrefixNotificationItem) HasRetNexthopType() bool`

HasRetNexthopType returns a boolean if a field has been set.

### GetNumServiceNexthop

`func (o *QueryPrefixNotificationItem) GetNumServiceNexthop() int32`

GetNumServiceNexthop returns the NumServiceNexthop field if non-nil, zero value otherwise.

### GetNumServiceNexthopOk

`func (o *QueryPrefixNotificationItem) GetNumServiceNexthopOk() (*int32, bool)`

GetNumServiceNexthopOk returns a tuple with the NumServiceNexthop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumServiceNexthop

`func (o *QueryPrefixNotificationItem) SetNumServiceNexthop(v int32)`

SetNumServiceNexthop sets NumServiceNexthop field to given value.

### HasNumServiceNexthop

`func (o *QueryPrefixNotificationItem) HasNumServiceNexthop() bool`

HasNumServiceNexthop returns a boolean if a field has been set.

### GetNumIpNexthop

`func (o *QueryPrefixNotificationItem) GetNumIpNexthop() int32`

GetNumIpNexthop returns the NumIpNexthop field if non-nil, zero value otherwise.

### GetNumIpNexthopOk

`func (o *QueryPrefixNotificationItem) GetNumIpNexthopOk() (*int32, bool)`

GetNumIpNexthopOk returns a tuple with the NumIpNexthop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumIpNexthop

`func (o *QueryPrefixNotificationItem) SetNumIpNexthop(v int32)`

SetNumIpNexthop sets NumIpNexthop field to given value.

### HasNumIpNexthop

`func (o *QueryPrefixNotificationItem) HasNumIpNexthop() bool`

HasNumIpNexthop returns a boolean if a field has been set.

### GetNumLspNexthop

`func (o *QueryPrefixNotificationItem) GetNumLspNexthop() int32`

GetNumLspNexthop returns the NumLspNexthop field if non-nil, zero value otherwise.

### GetNumLspNexthopOk

`func (o *QueryPrefixNotificationItem) GetNumLspNexthopOk() (*int32, bool)`

GetNumLspNexthopOk returns a tuple with the NumLspNexthop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLspNexthop

`func (o *QueryPrefixNotificationItem) SetNumLspNexthop(v int32)`

SetNumLspNexthop sets NumLspNexthop field to given value.

### HasNumLspNexthop

`func (o *QueryPrefixNotificationItem) HasNumLspNexthop() bool`

HasNumLspNexthop returns a boolean if a field has been set.

### GetInterNetworkInstanceFlag

`func (o *QueryPrefixNotificationItem) GetInterNetworkInstanceFlag() bool`

GetInterNetworkInstanceFlag returns the InterNetworkInstanceFlag field if non-nil, zero value otherwise.

### GetInterNetworkInstanceFlagOk

`func (o *QueryPrefixNotificationItem) GetInterNetworkInstanceFlagOk() (*bool, bool)`

GetInterNetworkInstanceFlagOk returns a tuple with the InterNetworkInstanceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterNetworkInstanceFlag

`func (o *QueryPrefixNotificationItem) SetInterNetworkInstanceFlag(v bool)`

SetInterNetworkInstanceFlag sets InterNetworkInstanceFlag field to given value.

### HasInterNetworkInstanceFlag

`func (o *QueryPrefixNotificationItem) HasInterNetworkInstanceFlag() bool`

HasInterNetworkInstanceFlag returns a boolean if a field has been set.

### GetMatchedPrefix

`func (o *QueryPrefixNotificationItem) GetMatchedPrefix() string`

GetMatchedPrefix returns the MatchedPrefix field if non-nil, zero value otherwise.

### GetMatchedPrefixOk

`func (o *QueryPrefixNotificationItem) GetMatchedPrefixOk() (*string, bool)`

GetMatchedPrefixOk returns a tuple with the MatchedPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedPrefix

`func (o *QueryPrefixNotificationItem) SetMatchedPrefix(v string)`

SetMatchedPrefix sets MatchedPrefix field to given value.

### HasMatchedPrefix

`func (o *QueryPrefixNotificationItem) HasMatchedPrefix() bool`

HasMatchedPrefix returns a boolean if a field has been set.

### GetNexthops

`func (o *QueryPrefixNotificationItem) GetNexthops() []QueryPrefixNotificationNexthopItem`

GetNexthops returns the Nexthops field if non-nil, zero value otherwise.

### GetNexthopsOk

`func (o *QueryPrefixNotificationItem) GetNexthopsOk() (*[]QueryPrefixNotificationNexthopItem, bool)`

GetNexthopsOk returns a tuple with the Nexthops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthops

`func (o *QueryPrefixNotificationItem) SetNexthops(v []QueryPrefixNotificationNexthopItem)`

SetNexthops sets Nexthops field to given value.

### HasNexthops

`func (o *QueryPrefixNotificationItem) HasNexthops() bool`

HasNexthops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



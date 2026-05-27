# PrefixRegisterItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to [**IpPrefix**](IpPrefix.md) |  | [optional] 
**AddressFamily** | Pointer to **string** |  | [optional] 
**LookupType** | Pointer to **string** |  | [optional] 
**NexthopType** | Pointer to **string** |  | [optional] 

## Methods

### NewPrefixRegisterItem

`func NewPrefixRegisterItem() *PrefixRegisterItem`

NewPrefixRegisterItem instantiates a new PrefixRegisterItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefixRegisterItemWithDefaults

`func NewPrefixRegisterItemWithDefaults() *PrefixRegisterItem`

NewPrefixRegisterItemWithDefaults instantiates a new PrefixRegisterItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *PrefixRegisterItem) GetPrefix() IpPrefix`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PrefixRegisterItem) GetPrefixOk() (*IpPrefix, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PrefixRegisterItem) SetPrefix(v IpPrefix)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PrefixRegisterItem) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetAddressFamily

`func (o *PrefixRegisterItem) GetAddressFamily() string`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *PrefixRegisterItem) GetAddressFamilyOk() (*string, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *PrefixRegisterItem) SetAddressFamily(v string)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *PrefixRegisterItem) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetLookupType

`func (o *PrefixRegisterItem) GetLookupType() string`

GetLookupType returns the LookupType field if non-nil, zero value otherwise.

### GetLookupTypeOk

`func (o *PrefixRegisterItem) GetLookupTypeOk() (*string, bool)`

GetLookupTypeOk returns a tuple with the LookupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupType

`func (o *PrefixRegisterItem) SetLookupType(v string)`

SetLookupType sets LookupType field to given value.

### HasLookupType

`func (o *PrefixRegisterItem) HasLookupType() bool`

HasLookupType returns a boolean if a field has been set.

### GetNexthopType

`func (o *PrefixRegisterItem) GetNexthopType() string`

GetNexthopType returns the NexthopType field if non-nil, zero value otherwise.

### GetNexthopTypeOk

`func (o *PrefixRegisterItem) GetNexthopTypeOk() (*string, bool)`

GetNexthopTypeOk returns a tuple with the NexthopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthopType

`func (o *PrefixRegisterItem) SetNexthopType(v string)`

SetNexthopType sets NexthopType field to given value.

### HasNexthopType

`func (o *PrefixRegisterItem) HasNexthopType() bool`

HasNexthopType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



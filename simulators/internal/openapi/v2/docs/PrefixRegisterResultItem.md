# PrefixRegisterResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInstance** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** | A valid IP prefix. | [optional] 
**AddressFamily** | Pointer to [**AddressFamily**](AddressFamily.md) |  | [optional] 
**Status** | Pointer to **int32** | http status code, such as 2xx/3xx/4xx/5xx etc.. | [optional] 

## Methods

### NewPrefixRegisterResultItem

`func NewPrefixRegisterResultItem() *PrefixRegisterResultItem`

NewPrefixRegisterResultItem instantiates a new PrefixRegisterResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefixRegisterResultItemWithDefaults

`func NewPrefixRegisterResultItemWithDefaults() *PrefixRegisterResultItem`

NewPrefixRegisterResultItemWithDefaults instantiates a new PrefixRegisterResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkInstance

`func (o *PrefixRegisterResultItem) GetNetworkInstance() string`

GetNetworkInstance returns the NetworkInstance field if non-nil, zero value otherwise.

### GetNetworkInstanceOk

`func (o *PrefixRegisterResultItem) GetNetworkInstanceOk() (*string, bool)`

GetNetworkInstanceOk returns a tuple with the NetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstance

`func (o *PrefixRegisterResultItem) SetNetworkInstance(v string)`

SetNetworkInstance sets NetworkInstance field to given value.

### HasNetworkInstance

`func (o *PrefixRegisterResultItem) HasNetworkInstance() bool`

HasNetworkInstance returns a boolean if a field has been set.

### GetPrefix

`func (o *PrefixRegisterResultItem) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PrefixRegisterResultItem) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PrefixRegisterResultItem) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PrefixRegisterResultItem) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetAddressFamily

`func (o *PrefixRegisterResultItem) GetAddressFamily() AddressFamily`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *PrefixRegisterResultItem) GetAddressFamilyOk() (*AddressFamily, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *PrefixRegisterResultItem) SetAddressFamily(v AddressFamily)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *PrefixRegisterResultItem) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetStatus

`func (o *PrefixRegisterResultItem) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrefixRegisterResultItem) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrefixRegisterResultItem) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrefixRegisterResultItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



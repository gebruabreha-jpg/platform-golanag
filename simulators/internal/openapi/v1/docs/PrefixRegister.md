# PrefixRegister

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to [**ClientIdentifier**](ClientIdentifier.md) |  | [optional] 
**Prefixes** | Pointer to [**[]PrefixRegisterItem**](PrefixRegisterItem.md) |  | [optional] 

## Methods

### NewPrefixRegister

`func NewPrefixRegister() *PrefixRegister`

NewPrefixRegister instantiates a new PrefixRegister object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefixRegisterWithDefaults

`func NewPrefixRegisterWithDefaults() *PrefixRegister`

NewPrefixRegisterWithDefaults instantiates a new PrefixRegister object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *PrefixRegister) GetClientId() ClientIdentifier`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PrefixRegister) GetClientIdOk() (*ClientIdentifier, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PrefixRegister) SetClientId(v ClientIdentifier)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PrefixRegister) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetPrefixes

`func (o *PrefixRegister) GetPrefixes() []PrefixRegisterItem`

GetPrefixes returns the Prefixes field if non-nil, zero value otherwise.

### GetPrefixesOk

`func (o *PrefixRegister) GetPrefixesOk() (*[]PrefixRegisterItem, bool)`

GetPrefixesOk returns a tuple with the Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixes

`func (o *PrefixRegister) SetPrefixes(v []PrefixRegisterItem)`

SetPrefixes sets Prefixes field to given value.

### HasPrefixes

`func (o *PrefixRegister) HasPrefixes() bool`

HasPrefixes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



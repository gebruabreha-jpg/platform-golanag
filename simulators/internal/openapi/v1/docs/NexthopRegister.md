# NexthopRegister

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to [**ClientIdentifier**](ClientIdentifier.md) |  | [optional] 
**Nexthops** | Pointer to [**[]NexthopRegisterItem**](NexthopRegisterItem.md) |  | [optional] 

## Methods

### NewNexthopRegister

`func NewNexthopRegister() *NexthopRegister`

NewNexthopRegister instantiates a new NexthopRegister object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexthopRegisterWithDefaults

`func NewNexthopRegisterWithDefaults() *NexthopRegister`

NewNexthopRegisterWithDefaults instantiates a new NexthopRegister object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *NexthopRegister) GetClientId() ClientIdentifier`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *NexthopRegister) GetClientIdOk() (*ClientIdentifier, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *NexthopRegister) SetClientId(v ClientIdentifier)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *NexthopRegister) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetNexthops

`func (o *NexthopRegister) GetNexthops() []NexthopRegisterItem`

GetNexthops returns the Nexthops field if non-nil, zero value otherwise.

### GetNexthopsOk

`func (o *NexthopRegister) GetNexthopsOk() (*[]NexthopRegisterItem, bool)`

GetNexthopsOk returns a tuple with the Nexthops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthops

`func (o *NexthopRegister) SetNexthops(v []NexthopRegisterItem)`

SetNexthops sets Nexthops field to given value.

### HasNexthops

`func (o *NexthopRegister) HasNexthops() bool`

HasNexthops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



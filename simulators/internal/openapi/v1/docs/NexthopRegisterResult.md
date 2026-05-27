# NexthopRegisterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to [**ClientIdentifier**](ClientIdentifier.md) |  | [optional] 
**RegisterResult** | Pointer to [**[]NexthopRegisterResultItem**](NexthopRegisterResultItem.md) |  | [optional] 

## Methods

### NewNexthopRegisterResult

`func NewNexthopRegisterResult() *NexthopRegisterResult`

NewNexthopRegisterResult instantiates a new NexthopRegisterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexthopRegisterResultWithDefaults

`func NewNexthopRegisterResultWithDefaults() *NexthopRegisterResult`

NewNexthopRegisterResultWithDefaults instantiates a new NexthopRegisterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *NexthopRegisterResult) GetClientId() ClientIdentifier`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *NexthopRegisterResult) GetClientIdOk() (*ClientIdentifier, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *NexthopRegisterResult) SetClientId(v ClientIdentifier)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *NexthopRegisterResult) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetRegisterResult

`func (o *NexthopRegisterResult) GetRegisterResult() []NexthopRegisterResultItem`

GetRegisterResult returns the RegisterResult field if non-nil, zero value otherwise.

### GetRegisterResultOk

`func (o *NexthopRegisterResult) GetRegisterResultOk() (*[]NexthopRegisterResultItem, bool)`

GetRegisterResultOk returns a tuple with the RegisterResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterResult

`func (o *NexthopRegisterResult) SetRegisterResult(v []NexthopRegisterResultItem)`

SetRegisterResult sets RegisterResult field to given value.

### HasRegisterResult

`func (o *NexthopRegisterResult) HasRegisterResult() bool`

HasRegisterResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



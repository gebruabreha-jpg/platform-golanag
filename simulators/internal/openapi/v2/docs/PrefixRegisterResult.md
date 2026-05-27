# PrefixRegisterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Indicates client name, for either Producer or Consumer. For route-engine, client name is like {cre-pod-name} + {static/bgp/ifm}, for data-plane, client name is like {dp-pod-name}, for client that is service, client name is like {service-name}. | [optional] 
**RegisterResult** | Pointer to [**[]PrefixRegisterResultItem**](PrefixRegisterResultItem.md) |  | [optional] 

## Methods

### NewPrefixRegisterResult

`func NewPrefixRegisterResult() *PrefixRegisterResult`

NewPrefixRegisterResult instantiates a new PrefixRegisterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefixRegisterResultWithDefaults

`func NewPrefixRegisterResultWithDefaults() *PrefixRegisterResult`

NewPrefixRegisterResultWithDefaults instantiates a new PrefixRegisterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *PrefixRegisterResult) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PrefixRegisterResult) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PrefixRegisterResult) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PrefixRegisterResult) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetRegisterResult

`func (o *PrefixRegisterResult) GetRegisterResult() []PrefixRegisterResultItem`

GetRegisterResult returns the RegisterResult field if non-nil, zero value otherwise.

### GetRegisterResultOk

`func (o *PrefixRegisterResult) GetRegisterResultOk() (*[]PrefixRegisterResultItem, bool)`

GetRegisterResultOk returns a tuple with the RegisterResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterResult

`func (o *PrefixRegisterResult) SetRegisterResult(v []PrefixRegisterResultItem)`

SetRegisterResult sets RegisterResult field to given value.

### HasRegisterResult

`func (o *PrefixRegisterResult) HasRegisterResult() bool`

HasRegisterResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



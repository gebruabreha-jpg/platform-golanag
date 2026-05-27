# ClientIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Indicates client name. For route-engine, client name is like {cre-pod-name} + {static/bgp/ifm}, for data-plane, client name is like {dp-pod-name}, for client that is service, client name is like {service-name}. | [optional] 

## Methods

### NewClientIdentifier

`func NewClientIdentifier() *ClientIdentifier`

NewClientIdentifier instantiates a new ClientIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientIdentifierWithDefaults

`func NewClientIdentifierWithDefaults() *ClientIdentifier`

NewClientIdentifierWithDefaults instantiates a new ClientIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClientIdentifier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientIdentifier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientIdentifier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientIdentifier) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ClientUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to [**ClientIdentifier**](ClientIdentifier.md) |  | [optional] 
**ConsumerInfo** | Pointer to [**ConsumerUpdateInfo**](ConsumerUpdateInfo.md) |  | [optional] 

## Methods

### NewClientUpdate

`func NewClientUpdate() *ClientUpdate`

NewClientUpdate instantiates a new ClientUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientUpdateWithDefaults

`func NewClientUpdateWithDefaults() *ClientUpdate`

NewClientUpdateWithDefaults instantiates a new ClientUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ClientUpdate) GetClientId() ClientIdentifier`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClientUpdate) GetClientIdOk() (*ClientIdentifier, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClientUpdate) SetClientId(v ClientIdentifier)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ClientUpdate) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetConsumerInfo

`func (o *ClientUpdate) GetConsumerInfo() ConsumerUpdateInfo`

GetConsumerInfo returns the ConsumerInfo field if non-nil, zero value otherwise.

### GetConsumerInfoOk

`func (o *ClientUpdate) GetConsumerInfoOk() (*ConsumerUpdateInfo, bool)`

GetConsumerInfoOk returns a tuple with the ConsumerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerInfo

`func (o *ClientUpdate) SetConsumerInfo(v ConsumerUpdateInfo)`

SetConsumerInfo sets ConsumerInfo field to given value.

### HasConsumerInfo

`func (o *ClientUpdate) HasConsumerInfo() bool`

HasConsumerInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



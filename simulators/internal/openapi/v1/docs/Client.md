# Client

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to [**ClientIdentifier**](ClientIdentifier.md) |  | [optional] 
**ProducerInfo** | Pointer to [**ProducerInfo**](ProducerInfo.md) |  | [optional] 
**ConsumerInfo** | Pointer to [**ConsumerInfo**](ConsumerInfo.md) |  | [optional] 

## Methods

### NewClient

`func NewClient() *Client`

NewClient instantiates a new Client object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientWithDefaults

`func NewClientWithDefaults() *Client`

NewClientWithDefaults instantiates a new Client object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *Client) GetClientId() ClientIdentifier`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Client) GetClientIdOk() (*ClientIdentifier, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Client) SetClientId(v ClientIdentifier)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Client) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetProducerInfo

`func (o *Client) GetProducerInfo() ProducerInfo`

GetProducerInfo returns the ProducerInfo field if non-nil, zero value otherwise.

### GetProducerInfoOk

`func (o *Client) GetProducerInfoOk() (*ProducerInfo, bool)`

GetProducerInfoOk returns a tuple with the ProducerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerInfo

`func (o *Client) SetProducerInfo(v ProducerInfo)`

SetProducerInfo sets ProducerInfo field to given value.

### HasProducerInfo

`func (o *Client) HasProducerInfo() bool`

HasProducerInfo returns a boolean if a field has been set.

### GetConsumerInfo

`func (o *Client) GetConsumerInfo() ConsumerInfo`

GetConsumerInfo returns the ConsumerInfo field if non-nil, zero value otherwise.

### GetConsumerInfoOk

`func (o *Client) GetConsumerInfoOk() (*ConsumerInfo, bool)`

GetConsumerInfoOk returns a tuple with the ConsumerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerInfo

`func (o *Client) SetConsumerInfo(v ConsumerInfo)`

SetConsumerInfo sets ConsumerInfo field to given value.

### HasConsumerInfo

`func (o *Client) HasConsumerInfo() bool`

HasConsumerInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



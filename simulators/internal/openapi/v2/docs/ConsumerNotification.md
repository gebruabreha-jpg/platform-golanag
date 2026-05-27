# ConsumerNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerId** | Pointer to **string** | Indicates client name, for either Producer or Consumer. For route-engine, client name is like {cre-pod-name} + {static/bgp/ifm}, for data-plane, client name is like {dp-pod-name}, for client that is service, client name is like {service-name}. | [optional] 
**Msg** | Pointer to [**ConsumerNotificationMsg**](ConsumerNotificationMsg.md) |  | [optional] 

## Methods

### NewConsumerNotification

`func NewConsumerNotification() *ConsumerNotification`

NewConsumerNotification instantiates a new ConsumerNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerNotificationWithDefaults

`func NewConsumerNotificationWithDefaults() *ConsumerNotification`

NewConsumerNotificationWithDefaults instantiates a new ConsumerNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerId

`func (o *ConsumerNotification) GetConsumerId() string`

GetConsumerId returns the ConsumerId field if non-nil, zero value otherwise.

### GetConsumerIdOk

`func (o *ConsumerNotification) GetConsumerIdOk() (*string, bool)`

GetConsumerIdOk returns a tuple with the ConsumerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerId

`func (o *ConsumerNotification) SetConsumerId(v string)`

SetConsumerId sets ConsumerId field to given value.

### HasConsumerId

`func (o *ConsumerNotification) HasConsumerId() bool`

HasConsumerId returns a boolean if a field has been set.

### GetMsg

`func (o *ConsumerNotification) GetMsg() ConsumerNotificationMsg`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ConsumerNotification) GetMsgOk() (*ConsumerNotificationMsg, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ConsumerNotification) SetMsg(v ConsumerNotificationMsg)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ConsumerNotification) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



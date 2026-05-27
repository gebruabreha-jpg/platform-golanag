# ProducerNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProducerId** | Pointer to **string** | Indicates client name, for either Producer or Consumer. For route-engine, client name is like {cre-pod-name} + {static/bgp/ifm}, for data-plane, client name is like {dp-pod-name}, for client that is service, client name is like {service-name}. | [optional] 
**Msg** | Pointer to [**ProducerNotificationMsg**](ProducerNotificationMsg.md) |  | [optional] 

## Methods

### NewProducerNotification

`func NewProducerNotification() *ProducerNotification`

NewProducerNotification instantiates a new ProducerNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProducerNotificationWithDefaults

`func NewProducerNotificationWithDefaults() *ProducerNotification`

NewProducerNotificationWithDefaults instantiates a new ProducerNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducerId

`func (o *ProducerNotification) GetProducerId() string`

GetProducerId returns the ProducerId field if non-nil, zero value otherwise.

### GetProducerIdOk

`func (o *ProducerNotification) GetProducerIdOk() (*string, bool)`

GetProducerIdOk returns a tuple with the ProducerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerId

`func (o *ProducerNotification) SetProducerId(v string)`

SetProducerId sets ProducerId field to given value.

### HasProducerId

`func (o *ProducerNotification) HasProducerId() bool`

HasProducerId returns a boolean if a field has been set.

### GetMsg

`func (o *ProducerNotification) GetMsg() ProducerNotificationMsg`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ProducerNotification) GetMsgOk() (*ProducerNotificationMsg, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ProducerNotification) SetMsg(v ProducerNotificationMsg)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ProducerNotification) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# NotifyInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to [**ClientIdentifier**](ClientIdentifier.md) |  | [optional] 
**Msg** | Pointer to [**NotifyInformationMsg**](NotifyInformationMsg.md) |  | [optional] 

## Methods

### NewNotifyInformation

`func NewNotifyInformation() *NotifyInformation`

NewNotifyInformation instantiates a new NotifyInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyInformationWithDefaults

`func NewNotifyInformationWithDefaults() *NotifyInformation`

NewNotifyInformationWithDefaults instantiates a new NotifyInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *NotifyInformation) GetClientId() ClientIdentifier`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *NotifyInformation) GetClientIdOk() (*ClientIdentifier, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *NotifyInformation) SetClientId(v ClientIdentifier)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *NotifyInformation) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetMsg

`func (o *NotifyInformation) GetMsg() NotifyInformationMsg`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *NotifyInformation) GetMsgOk() (*NotifyInformationMsg, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *NotifyInformation) SetMsg(v NotifyInformationMsg)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *NotifyInformation) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



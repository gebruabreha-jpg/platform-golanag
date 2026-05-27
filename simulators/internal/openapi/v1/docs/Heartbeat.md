# Heartbeat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to [**ClientIdentifier**](ClientIdentifier.md) |  | [optional] 
**ClientType** | Pointer to **string** | Indicates the role of client | [optional] 

## Methods

### NewHeartbeat

`func NewHeartbeat() *Heartbeat`

NewHeartbeat instantiates a new Heartbeat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeartbeatWithDefaults

`func NewHeartbeatWithDefaults() *Heartbeat`

NewHeartbeatWithDefaults instantiates a new Heartbeat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *Heartbeat) GetClientId() ClientIdentifier`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Heartbeat) GetClientIdOk() (*ClientIdentifier, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Heartbeat) SetClientId(v ClientIdentifier)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Heartbeat) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientType

`func (o *Heartbeat) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *Heartbeat) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *Heartbeat) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *Heartbeat) HasClientType() bool`

HasClientType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



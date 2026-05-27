# RedistributionNotificationHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInstance** | Pointer to **string** |  | [optional] 
**AddressFamily** | Pointer to **string** |  | [optional] 
**ControlCommand** | Pointer to **string** | Indicates how BGP shall do. | [optional] 
**RedistributionEnd** | Pointer to **bool** | Indicates finish redistributing route in this round. | [optional] 

## Methods

### NewRedistributionNotificationHeader

`func NewRedistributionNotificationHeader() *RedistributionNotificationHeader`

NewRedistributionNotificationHeader instantiates a new RedistributionNotificationHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedistributionNotificationHeaderWithDefaults

`func NewRedistributionNotificationHeaderWithDefaults() *RedistributionNotificationHeader`

NewRedistributionNotificationHeaderWithDefaults instantiates a new RedistributionNotificationHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkInstance

`func (o *RedistributionNotificationHeader) GetNetworkInstance() string`

GetNetworkInstance returns the NetworkInstance field if non-nil, zero value otherwise.

### GetNetworkInstanceOk

`func (o *RedistributionNotificationHeader) GetNetworkInstanceOk() (*string, bool)`

GetNetworkInstanceOk returns a tuple with the NetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstance

`func (o *RedistributionNotificationHeader) SetNetworkInstance(v string)`

SetNetworkInstance sets NetworkInstance field to given value.

### HasNetworkInstance

`func (o *RedistributionNotificationHeader) HasNetworkInstance() bool`

HasNetworkInstance returns a boolean if a field has been set.

### GetAddressFamily

`func (o *RedistributionNotificationHeader) GetAddressFamily() string`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *RedistributionNotificationHeader) GetAddressFamilyOk() (*string, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *RedistributionNotificationHeader) SetAddressFamily(v string)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *RedistributionNotificationHeader) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetControlCommand

`func (o *RedistributionNotificationHeader) GetControlCommand() string`

GetControlCommand returns the ControlCommand field if non-nil, zero value otherwise.

### GetControlCommandOk

`func (o *RedistributionNotificationHeader) GetControlCommandOk() (*string, bool)`

GetControlCommandOk returns a tuple with the ControlCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCommand

`func (o *RedistributionNotificationHeader) SetControlCommand(v string)`

SetControlCommand sets ControlCommand field to given value.

### HasControlCommand

`func (o *RedistributionNotificationHeader) HasControlCommand() bool`

HasControlCommand returns a boolean if a field has been set.

### GetRedistributionEnd

`func (o *RedistributionNotificationHeader) GetRedistributionEnd() bool`

GetRedistributionEnd returns the RedistributionEnd field if non-nil, zero value otherwise.

### GetRedistributionEndOk

`func (o *RedistributionNotificationHeader) GetRedistributionEndOk() (*bool, bool)`

GetRedistributionEndOk returns a tuple with the RedistributionEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedistributionEnd

`func (o *RedistributionNotificationHeader) SetRedistributionEnd(v bool)`

SetRedistributionEnd sets RedistributionEnd field to given value.

### HasRedistributionEnd

`func (o *RedistributionNotificationHeader) HasRedistributionEnd() bool`

HasRedistributionEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



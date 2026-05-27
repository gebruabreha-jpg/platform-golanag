# RedistributionNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Header** | Pointer to [**RedistributionNotificationHeader**](RedistributionNotificationHeader.md) |  | [optional] 
**Routes** | Pointer to [**[]RedistributionNotificationRoute**](RedistributionNotificationRoute.md) |  | [optional] 

## Methods

### NewRedistributionNotification

`func NewRedistributionNotification() *RedistributionNotification`

NewRedistributionNotification instantiates a new RedistributionNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedistributionNotificationWithDefaults

`func NewRedistributionNotificationWithDefaults() *RedistributionNotification`

NewRedistributionNotificationWithDefaults instantiates a new RedistributionNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RedistributionNotification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RedistributionNotification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RedistributionNotification) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RedistributionNotification) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHeader

`func (o *RedistributionNotification) GetHeader() RedistributionNotificationHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *RedistributionNotification) GetHeaderOk() (*RedistributionNotificationHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *RedistributionNotification) SetHeader(v RedistributionNotificationHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *RedistributionNotification) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetRoutes

`func (o *RedistributionNotification) GetRoutes() []RedistributionNotificationRoute`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *RedistributionNotification) GetRoutesOk() (*[]RedistributionNotificationRoute, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *RedistributionNotification) SetRoutes(v []RedistributionNotificationRoute)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *RedistributionNotification) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



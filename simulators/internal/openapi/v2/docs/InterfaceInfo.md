# InterfaceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceName** | Pointer to **string** | For subnet route, needs to carry interface name for display. for other routes, interface name is not needed. | [optional] 
**InterfaceType** | Pointer to [**InterfaceType**](InterfaceType.md) |  | [optional] 

## Methods

### NewInterfaceInfo

`func NewInterfaceInfo() *InterfaceInfo`

NewInterfaceInfo instantiates a new InterfaceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceInfoWithDefaults

`func NewInterfaceInfoWithDefaults() *InterfaceInfo`

NewInterfaceInfoWithDefaults instantiates a new InterfaceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceName

`func (o *InterfaceInfo) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *InterfaceInfo) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *InterfaceInfo) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *InterfaceInfo) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetInterfaceType

`func (o *InterfaceInfo) GetInterfaceType() InterfaceType`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *InterfaceInfo) GetInterfaceTypeOk() (*InterfaceType, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *InterfaceInfo) SetInterfaceType(v InterfaceType)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *InterfaceInfo) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



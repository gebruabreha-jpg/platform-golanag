# Routes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to [**ClientIdentifier**](ClientIdentifier.md) |  | [optional] 
**Routes** | Pointer to [**[]RouteArrayItem**](RouteArrayItem.md) |  | [optional] 

## Methods

### NewRoutes

`func NewRoutes() *Routes`

NewRoutes instantiates a new Routes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutesWithDefaults

`func NewRoutesWithDefaults() *Routes`

NewRoutesWithDefaults instantiates a new Routes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *Routes) GetClientId() ClientIdentifier`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Routes) GetClientIdOk() (*ClientIdentifier, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Routes) SetClientId(v ClientIdentifier)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Routes) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetRoutes

`func (o *Routes) GetRoutes() []RouteArrayItem`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *Routes) GetRoutesOk() (*[]RouteArrayItem, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *Routes) SetRoutes(v []RouteArrayItem)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *Routes) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



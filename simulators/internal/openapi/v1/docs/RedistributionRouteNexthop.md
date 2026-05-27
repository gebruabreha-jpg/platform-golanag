# RedistributionRouteNexthop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamily** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**IpAddress**](IpAddress.md) |  | [optional] 
**Label** | Pointer to **int32** |  | [optional] 

## Methods

### NewRedistributionRouteNexthop

`func NewRedistributionRouteNexthop() *RedistributionRouteNexthop`

NewRedistributionRouteNexthop instantiates a new RedistributionRouteNexthop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedistributionRouteNexthopWithDefaults

`func NewRedistributionRouteNexthopWithDefaults() *RedistributionRouteNexthop`

NewRedistributionRouteNexthopWithDefaults instantiates a new RedistributionRouteNexthop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamily

`func (o *RedistributionRouteNexthop) GetAddressFamily() string`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *RedistributionRouteNexthop) GetAddressFamilyOk() (*string, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *RedistributionRouteNexthop) SetAddressFamily(v string)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *RedistributionRouteNexthop) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetAddress

`func (o *RedistributionRouteNexthop) GetAddress() IpAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RedistributionRouteNexthop) GetAddressOk() (*IpAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RedistributionRouteNexthop) SetAddress(v IpAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *RedistributionRouteNexthop) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLabel

`func (o *RedistributionRouteNexthop) GetLabel() int32`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RedistributionRouteNexthop) GetLabelOk() (*int32, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RedistributionRouteNexthop) SetLabel(v int32)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RedistributionRouteNexthop) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



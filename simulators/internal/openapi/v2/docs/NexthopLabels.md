# NexthopLabels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpnLabel** | Pointer to **int32** | Nexthop vpn label, when attr &#x3D; vpn_label of nexthopInformation, the field should be filled, indicates private network routing label | [optional] 
**Label** | Pointer to **int32** | Nexthop label, When attr &#x3D; label of nexthopInformation, the field should be filled, indicates public network routing label | [optional] 

## Methods

### NewNexthopLabels

`func NewNexthopLabels() *NexthopLabels`

NewNexthopLabels instantiates a new NexthopLabels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexthopLabelsWithDefaults

`func NewNexthopLabelsWithDefaults() *NexthopLabels`

NewNexthopLabelsWithDefaults instantiates a new NexthopLabels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpnLabel

`func (o *NexthopLabels) GetVpnLabel() int32`

GetVpnLabel returns the VpnLabel field if non-nil, zero value otherwise.

### GetVpnLabelOk

`func (o *NexthopLabels) GetVpnLabelOk() (*int32, bool)`

GetVpnLabelOk returns a tuple with the VpnLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnLabel

`func (o *NexthopLabels) SetVpnLabel(v int32)`

SetVpnLabel sets VpnLabel field to given value.

### HasVpnLabel

`func (o *NexthopLabels) HasVpnLabel() bool`

HasVpnLabel returns a boolean if a field has been set.

### GetLabel

`func (o *NexthopLabels) GetLabel() int32`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *NexthopLabels) GetLabelOk() (*int32, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *NexthopLabels) SetLabel(v int32)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *NexthopLabels) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



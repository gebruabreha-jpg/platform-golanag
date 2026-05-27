# NexthopInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInstance** | **string** | Network instance. | 
**Afi** | **string** | Address family. | 
**Attr** | Pointer to **[]string** | Attributes of the nexthop of the route. | [optional] 
**NhAddr** | [**IpAddress**](IpAddress.md) |  | 
**IfInfo** | Pointer to [**InterfaceInfo**](InterfaceInfo.md) |  | [optional] 
**Labels** | Pointer to [**NexthopLabels**](NexthopLabels.md) |  | [optional] 

## Methods

### NewNexthopInformation

`func NewNexthopInformation(networkInstance string, afi string, nhAddr IpAddress, ) *NexthopInformation`

NewNexthopInformation instantiates a new NexthopInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexthopInformationWithDefaults

`func NewNexthopInformationWithDefaults() *NexthopInformation`

NewNexthopInformationWithDefaults instantiates a new NexthopInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkInstance

`func (o *NexthopInformation) GetNetworkInstance() string`

GetNetworkInstance returns the NetworkInstance field if non-nil, zero value otherwise.

### GetNetworkInstanceOk

`func (o *NexthopInformation) GetNetworkInstanceOk() (*string, bool)`

GetNetworkInstanceOk returns a tuple with the NetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstance

`func (o *NexthopInformation) SetNetworkInstance(v string)`

SetNetworkInstance sets NetworkInstance field to given value.


### GetAfi

`func (o *NexthopInformation) GetAfi() string`

GetAfi returns the Afi field if non-nil, zero value otherwise.

### GetAfiOk

`func (o *NexthopInformation) GetAfiOk() (*string, bool)`

GetAfiOk returns a tuple with the Afi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfi

`func (o *NexthopInformation) SetAfi(v string)`

SetAfi sets Afi field to given value.


### GetAttr

`func (o *NexthopInformation) GetAttr() []string`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *NexthopInformation) GetAttrOk() (*[]string, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *NexthopInformation) SetAttr(v []string)`

SetAttr sets Attr field to given value.

### HasAttr

`func (o *NexthopInformation) HasAttr() bool`

HasAttr returns a boolean if a field has been set.

### GetNhAddr

`func (o *NexthopInformation) GetNhAddr() IpAddress`

GetNhAddr returns the NhAddr field if non-nil, zero value otherwise.

### GetNhAddrOk

`func (o *NexthopInformation) GetNhAddrOk() (*IpAddress, bool)`

GetNhAddrOk returns a tuple with the NhAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNhAddr

`func (o *NexthopInformation) SetNhAddr(v IpAddress)`

SetNhAddr sets NhAddr field to given value.


### GetIfInfo

`func (o *NexthopInformation) GetIfInfo() InterfaceInfo`

GetIfInfo returns the IfInfo field if non-nil, zero value otherwise.

### GetIfInfoOk

`func (o *NexthopInformation) GetIfInfoOk() (*InterfaceInfo, bool)`

GetIfInfoOk returns a tuple with the IfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfInfo

`func (o *NexthopInformation) SetIfInfo(v InterfaceInfo)`

SetIfInfo sets IfInfo field to given value.

### HasIfInfo

`func (o *NexthopInformation) HasIfInfo() bool`

HasIfInfo returns a boolean if a field has been set.

### GetLabels

`func (o *NexthopInformation) GetLabels() NexthopLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NexthopInformation) GetLabelsOk() (*NexthopLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NexthopInformation) SetLabels(v NexthopLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *NexthopInformation) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



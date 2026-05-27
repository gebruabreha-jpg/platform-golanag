# Redistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to [**ClientIdentifier**](ClientIdentifier.md) |  | [optional] 
**AddressFamily** | Pointer to **string** |  | [optional] 
**ProtocolType** | Pointer to [**RouteType**](RouteType.md) |  | [optional] 
**SrcTag** | Pointer to **string** |  | [optional] 
**DestProtocolType** | Pointer to **string** |  | [optional] 
**DestTag** | Pointer to **string** |  | [optional] 

## Methods

### NewRedistribution

`func NewRedistribution() *Redistribution`

NewRedistribution instantiates a new Redistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedistributionWithDefaults

`func NewRedistributionWithDefaults() *Redistribution`

NewRedistributionWithDefaults instantiates a new Redistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *Redistribution) GetClientId() ClientIdentifier`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Redistribution) GetClientIdOk() (*ClientIdentifier, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Redistribution) SetClientId(v ClientIdentifier)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Redistribution) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetAddressFamily

`func (o *Redistribution) GetAddressFamily() string`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *Redistribution) GetAddressFamilyOk() (*string, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *Redistribution) SetAddressFamily(v string)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *Redistribution) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetProtocolType

`func (o *Redistribution) GetProtocolType() RouteType`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *Redistribution) GetProtocolTypeOk() (*RouteType, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *Redistribution) SetProtocolType(v RouteType)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *Redistribution) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetSrcTag

`func (o *Redistribution) GetSrcTag() string`

GetSrcTag returns the SrcTag field if non-nil, zero value otherwise.

### GetSrcTagOk

`func (o *Redistribution) GetSrcTagOk() (*string, bool)`

GetSrcTagOk returns a tuple with the SrcTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcTag

`func (o *Redistribution) SetSrcTag(v string)`

SetSrcTag sets SrcTag field to given value.

### HasSrcTag

`func (o *Redistribution) HasSrcTag() bool`

HasSrcTag returns a boolean if a field has been set.

### GetDestProtocolType

`func (o *Redistribution) GetDestProtocolType() string`

GetDestProtocolType returns the DestProtocolType field if non-nil, zero value otherwise.

### GetDestProtocolTypeOk

`func (o *Redistribution) GetDestProtocolTypeOk() (*string, bool)`

GetDestProtocolTypeOk returns a tuple with the DestProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestProtocolType

`func (o *Redistribution) SetDestProtocolType(v string)`

SetDestProtocolType sets DestProtocolType field to given value.

### HasDestProtocolType

`func (o *Redistribution) HasDestProtocolType() bool`

HasDestProtocolType returns a boolean if a field has been set.

### GetDestTag

`func (o *Redistribution) GetDestTag() string`

GetDestTag returns the DestTag field if non-nil, zero value otherwise.

### GetDestTagOk

`func (o *Redistribution) GetDestTagOk() (*string, bool)`

GetDestTagOk returns a tuple with the DestTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestTag

`func (o *Redistribution) SetDestTag(v string)`

SetDestTag sets DestTag field to given value.

### HasDestTag

`func (o *Redistribution) HasDestTag() bool`

HasDestTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



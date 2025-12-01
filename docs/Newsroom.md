# Newsroom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Company symbol. | [optional] 
**Data** | Pointer to [**[]NewsroomArticle**](NewsroomArticle.md) | Array of articles. | [optional] 

## Methods

### NewNewsroom

`func NewNewsroom() *Newsroom`

NewNewsroom instantiates a new Newsroom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewsroomWithDefaults

`func NewNewsroomWithDefaults() *Newsroom`

NewNewsroomWithDefaults instantiates a new Newsroom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *Newsroom) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Newsroom) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Newsroom) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Newsroom) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetData

`func (o *Newsroom) GetData() []NewsroomArticle`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Newsroom) GetDataOk() (*[]NewsroomArticle, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Newsroom) SetData(v []NewsroomArticle)`

SetData sets Data field to given value.

### HasData

`func (o *Newsroom) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



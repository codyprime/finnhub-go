# NewsroomArticle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtDate** | Pointer to **string** | Published time in &lt;code&gt;YYYY-MM-DD HH:MM:SS&lt;/code&gt; format (EST timezone). | [optional] 
**Title** | Pointer to **string** | Title. | [optional] 
**FullText** | Pointer to **string** | URL to download the full text data. | [optional] 
**Url** | Pointer to **string** | Original URL. | [optional] 

## Methods

### NewNewsroomArticle

`func NewNewsroomArticle() *NewsroomArticle`

NewNewsroomArticle instantiates a new NewsroomArticle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewsroomArticleWithDefaults

`func NewNewsroomArticleWithDefaults() *NewsroomArticle`

NewNewsroomArticleWithDefaults instantiates a new NewsroomArticle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtDate

`func (o *NewsroomArticle) GetAtDate() string`

GetAtDate returns the AtDate field if non-nil, zero value otherwise.

### GetAtDateOk

`func (o *NewsroomArticle) GetAtDateOk() (*string, bool)`

GetAtDateOk returns a tuple with the AtDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtDate

`func (o *NewsroomArticle) SetAtDate(v string)`

SetAtDate sets AtDate field to given value.

### HasAtDate

`func (o *NewsroomArticle) HasAtDate() bool`

HasAtDate returns a boolean if a field has been set.

### GetTitle

`func (o *NewsroomArticle) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewsroomArticle) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewsroomArticle) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NewsroomArticle) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFullText

`func (o *NewsroomArticle) GetFullText() string`

GetFullText returns the FullText field if non-nil, zero value otherwise.

### GetFullTextOk

`func (o *NewsroomArticle) GetFullTextOk() (*string, bool)`

GetFullTextOk returns a tuple with the FullText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullText

`func (o *NewsroomArticle) SetFullText(v string)`

SetFullText sets FullText field to given value.

### HasFullText

`func (o *NewsroomArticle) HasFullText() bool`

HasFullText returns a boolean if a field has been set.

### GetUrl

`func (o *NewsroomArticle) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NewsroomArticle) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NewsroomArticle) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NewsroomArticle) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



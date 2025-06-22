# ReviewsApi

All URIs are relative to *http://localhost:8080*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**apiV1ReviewsGet**](#apiv1reviewsget) | **GET** /api/v1/reviews | レビュー一覧取得|
|[**apiV1ReviewsIdDelete**](#apiv1reviewsiddelete) | **DELETE** /api/v1/reviews/{id} | レビュー削除|
|[**apiV1ReviewsIdGet**](#apiv1reviewsidget) | **GET** /api/v1/reviews/{id} | レビュー詳細取得|
|[**apiV1ReviewsIdImagesPost**](#apiv1reviewsidimagespost) | **POST** /api/v1/reviews/{id}/images | レビュー画像アップロード|
|[**apiV1ReviewsIdPut**](#apiv1reviewsidput) | **PUT** /api/v1/reviews/{id} | レビュー更新|
|[**apiV1ReviewsPost**](#apiv1reviewspost) | **POST** /api/v1/reviews | レビュー投稿|

# **apiV1ReviewsGet**
> Array<Review> apiV1ReviewsGet()

全レビューの一覧を取得

### Example

```typescript
import {
    ReviewsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ReviewsApi(configuration);

let limit: number; //取得件数制限 (optional) (default to undefined)
let offset: number; //取得開始位置 (optional) (default to undefined)
let after: string; // (optional) (default to undefined)
let before: string; //指定日時以降のレビューを取得 (optional) (default to undefined)
let shopId: string; //店舗IDでフィルタ (optional) (default to undefined)
let authorId: string; //投稿者IDでフィルタ (optional) (default to undefined)

const { status, data } = await apiInstance.apiV1ReviewsGet(
    limit,
    offset,
    after,
    before,
    shopId,
    authorId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | [**number**] | 取得件数制限 | (optional) defaults to undefined|
| **offset** | [**number**] | 取得開始位置 | (optional) defaults to undefined|
| **after** | [**string**] |  | (optional) defaults to undefined|
| **before** | [**string**] | 指定日時以降のレビューを取得 | (optional) defaults to undefined|
| **shopId** | [**string**] | 店舗IDでフィルタ | (optional) defaults to undefined|
| **authorId** | [**string**] | 投稿者IDでフィルタ | (optional) defaults to undefined|


### Return type

**Array<Review>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | レビュー一覧 |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1ReviewsIdDelete**
> apiV1ReviewsIdDelete()

指定されたIDのレビューを削除

### Example

```typescript
import {
    ReviewsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ReviewsApi(configuration);

let id: string; //レビューID (default to undefined)

const { status, data } = await apiInstance.apiV1ReviewsIdDelete(
    id
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **id** | [**string**] | レビューID | defaults to undefined|


### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**204** | レビューの削除に成功 |  -  |
|**404** | レビューが見つかりません |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1ReviewsIdGet**
> Review apiV1ReviewsIdGet()

指定されたIDのレビュー詳細を取得

### Example

```typescript
import {
    ReviewsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ReviewsApi(configuration);

let id: string; //レビューID (default to undefined)

const { status, data } = await apiInstance.apiV1ReviewsIdGet(
    id
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **id** | [**string**] | レビューID | defaults to undefined|


### Return type

**Review**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | レビュー詳細 |  -  |
|**404** | レビューが見つかりません |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1ReviewsIdImagesPost**
> ApiV1ShopsIdImagesPost200Response apiV1ReviewsIdImagesPost()

指定されたレビューに画像をアップロード

### Example

```typescript
import {
    ReviewsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ReviewsApi(configuration);

let id: string; //レビューID (default to undefined)
let image: File; // (default to undefined)

const { status, data } = await apiInstance.apiV1ReviewsIdImagesPost(
    id,
    image
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **id** | [**string**] | レビューID | defaults to undefined|
| **image** | [**File**] |  | defaults to undefined|


### Return type

**ApiV1ShopsIdImagesPost200Response**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | 画像のアップロードに成功 |  -  |
|**404** | レビューが見つかりません |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1ReviewsIdPut**
> Review apiV1ReviewsIdPut(review)

指定されたIDのレビューを更新

### Example

```typescript
import {
    ReviewsApi,
    Configuration,
    Review
} from './api';

const configuration = new Configuration();
const apiInstance = new ReviewsApi(configuration);

let id: string; //レビューID (default to undefined)
let review: Review; //

const { status, data } = await apiInstance.apiV1ReviewsIdPut(
    id,
    review
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **review** | **Review**|  | |
| **id** | [**string**] | レビューID | defaults to undefined|


### Return type

**Review**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | 更新されたレビュー |  -  |
|**404** | レビューが見つかりません |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1ReviewsPost**
> Review apiV1ReviewsPost(apiV1ReviewsPostRequest)

新規レビューを投稿

### Example

```typescript
import {
    ReviewsApi,
    Configuration,
    ApiV1ReviewsPostRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new ReviewsApi(configuration);

let apiV1ReviewsPostRequest: ApiV1ReviewsPostRequest; //

const { status, data } = await apiInstance.apiV1ReviewsPost(
    apiV1ReviewsPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **apiV1ReviewsPostRequest** | **ApiV1ReviewsPostRequest**|  | |


### Return type

**Review**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | 作成されたレビュー |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


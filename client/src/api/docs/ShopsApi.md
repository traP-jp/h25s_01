# ShopsApi

All URIs are relative to *http://localhost:8080*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**apiV1ShopsGet**](#apiv1shopsget) | **GET** /api/v1/shops | 店舗一覧取得|
|[**apiV1ShopsIdDelete**](#apiv1shopsiddelete) | **DELETE** /api/v1/shops/{id} | 店舗削除|
|[**apiV1ShopsIdGet**](#apiv1shopsidget) | **GET** /api/v1/shops/{id} | 店舗詳細取得|
|[**apiV1ShopsIdImagesDelete**](#apiv1shopsidimagesdelete) | **DELETE** /api/v1/shops/{id}/images | 店舗画像削除|
|[**apiV1ShopsIdImagesPost**](#apiv1shopsidimagespost) | **POST** /api/v1/shops/{id}/images | 店舗画像アップロード|
|[**apiV1ShopsIdPut**](#apiv1shopsidput) | **PUT** /api/v1/shops/{id} | 店舗情報更新|
|[**apiV1ShopsPost**](#apiv1shopspost) | **POST** /api/v1/shops | 店舗作成|

# **apiV1ShopsGet**
> Array<Shop> apiV1ShopsGet()

全店舗の一覧を取得します

### Example

```typescript
import {
    ShopsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ShopsApi(configuration);

let limit: number; //取得件数制限 (optional) (default to undefined)
let latitude: string; //緯度（位置検索用） (optional) (default to undefined)
let longitude: string; //経度（位置検索用） (optional) (default to undefined)
let radius: number; //検索半径（km） (optional) (default to undefined)
let stationId: string; //駅IDでフィルタ (optional) (default to undefined)

const { status, data } = await apiInstance.apiV1ShopsGet(
    limit,
    latitude,
    longitude,
    radius,
    stationId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | [**number**] | 取得件数制限 | (optional) defaults to undefined|
| **latitude** | [**string**] | 緯度（位置検索用） | (optional) defaults to undefined|
| **longitude** | [**string**] | 経度（位置検索用） | (optional) defaults to undefined|
| **radius** | [**number**] | 検索半径（km） | (optional) defaults to undefined|
| **stationId** | [**string**] | 駅IDでフィルタ | (optional) defaults to undefined|


### Return type

**Array<Shop>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | 店舗一覧 |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1ShopsIdDelete**
> apiV1ShopsIdDelete()

指定されたIDの店舗を削除

### Example

```typescript
import {
    ShopsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ShopsApi(configuration);

let id: string; //店舗ID (default to undefined)

const { status, data } = await apiInstance.apiV1ShopsIdDelete(
    id
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **id** | [**string**] | 店舗ID | defaults to undefined|


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
|**204** | 店舗の削除に成功 |  -  |
|**404** | 店舗が見つかりません |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1ShopsIdGet**
> Shop apiV1ShopsIdGet()

指定されたIDの店舗詳細情報を取得

### Example

```typescript
import {
    ShopsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ShopsApi(configuration);

let id: string; //店舗ID (default to undefined)

const { status, data } = await apiInstance.apiV1ShopsIdGet(
    id
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **id** | [**string**] | 店舗ID | defaults to undefined|


### Return type

**Shop**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | 店舗詳細 |  -  |
|**404** | 店舗が見つかりません |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1ShopsIdImagesDelete**
> apiV1ShopsIdImagesDelete(apiV1ShopsIdImagesDeleteRequest)

指定された店舗の画像を削除

### Example

```typescript
import {
    ShopsApi,
    Configuration,
    ApiV1ShopsIdImagesDeleteRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new ShopsApi(configuration);

let id: string; //店舗ID (default to undefined)
let apiV1ShopsIdImagesDeleteRequest: ApiV1ShopsIdImagesDeleteRequest; //

const { status, data } = await apiInstance.apiV1ShopsIdImagesDelete(
    id,
    apiV1ShopsIdImagesDeleteRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **apiV1ShopsIdImagesDeleteRequest** | **ApiV1ShopsIdImagesDeleteRequest**|  | |
| **id** | [**string**] | 店舗ID | defaults to undefined|


### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**204** | 画像の削除に成功 |  -  |
|**404** | 店舗または画像が見つかりません |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1ShopsIdImagesPost**
> ApiV1ShopsIdImagesPost200Response apiV1ShopsIdImagesPost()

指定された店舗に画像をアップロード

### Example

```typescript
import {
    ShopsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ShopsApi(configuration);

let id: string; //店舗ID (default to undefined)
let image: File; // (default to undefined)

const { status, data } = await apiInstance.apiV1ShopsIdImagesPost(
    id,
    image
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **id** | [**string**] | 店舗ID | defaults to undefined|
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
|**404** | 店舗が見つかりません |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1ShopsIdPut**
> Shop apiV1ShopsIdPut(shop)

指定されたIDの店舗情報を更新

### Example

```typescript
import {
    ShopsApi,
    Configuration,
    Shop
} from './api';

const configuration = new Configuration();
const apiInstance = new ShopsApi(configuration);

let id: string; //店舗ID (default to undefined)
let shop: Shop; //

const { status, data } = await apiInstance.apiV1ShopsIdPut(
    id,
    shop
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **shop** | **Shop**|  | |
| **id** | [**string**] | 店舗ID | defaults to undefined|


### Return type

**Shop**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | 更新された店舗 |  -  |
|**404** | 店舗が見つかりません |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1ShopsPost**
> Shop apiV1ShopsPost(apiV1ShopsPostRequest)

新規店舗を登録します

### Example

```typescript
import {
    ShopsApi,
    Configuration,
    ApiV1ShopsPostRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new ShopsApi(configuration);

let apiV1ShopsPostRequest: ApiV1ShopsPostRequest; //

const { status, data } = await apiInstance.apiV1ShopsPost(
    apiV1ShopsPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **apiV1ShopsPostRequest** | **ApiV1ShopsPostRequest**|  | |


### Return type

**Shop**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | 作成された店舗 |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


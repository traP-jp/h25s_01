# StationsApi

All URIs are relative to *http://localhost:8080*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**apiV1StationsGet**](#apiv1stationsget) | **GET** /api/v1/stations | 駅一覧取得|
|[**apiV1StationsIdDelete**](#apiv1stationsiddelete) | **DELETE** /api/v1/stations/{id} | 駅削除|
|[**apiV1StationsIdGet**](#apiv1stationsidget) | **GET** /api/v1/stations/{id} | 駅詳細取得|
|[**apiV1StationsIdPut**](#apiv1stationsidput) | **PUT** /api/v1/stations/{id} | 駅情報更新|
|[**apiV1StationsIdShopsGet**](#apiv1stationsidshopsget) | **GET** /api/v1/stations/{id}/shops | 駅周辺の店舗取得|
|[**apiV1StationsPost**](#apiv1stationspost) | **POST** /api/v1/stations | 駅作成|

# **apiV1StationsGet**
> Array<Station> apiV1StationsGet()

全駅の一覧を取得します

### Example

```typescript
import {
    StationsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new StationsApi(configuration);

let limit: number; //取得件数制限 (optional) (default to undefined)
let name: string; //駅名で部分一致検索 (optional) (default to undefined)

const { status, data } = await apiInstance.apiV1StationsGet(
    limit,
    name
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | [**number**] | 取得件数制限 | (optional) defaults to undefined|
| **name** | [**string**] | 駅名で部分一致検索 | (optional) defaults to undefined|


### Return type

**Array<Station>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | 駅一覧 |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1StationsIdDelete**
> apiV1StationsIdDelete()

指定されたIDの駅を削除します

### Example

```typescript
import {
    StationsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new StationsApi(configuration);

let id: string; //駅ID (default to undefined)

const { status, data } = await apiInstance.apiV1StationsIdDelete(
    id
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **id** | [**string**] | 駅ID | defaults to undefined|


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
|**204** | 駅の削除に成功 |  -  |
|**404** | 駅が見つかりません |  -  |
|**409** | この駅は店舗で使用されているため削除できません |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1StationsIdGet**
> Station apiV1StationsIdGet()

指定されたIDの駅詳細情報を取得

### Example

```typescript
import {
    StationsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new StationsApi(configuration);

let id: string; //駅ID (default to undefined)

const { status, data } = await apiInstance.apiV1StationsIdGet(
    id
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **id** | [**string**] | 駅ID | defaults to undefined|


### Return type

**Station**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | 駅詳細 |  -  |
|**404** | 駅が見つかりません |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1StationsIdPut**
> Station apiV1StationsIdPut(apiV1StationsPostRequest)

指定されたIDの駅情報を更新

### Example

```typescript
import {
    StationsApi,
    Configuration,
    ApiV1StationsPostRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new StationsApi(configuration);

let id: string; //駅ID (default to undefined)
let apiV1StationsPostRequest: ApiV1StationsPostRequest; //

const { status, data } = await apiInstance.apiV1StationsIdPut(
    id,
    apiV1StationsPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **apiV1StationsPostRequest** | **ApiV1StationsPostRequest**|  | |
| **id** | [**string**] | 駅ID | defaults to undefined|


### Return type

**Station**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | 更新された駅 |  -  |
|**404** | 駅が見つかりません |  -  |
|**400** | 無効なリクエスト |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1StationsIdShopsGet**
> Array<Shop> apiV1StationsIdShopsGet()

指定された駅周辺の店舗一覧を取得します

### Example

```typescript
import {
    StationsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new StationsApi(configuration);

let id: string; //駅ID (default to undefined)
let limit: number; //取得件数制限 (optional) (default to undefined)

const { status, data } = await apiInstance.apiV1StationsIdShopsGet(
    id,
    limit
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **id** | [**string**] | 駅ID | defaults to undefined|
| **limit** | [**number**] | 取得件数制限 | (optional) defaults to undefined|


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
|**200** | 駅周辺の店舗一覧 |  -  |
|**404** | 駅が見つかりません |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiV1StationsPost**
> Station apiV1StationsPost(apiV1StationsPostRequest)

新規駅を登録します

### Example

```typescript
import {
    StationsApi,
    Configuration,
    ApiV1StationsPostRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new StationsApi(configuration);

let apiV1StationsPostRequest: ApiV1StationsPostRequest; //

const { status, data } = await apiInstance.apiV1StationsPost(
    apiV1StationsPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **apiV1StationsPostRequest** | **ApiV1StationsPostRequest**|  | |


### Return type

**Station**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | 作成された駅 |  -  |
|**400** | 無効なリクエスト |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


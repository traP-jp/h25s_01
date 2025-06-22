# ApiV1ShopsPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** |  | [default to undefined]
**post_code** | **string** |  | [optional] [default to undefined]
**address** | **string** |  | [default to undefined]
**latitude** | **string** |  | [optional] [default to undefined]
**longitude** | **string** |  | [optional] [default to undefined]
**images** | **Array&lt;string&gt;** |  | [optional] [default to undefined]
**payment_methods** | **Array&lt;string&gt;** |  | [optional] [default to undefined]
**stations** | **Array&lt;string&gt;** | 関連する駅のID配列 | [optional] [default to undefined]
**registerer** | **string** |  | [optional] [default to undefined]

## Example

```typescript
import { ApiV1ShopsPostRequest } from './api';

const instance: ApiV1ShopsPostRequest = {
    name,
    post_code,
    address,
    latitude,
    longitude,
    images,
    payment_methods,
    stations,
    registerer,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

# Review


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** |  | [default to undefined]
**author** | **string** | レビュー投稿者のユーザーID | [default to undefined]
**shop** | **string** | レビュー対象の店舗ID | [default to undefined]
**rating** | **number** | 評価（0から3まで） | [optional] [default to undefined]
**created_at** | **string** |  | [optional] [default to undefined]
**updated_at** | **string** |  | [optional] [default to undefined]
**content** | **string** |  | [optional] [default to undefined]
**images** | **Array&lt;string&gt;** |  | [optional] [default to undefined]

## Example

```typescript
import { Review } from './api';

const instance: Review = {
    id,
    author,
    shop,
    rating,
    created_at,
    updated_at,
    content,
    images,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

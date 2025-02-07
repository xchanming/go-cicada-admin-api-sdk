package go_cicada_admin_sdk

import (
	"net/http"
)

type ProductKeywordDictionaryRepository ClientService

func (t ProductKeywordDictionaryRepository) Search(ctx ApiContext, criteria Criteria) (*ProductKeywordDictionaryCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-keyword-dictionary", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductKeywordDictionaryCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductKeywordDictionaryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductKeywordDictionaryCollection, *http.Response, error) {
	if criteria.Limit == 0 {
		criteria.Limit = 50
	}

	if criteria.Page == 0 {
		criteria.Page = 1
	}

	c, resp, err := t.Search(ctx, criteria)

	if err != nil {
		return c, resp, err
	}

	for {
		criteria.Page++

		nextC, nextResp, nextErr := t.Search(ctx, criteria)

		if nextErr != nil {
			return c, nextResp, nextErr
		}

		if len(nextC.Data) == 0 {
			break
		}

		c.Data = append(c.Data, nextC.Data...)
	}

	c.Total = int64(len(c.Data))

	return c, resp, err
}

func (t ProductKeywordDictionaryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-keyword-dictionary", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SearchIdsResponse)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductKeywordDictionaryRepository) Upsert(ctx ApiContext, entity []ProductKeywordDictionary) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_keyword_dictionary": {
		Entity:  "product_keyword_dictionary",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductKeywordDictionaryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_keyword_dictionary": {
		Entity:  "product_keyword_dictionary",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductKeywordDictionary struct {
	Id string `json:"id,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Keyword string `json:"keyword,omitempty"`

	Reversed string `json:"reversed,omitempty"`

	Language *Language `json:"language,omitempty"`
}

type ProductKeywordDictionaryCollection struct {
	EntityCollection

	Data []ProductKeywordDictionary `json:"data"`
}

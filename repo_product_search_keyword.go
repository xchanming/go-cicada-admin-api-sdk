package go_cicada_admin_sdk

import (
	"net/http"

	"time"
)

type ProductSearchKeywordRepository ClientService

func (t ProductSearchKeywordRepository) Search(ctx ApiContext, criteria Criteria) (*ProductSearchKeywordCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-search-keyword", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductSearchKeywordCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductSearchKeywordRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductSearchKeywordCollection, *http.Response, error) {
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

func (t ProductSearchKeywordRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-search-keyword", criteria)

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

func (t ProductSearchKeywordRepository) Upsert(ctx ApiContext, entity []ProductSearchKeyword) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_search_keyword": {
		Entity:  "product_search_keyword",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductSearchKeywordRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_search_keyword": {
		Entity:  "product_search_keyword",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductSearchKeyword struct {
	Keyword string `json:"keyword,omitempty"`

	Product *Product `json:"product,omitempty"`

	Language *Language `json:"language,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	Ranking float64 `json:"ranking,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	ProductId string `json:"productId,omitempty"`
}

type ProductSearchKeywordCollection struct {
	EntityCollection

	Data []ProductSearchKeyword `json:"data"`
}

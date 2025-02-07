package go_cicada_admin_sdk

import (
	"net/http"

	"time"
)

type ProductMediaRepository ClientService

func (t ProductMediaRepository) Search(ctx ApiContext, criteria Criteria) (*ProductMediaCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-media", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductMediaCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductMediaRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ProductMediaCollection, *http.Response, error) {
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

func (t ProductMediaRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-media", criteria)

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

func (t ProductMediaRepository) Upsert(ctx ApiContext, entity []ProductMedia) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_media": {
		Entity:  "product_media",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductMediaRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_media": {
		Entity:  "product_media",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductMedia struct {
	Media *Media `json:"media,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ProductId string `json:"productId,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	Position float64 `json:"position,omitempty"`

	Product *Product `json:"product,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	VersionId string `json:"versionId,omitempty"`
}

type ProductMediaCollection struct {
	EntityCollection

	Data []ProductMedia `json:"data"`
}

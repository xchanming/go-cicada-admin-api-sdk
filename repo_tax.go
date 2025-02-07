package go_cicada_admin_sdk

import (
	"net/http"

	"time"
)

type TaxRepository ClientService

func (t TaxRepository) Search(ctx ApiContext, criteria Criteria) (*TaxCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/tax", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(TaxCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t TaxRepository) SearchAll(ctx ApiContext, criteria Criteria) (*TaxCollection, *http.Response, error) {
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

func (t TaxRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/tax", criteria)

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

func (t TaxRepository) Upsert(ctx ApiContext, entity []Tax) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"tax": {
		Entity:  "tax",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t TaxRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"tax": {
		Entity:  "tax",
		Action:  "delete",
		Payload: payload,
	}})
}

type Tax struct {
	Id string `json:"id,omitempty"`

	TaxRate float64 `json:"taxRate,omitempty"`

	Position float64 `json:"position,omitempty"`

	Rules []TaxRule `json:"rules,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Name string `json:"name,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Products []Product `json:"products,omitempty"`

	ShippingMethods []ShippingMethod `json:"shippingMethods,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type TaxCollection struct {
	EntityCollection

	Data []Tax `json:"data"`
}

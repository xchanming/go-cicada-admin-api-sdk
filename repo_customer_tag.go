package go_cicada_admin_sdk

import (
	"net/http"
)

type CustomerTagRepository ClientService

func (t CustomerTagRepository) Search(ctx ApiContext, criteria Criteria) (*CustomerTagCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/customer-tag", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CustomerTagCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CustomerTagRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CustomerTagCollection, *http.Response, error) {
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

func (t CustomerTagRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/customer-tag", criteria)

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

func (t CustomerTagRepository) Upsert(ctx ApiContext, entity []CustomerTag) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer_tag": {
		Entity:  "customer_tag",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CustomerTagRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"customer_tag": {
		Entity:  "customer_tag",
		Action:  "delete",
		Payload: payload,
	}})
}

type CustomerTag struct {
	TagId string `json:"tagId,omitempty"`

	Customer *Customer `json:"customer,omitempty"`

	Tag *Tag `json:"tag,omitempty"`

	CustomerId string `json:"customerId,omitempty"`
}

type CustomerTagCollection struct {
	EntityCollection

	Data []CustomerTag `json:"data"`
}
package go_cicada_admin_sdk

import (
	"net/http"

	"time"
)

type PromotionSalesChannelRepository ClientService

func (t PromotionSalesChannelRepository) Search(ctx ApiContext, criteria Criteria) (*PromotionSalesChannelCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/promotion-sales-channel", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(PromotionSalesChannelCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t PromotionSalesChannelRepository) SearchAll(ctx ApiContext, criteria Criteria) (*PromotionSalesChannelCollection, *http.Response, error) {
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

func (t PromotionSalesChannelRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/promotion-sales-channel", criteria)

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

func (t PromotionSalesChannelRepository) Upsert(ctx ApiContext, entity []PromotionSalesChannel) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_sales_channel": {
		Entity:  "promotion_sales_channel",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t PromotionSalesChannelRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"promotion_sales_channel": {
		Entity:  "promotion_sales_channel",
		Action:  "delete",
		Payload: payload,
	}})
}

type PromotionSalesChannel struct {
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Id string `json:"id,omitempty"`

	PromotionId string `json:"promotionId,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`

	Priority float64 `json:"priority,omitempty"`

	Promotion *Promotion `json:"promotion,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type PromotionSalesChannelCollection struct {
	EntityCollection

	Data []PromotionSalesChannel `json:"data"`
}

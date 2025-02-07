package go_cicada_admin_sdk

import (
	"net/http"
)

type SalesChannelPaymentMethodRepository ClientService

func (t SalesChannelPaymentMethodRepository) Search(ctx ApiContext, criteria Criteria) (*SalesChannelPaymentMethodCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/sales-channel-payment-method", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SalesChannelPaymentMethodCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t SalesChannelPaymentMethodRepository) SearchAll(ctx ApiContext, criteria Criteria) (*SalesChannelPaymentMethodCollection, *http.Response, error) {
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

func (t SalesChannelPaymentMethodRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/sales-channel-payment-method", criteria)

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

func (t SalesChannelPaymentMethodRepository) Upsert(ctx ApiContext, entity []SalesChannelPaymentMethod) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_payment_method": {
		Entity:  "sales_channel_payment_method",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t SalesChannelPaymentMethodRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"sales_channel_payment_method": {
		Entity:  "sales_channel_payment_method",
		Action:  "delete",
		Payload: payload,
	}})
}

type SalesChannelPaymentMethod struct {
	SalesChannelId string `json:"salesChannelId,omitempty"`

	PaymentMethodId string `json:"paymentMethodId,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	PaymentMethod *PaymentMethod `json:"paymentMethod,omitempty"`
}

type SalesChannelPaymentMethodCollection struct {
	EntityCollection

	Data []SalesChannelPaymentMethod `json:"data"`
}

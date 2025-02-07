package go_cicada_admin_sdk

import (
	"net/http"

	"time"
)

type OrderCustomerRepository ClientService

func (t OrderCustomerRepository) Search(ctx ApiContext, criteria Criteria) (*OrderCustomerCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/order-customer", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(OrderCustomerCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t OrderCustomerRepository) SearchAll(ctx ApiContext, criteria Criteria) (*OrderCustomerCollection, *http.Response, error) {
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

func (t OrderCustomerRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/order-customer", criteria)

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

func (t OrderCustomerRepository) Upsert(ctx ApiContext, entity []OrderCustomer) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_customer": {
		Entity:  "order_customer",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t OrderCustomerRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"order_customer": {
		Entity:  "order_customer",
		Action:  "delete",
		Payload: payload,
	}})
}

type OrderCustomer struct {
	Email string `json:"email,omitempty"`

	LastName string `json:"lastName,omitempty"`

	CustomerNumber string `json:"customerNumber,omitempty"`

	Customer *Customer `json:"customer,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CustomerId string `json:"customerId,omitempty"`

	Salutation *Salutation `json:"salutation,omitempty"`

	RemoteAddress interface{} `json:"remoteAddress,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	Company string `json:"company,omitempty"`

	SalutationId string `json:"salutationId,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	OrderId string `json:"orderId,omitempty"`

	OrderVersionId string `json:"orderVersionId,omitempty"`

	Title string `json:"title,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Order *Order `json:"order,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`
}

type OrderCustomerCollection struct {
	EntityCollection

	Data []OrderCustomer `json:"data"`
}

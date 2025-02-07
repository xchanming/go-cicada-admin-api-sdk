package go_cicada_admin_sdk

import (
	"net/http"

	"time"
)

type CountryTranslationRepository ClientService

func (t CountryTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*CountryTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/country-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CountryTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CountryTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CountryTranslationCollection, *http.Response, error) {
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

func (t CountryTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/country-translation", criteria)

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

func (t CountryTranslationRepository) Upsert(ctx ApiContext, entity []CountryTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"country_translation": {
		Entity:  "country_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CountryTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"country_translation": {
		Entity:  "country_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type CountryTranslation struct {
	CustomFields interface{} `json:"customFields,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	CountryId string `json:"countryId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Name string `json:"name,omitempty"`

	AddressFormat interface{} `json:"addressFormat,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Country *Country `json:"country,omitempty"`

	Language *Language `json:"language,omitempty"`
}

type CountryTranslationCollection struct {
	EntityCollection

	Data []CountryTranslation `json:"data"`
}

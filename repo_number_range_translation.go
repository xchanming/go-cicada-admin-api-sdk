package go_cicada_admin_sdk

import (
	"net/http"

	"time"
)

type NumberRangeTranslationRepository ClientService

func (t NumberRangeTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*NumberRangeTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/number-range-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(NumberRangeTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t NumberRangeTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*NumberRangeTranslationCollection, *http.Response, error) {
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

func (t NumberRangeTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/number-range-translation", criteria)

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

func (t NumberRangeTranslationRepository) Upsert(ctx ApiContext, entity []NumberRangeTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"number_range_translation": {
		Entity:  "number_range_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t NumberRangeTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"number_range_translation": {
		Entity:  "number_range_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type NumberRangeTranslation struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	NumberRange *NumberRange `json:"numberRange,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	NumberRangeId string `json:"numberRangeId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Language *Language `json:"language,omitempty"`
}

type NumberRangeTranslationCollection struct {
	EntityCollection

	Data []NumberRangeTranslation `json:"data"`
}

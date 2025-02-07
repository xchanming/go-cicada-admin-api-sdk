package go_cicada_admin_sdk

import (
	"net/http"

	"time"
)

type AppActionButtonTranslationRepository ClientService

func (t AppActionButtonTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*AppActionButtonTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/app-action-button-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(AppActionButtonTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t AppActionButtonTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*AppActionButtonTranslationCollection, *http.Response, error) {
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

func (t AppActionButtonTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/app-action-button-translation", criteria)

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

func (t AppActionButtonTranslationRepository) Upsert(ctx ApiContext, entity []AppActionButtonTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_action_button_translation": {
		Entity:  "app_action_button_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t AppActionButtonTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_action_button_translation": {
		Entity:  "app_action_button_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type AppActionButtonTranslation struct {
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	AppActionButtonId string `json:"appActionButtonId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	AppActionButton *AppActionButton `json:"appActionButton,omitempty"`

	Language *Language `json:"language,omitempty"`

	Label string `json:"label,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type AppActionButtonTranslationCollection struct {
	EntityCollection

	Data []AppActionButtonTranslation `json:"data"`
}

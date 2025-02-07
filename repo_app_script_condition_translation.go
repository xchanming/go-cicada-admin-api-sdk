package go_cicada_admin_sdk

import (
	"net/http"

	"time"
)

type AppScriptConditionTranslationRepository ClientService

func (t AppScriptConditionTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*AppScriptConditionTranslationCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/app-script-condition-translation", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(AppScriptConditionTranslationCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t AppScriptConditionTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*AppScriptConditionTranslationCollection, *http.Response, error) {
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

func (t AppScriptConditionTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/app-script-condition-translation", criteria)

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

func (t AppScriptConditionTranslationRepository) Upsert(ctx ApiContext, entity []AppScriptConditionTranslation) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_script_condition_translation": {
		Entity:  "app_script_condition_translation",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t AppScriptConditionTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"app_script_condition_translation": {
		Entity:  "app_script_condition_translation",
		Action:  "delete",
		Payload: payload,
	}})
}

type AppScriptConditionTranslation struct {
	AppScriptConditionId string `json:"appScriptConditionId,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	AppScriptCondition *AppScriptCondition `json:"appScriptCondition,omitempty"`

	Language *Language `json:"language,omitempty"`

	Name string `json:"name,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type AppScriptConditionTranslationCollection struct {
	EntityCollection

	Data []AppScriptConditionTranslation `json:"data"`
}

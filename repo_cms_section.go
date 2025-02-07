package go_cicada_admin_sdk

import (
	"net/http"

	"time"
)

type CmsSectionRepository ClientService

func (t CmsSectionRepository) Search(ctx ApiContext, criteria Criteria) (*CmsSectionCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/cms-section", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(CmsSectionCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t CmsSectionRepository) SearchAll(ctx ApiContext, criteria Criteria) (*CmsSectionCollection, *http.Response, error) {
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

func (t CmsSectionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/cms-section", criteria)

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

func (t CmsSectionRepository) Upsert(ctx ApiContext, entity []CmsSection) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"cms_section": {
		Entity:  "cms_section",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t CmsSectionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"cms_section": {
		Entity:  "cms_section",
		Action:  "delete",
		Payload: payload,
	}})
}

type CmsSection struct {
	Visibility interface{} `json:"visibility,omitempty"`

	BackgroundMedia *Media `json:"backgroundMedia,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	MobileBehavior string `json:"mobileBehavior,omitempty"`

	BackgroundMediaId string `json:"backgroundMediaId,omitempty"`

	BackgroundMediaMode string `json:"backgroundMediaMode,omitempty"`

	SizingMode string `json:"sizingMode,omitempty"`

	Blocks []CmsBlock `json:"blocks,omitempty"`

	CmsPageVersionId string `json:"cmsPageVersionId,omitempty"`

	Id string `json:"id,omitempty"`

	Position float64 `json:"position,omitempty"`

	Type string `json:"type,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Locked bool `json:"locked,omitempty"`

	PageId string `json:"pageId,omitempty"`

	Page *CmsPage `json:"page,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Name string `json:"name,omitempty"`

	BackgroundColor string `json:"backgroundColor,omitempty"`

	CssClass string `json:"cssClass,omitempty"`
}

type CmsSectionCollection struct {
	EntityCollection

	Data []CmsSection `json:"data"`
}

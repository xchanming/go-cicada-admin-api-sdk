package go_cicada_admin_sdk

import (
	"net/http"

	"time"
)

type ScheduledTaskRepository ClientService

func (t ScheduledTaskRepository) Search(ctx ApiContext, criteria Criteria) (*ScheduledTaskCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/scheduled-task", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ScheduledTaskCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ScheduledTaskRepository) SearchAll(ctx ApiContext, criteria Criteria) (*ScheduledTaskCollection, *http.Response, error) {
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

func (t ScheduledTaskRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/scheduled-task", criteria)

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

func (t ScheduledTaskRepository) Upsert(ctx ApiContext, entity []ScheduledTask) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"scheduled_task": {
		Entity:  "scheduled_task",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ScheduledTaskRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"scheduled_task": {
		Entity:  "scheduled_task",
		Action:  "delete",
		Payload: payload,
	}})
}

type ScheduledTask struct {
	LastExecutionTime time.Time `json:"lastExecutionTime,omitempty"`

	NextExecutionTime time.Time `json:"nextExecutionTime,omitempty"`

	DefaultRunInterval float64 `json:"defaultRunInterval,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	RunInterval float64 `json:"runInterval,omitempty"`

	Status string `json:"status,omitempty"`

	ScheduledTaskClass string `json:"scheduledTaskClass,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`
}

type ScheduledTaskCollection struct {
	EntityCollection

	Data []ScheduledTask `json:"data"`
}

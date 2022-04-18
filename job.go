package stuartclient

import (
	"context"
	"fmt"
)

func (c ClientWrapper) CreateJob(ctx context.Context, model JobRequestModel) (*JobResponseModel, error) {
	builder := c.newRequest("/v2/jobs").
		BodyJSON(model)

	resp := new(JobResponseModel)
	if err := builder.ToJSON(resp).Fetch(ctx); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c ClientWrapper) GetJobPricing(ctx context.Context, model JobRequestModel) (*PricingProposalModel, error) {
	builder := c.newRequest("/v2/jobs/pricing").
		BodyJSON(model)

	pricing := new(PricingProposalModel)
	if err := builder.ToJSON(pricing).Fetch(ctx); err != nil {
		return nil, err
	}

	return pricing, nil
}

func (c ClientWrapper) ValidateJobParameters(ctx context.Context, model JobRequestModel) (bool, error) {
	builder := c.newRequest("/v2/jobs/validate").
		BodyJSON(model)

	if err := builder.Fetch(ctx); err != nil {
		return false, err
	}

	return true, nil
}

func (c ClientWrapper) GetJobOriginEta(ctx context.Context, model JobRequestModel) (int, error) {
	type etaResponse struct {
		Eta int `json:"eta"`
	}

	builder := c.newRequest("/v2/jobs/eta").
		BodyJSON(model)

	var resp etaResponse
	if err := builder.ToJSON(&resp).Fetch(ctx); err != nil {
		return 0, err
	}

	return resp.Eta, nil
}

func (c ClientWrapper) GetJob(ctx context.Context, id string) (*JobResponseModel, error) {
	builder := c.newRequest("/v2/jobs/" + id)
	resp := new(JobResponseModel)
	if err := builder.ToJSON(resp).Fetch(ctx); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c ClientWrapper) GetSchedulingSlots(ctx context.Context, zone string, requestType RequestType, date string) (*SchedulingSlotsResponseModel, error) {
	builder := c.newRequest(fmt.Sprintf("/v2/jobs/schedules/%s/%s/%s", zone, requestType, date))
	resp := new(SchedulingSlotsResponseModel)
	if err := builder.ToJSON(resp).Fetch(ctx); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c ClientWrapper) UpdateJob(ctx context.Context, jobId string, model JobRequestModel) (*JobResponseModel, error) {
	builder := c.newRequest("/v2/jobs/" + jobId).
		Method("PATCH").
		BodyJSON(model)

	resp := new(JobResponseModel)
	if err := builder.ToJSON(resp).Fetch(ctx); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c ClientWrapper) CancelJob(ctx context.Context, jobId string, cancelReq CancelRequestModel) error {
	builder := c.newRequest(fmt.Sprintf("/v2/jobs/%s/cancel", jobId)).
		BodyJSON(cancelReq)

	if err := builder.Fetch(ctx); err != nil {
		return err
	}

	return nil
}

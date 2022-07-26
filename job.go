package stuartclient

import (
	"bytes"
	"context"
	"fmt"
	_ "io"
	_ "net/http"
	"strconv"
)

func (c ClientWrapper) CreateJob(ctx context.Context, model JobRequestModel) ([]byte, error) {
	builder := c.newRequest("/v2/jobs").
		BodyJSON(model)

	var resp bytes.Buffer
	if err := builder.ToBytesBuffer(&resp).Fetch(ctx); err != nil {
		return []byte(" "), err
	}
	return resp.Bytes(), nil
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

func (c ClientWrapper) GetJobs(ctx context.Context, options GetJobsOptions) ([]JobResponseModel, error) {

	builder := c.newRequest("/v2/jobs/").
		Param("active", strconv.FormatBool(options.Active))

	if options.Status != "" {
		builder.Param("status", string(options.Status))
	}
	if options.Page > 0 {
		builder.ParamInt("page", options.Page)
	}
	if options.PerPage > 0 {
		builder.ParamInt("per_page", options.PerPage)
	}
	if options.ClientReference != "" {
		builder.Param("client_reference", string(options.ClientReference))
	}
	if options.Order != "" {
		builder.Param("order", string(options.Order))
	}

	var resp []JobResponseModel
	builder.ToJSON(resp)
	if err := builder.ToJSON(&resp).Fetch(ctx); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c ClientWrapper) GetJob(ctx context.Context, jobId string) (*JobResponseModel, error) {
	builder := c.newRequest("/v2/jobs/" + jobId)
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
		Patch().
		BodyJSON(model)

	resp := new(JobResponseModel)
	if err := builder.ToJSON(resp).Fetch(ctx); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c ClientWrapper) CancelJob(ctx context.Context, jobId string, cancelReq *CancelRequestModel) (int, error) {
	// var status int
	var resp string
	builder := c.newRequest(fmt.Sprintf("/v2/jobs/%s/cancel", jobId)).
		BodyBytes([]byte(``))
	if cancelReq != nil {
		builder.Param("public_reason_key", string(cancelReq.PublicReasonKey)).
			Param("comment", cancelReq.Comment)
	}
	if err := builder.ToString(&resp).Fetch(ctx); err != nil {
		fmt.Println(err.Error())
		return 0, err
	} else {
		fmt.Println(resp)
		return 201, nil
	}
}

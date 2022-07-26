package stuartclient

import (
	_ "bytes"
	"context"
	"github.com/paulmach/go.geojson"
)

type ClientInterface interface {
	GetDriverPhoneNumber(ctx context.Context, deliveryId string) (string, error)
	CancelDelivery(ctx context.Context, deliveryId string, cancelReq *CancelRequestModel) (int, error)
	ValidateAddress(ctx context.Context, address string, addressType RequestType, phone string) (bool, error)
	GetZoneCoverage(ctx context.Context, zone string, addressType RequestType) (*geojson.FeatureCollection, error)
	GetParcelShops(ctx context.Context, address string, date string) (*ParcelShopsResponse, error)
	CreateJob(ctx context.Context, model JobRequestModel) ([]byte, error)
	GetJobPricing(ctx context.Context, model JobRequestModel) (*PricingProposalModel, error)
	ValidateJobParameters(ctx context.Context, model JobRequestModel) (bool, error)
	GetJobOriginEta(ctx context.Context, model JobRequestModel) (int, error)
	GetJobs(ctx context.Context, options GetJobsOptions) ([]JobResponseModel, error)
	GetJob(ctx context.Context, jobId string) (*JobResponseModel, error)
	GetSchedulingSlots(ctx context.Context, zone string, requestType RequestType, date string) (*SchedulingSlotsResponseModel, error)
	UpdateJob(ctx context.Context, jobId string, model JobRequestModel) (*JobResponseModel, error)
	CancelJob(ctx context.Context, jobId string, cancelReq *CancelRequestModel) (int, error)
}

package stuartclient

import "time"

type RequestType string

const (
	PickupType  RequestType = "picking"
	DropoffType RequestType = "delivering"
)

/**
Request Models
*/

type JobRequestModel struct {
	Job JobModel `json:"job"`
}

type JobModel struct {
	PickupAt        time.Time             `json:"pickup_at"`
	AssignmentCode  string                `json:"assignment_code"`
	ClientReference string                `json:"client_reference"`
	Pickups         []PickupRequestModel  `json:"pickups"`
	Dropoffs        []DropoffRequestModel `json:"dropoffs"`
}

type ContactDetails struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Company   string `json:"company"`
}

type PackageType string

const (
	XSmall PackageType = "xsmall"
	Small  PackageType = "small"
	Medium PackageType = "medium"
	Large  PackageType = "large"
	Xlarge PackageType = "xlarge"
)

type AccessCodeType string

const (
	ScanQRText     AccessCodeType = "scan_qr_text"
	ScanBarcode128 AccessCodeType = "scan_barcode_128"
	QRText         AccessCodeType = "qr_text"
	Barcode128     AccessCodeType = "barcode_128"
)

type PickupRequestModel struct {
	Address string         `json:"address"`
	Comment string         `json:"comment"`
	Contact ContactDetails `json:"contact"`
	//	AccessCodes []AccessCodesModel `json:"access_codes"`
}

type DropoffRequestModel struct {
	PackageType        PackageType    `json:"package_type"`
	PackageDescription string         `json:"package_description"`
	ClientReference    string         `json:"client_reference"`
	Address            string         `json:"address"`
	Comment            string         `json:"comment,omitempty"`
	Contact            ContactDetails `json:"contact"`
	//	AccessCodes                []AccessCodesModel `json:"access_codes"`
	EndCustomerTimeWindowStart time.Time `json:"end_customer_time_window_start,omitempty"`
	EndCustomerTimeWindowEnd   time.Time `json:"end_customer_time_window_end,omitempty"`
}

type AccessCodesModel struct {
	Code         string         `json:"code"`
	Type         AccessCodeType `json:"type"`
	Title        string         `json:"title"`
	Instructions string         `json:"instructions"`
}

type AddressType struct {
	Street   string      `json:"street"`
	Postcode string      `json:"postcode"`
	City     interface{} `json:"city"`
	Zone     string      `json:"zone"`
	Country  string      `json:"country"`
}

/**
Response Models
*/

type JobResponseModel struct {
	Id             int                       `json:"id"`
	CreatedAt      time.Time                 `json:"created_at"`
	Status         string                    `json:"status"`
	PackageType    string                    `json:"package_type"`
	TransportType  string                    `json:"transport_type"`
	AssignmentCode string                    `json:"assignment_code"`
	DropoffAt      string                    `json:"dropoff_at"`
	PickupAt       string                    `json:"pickup_at"`
	EndedAt        string                    `json:"ended_at"`
	Comment        string                    `json:"comment"`
	Distance       float64                   `json:"distance"`
	Duration       int                       `json:"duration"`
	Deliveries     []DeliveriesResponseModel `json:"deliveries"`
	Driver         DriverModel               `json:"driver"`
	Pricing        PricingModel              `json:"pricing"`
	Cancellation   CancellationResponseModel `json:"cancelation"`
	Eta            EtaModel                  `json:"eta"`
	Rating         string                    `json:"rating"`
}

type DriverModel struct {
	Id            int     `json:"id"`
	DisplayName   string  `json:"display_name"`
	Phone         string  `json:"phone"`
	PictureUrl    string  `json:"picture_url"`
	TransportType string  `json:"transport_type"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
}

type PickupDropoffResponseModel struct {
	Id        int            `json:"id"`
	Latitude  float64        `json:"latitude"`
	Longitude float64        `json:"longitude"`
	Comment   string         `json:"comment"`
	Address   AddressType    `json:"address"`
	Contact   ContactDetails `json:"contact"`
}

type CancellationResponseModel struct {
	CanceledBy string `json:"canceled_by"`
	ReasonKey  string `json:"reason_key"`
	Comment    string `json:"comment"`
}

type EtaModel struct {
	Pickup  string `json:"pickup"`
	Dropoff string `json:"dropoff"`
}

type PricingModel struct {
	Currency         string  `json:"currency"`
	TaxPercentage    float64 `json:"tax_percentage"`
	PriceTaxIncluded float64 `json:"price_tax_included"`
	PriceTaxExcluded float64 `json:"price_tax_excluded"`
	TaxAmount        float64 `json:"tax_amount"`
	InvoiceUrl       string  `json:"invoice_url"`
}

type PricingProposalModel struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type ProofModel struct {
	SignatureUrl string `json:"signature_url"`
}

type DeliveriesResponseModel struct {
	Id                 int                        `json:"id"`
	Status             string                     `json:"status"`
	PickedAt           string                     `json:"picked_at"`
	DeliveredAt        string                     `json:"delivered_at"`
	TrackingUrl        string                     `json:"tracking_url"`
	ClientReference    string                     `json:"client_reference"`
	PackageDescription string                     `json:"package_description"`
	PackageType        string                     `json:"package_type"`
	Pickup             PickupDropoffResponseModel `json:"pickup"`
	Dropoff            PickupDropoffResponseModel `json:"dropoff"`
	Cancellation       CancellationResponseModel  `json:"cancellation"`
	Eta                EtaModel                   `json:"eta"`
	Proof              ProofModel                 `json:"proof"`
}

type TimeSlotModel struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type SchedulingSlotsResponseModel struct {
	Date  time.Time       `json:"date"`
	Type  string          `json:"type"`
	Slots []TimeSlotModel `json:"slots"`
}

type CancelReasonType string

const (
	CancelAddressError                  CancelReasonType = "address_error"
	CancelCourierIssue                  CancelReasonType = "courier_issue"
	CancelCustomerCancellationRequested CancelReasonType = "customer_cancellation_requested"
	CancelDuplicateJob                  CancelReasonType = "duplicate_job"
	CancelIncorrectPackage              CancelReasonType = "incrorrect_package"
	CancelNoSupply                      CancelReasonType = "no_supply"
	CancelOther                         CancelReasonType = "other"
	CancelPackageDamaged                CancelReasonType = "package_damaged"
	CancelPackageNotReady               CancelReasonType = "package_not_ready"
	CancelPUClosed                      CancelReasonType = "pu_closed"
	CancelTechnicalIssue                CancelReasonType = "technical_issue"
	CancelWrongTransportType            CancelReasonType = "wrong_transport_type"
)

type CancelRequestModel struct {
	PublicReasonKey CancelReasonType `json:"public_reason_key"`
	Comment         string           `json:"comment"`
}

type ParcelShopsResponse struct {
	Date     string         `json:"date"`
	Schedule []ScheduleType `json:"schedule"`
}

type ParcelContactType struct {
	Company string `json:"company"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type ParcelShopType struct {
	Address string            `json:"address"`
	Contact ParcelContactType `json:"contact"`
}

type ScheduleType struct {
	ParcelShop ParcelShopType `json:"parcel_shop"`
	From       time.Time      `json:"from"`
	To         time.Time      `json:"to"`
}

/**
Get Job Options
*/

type OrderValuesType string

const (
	ORDER_START_INVITING_AT_DESC OrderValuesType = "start_inviting_at:desc"
	ORDER_PICKUP_AT_DESC         OrderValuesType = "pickup_at:desc"
	ORDER_PICKUP_AT_ASC          OrderValuesType = "pickup_at:asc"
)

type StatusType string

const (
	STATUS_NEW         StatusType = "new"
	STATUS_SCHEDULED   StatusType = "scheduled"
	STATUS_SEARCHING   StatusType = "searching"
	STATUS_IN_PROGRESS StatusType = "in_progress"
	STATUS_FINISHED    StatusType = "finished"
	STATUS_CANCELLED   StatusType = "cancelled"
	STATUS_EXPIRED     StatusType = "expired"
)

type GetJobsOptions struct {
	Status          StatusType
	Page            int
	PerPage         int
	ClientReference string
	Active          bool
	Order           OrderValuesType
}

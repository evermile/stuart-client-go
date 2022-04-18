package stuartclient

import (
	"context"
	"fmt"
)

func (c ClientWrapper) GetDriverPhoneNumber(ctx context.Context, deliveryId string) (string, error) {
	type phoneResponse struct {
		PhoneNumber string `json:"phone_number"`
	}
	builder := c.newRequest(fmt.Sprintf("/v2/deliveries/%s/phone_number", deliveryId))
	var resp phoneResponse
	if err := builder.ToJSON(resp).Fetch(ctx); err != nil {
		return "", err
	}

	return resp.PhoneNumber, nil
}

func (c ClientWrapper) CancelDelivery(ctx context.Context, deliveryId string, cancelReq CancelRequestModel) error {
	builder := c.newRequest(fmt.Sprintf("/v2/deliveries/%s/cancel", deliveryId)).
		BodyJSON(cancelReq)

	if err := builder.Fetch(ctx); err != nil {
		return err
	}

	return nil
}

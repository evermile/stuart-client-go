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

func (c ClientWrapper) CancelDelivery(ctx context.Context, deliveryId string, cancelReq *CancelRequestModel) (int, error) {
	// var status int
	var resp string
	builder := c.newRequest(fmt.Sprintf("/v2/deliveries/%s/cancel", deliveryId)).
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

package stuartclient

import (
	"context"
	geojson "github.com/paulmach/go.geojson"
	url2 "net/url"
)

func (c ClientWrapper) ValidateAddress(ctx context.Context, address string, addressType RequestType, phone string) (bool, error) {
	builder := c.newRequest("/v2/addresses/validate").
		Param("address", url2.QueryEscape(address)).
		Param("type", string(addressType))
	if phone != "" {
		builder.Param("phone", phone)
	}
	if err := builder.Fetch(ctx); err != nil {
		return false, err
	}
	return true, nil
}

func (c ClientWrapper) GetZoneCoverage(ctx context.Context, zone string, addressType RequestType) (*geojson.FeatureCollection, error) {
	builder := c.newRequest("/v2/areas/" + zone)
	if addressType != "" {
		builder.Param("type", string(addressType))
	}
	resp := new(geojson.FeatureCollection)

	if err := builder.ToJSON(resp).Fetch(ctx); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c ClientWrapper) GetParcelShops(ctx context.Context, address string, date string) (*ParcelShopsResponse, error) {
	builder := c.newRequest("/v2/parcel_shops/around/schedule").
		Param("address", address).
		Param("date", date)

	resp := new(ParcelShopsResponse)
	if err := builder.ToJSON(resp).Fetch(ctx); err != nil {
		return nil, err
	}
	return resp, nil
}

package list

import (
	"context"
	"net/http"
	"net/url"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewClient(instance string, logger log.Logger, token string, opts ...httptransport.ClientOption) (Service, error) {
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}

	var getDEPAccountInfoEndpoint endpoint.Endpoint
	{
		getDEPAccountInfoEndpoint = httptransport.NewClient(
			"GET",
			copyURL(u, "/v1/dep/account"),
			encodeRequestWithToken(token, EncodeHTTPGenericRequest),
			DecodeDEPAccountInfoResponse,
			opts...,
		).Endpoint()
	}

	var getDEPDeviceDetailsEndpoint endpoint.Endpoint
	{
		getDEPDeviceDetailsEndpoint = httptransport.NewClient(
			"GET",
			copyURL(u, "/v1/dep/devices"),
			encodeRequestWithToken(token, EncodeHTTPGenericRequest),
			DecodeDEPDeviceDetailsReponse,
			opts...,
		).Endpoint()
	}

	var getDEPProfilesEndpoint endpoint.Endpoint
	{
		getDEPProfilesEndpoint = httptransport.NewClient(
			"GET",
			copyURL(u, "/v1/dep/profiles"),
			encodeRequestWithToken(token, EncodeHTTPGenericRequest),
			DecodeDEPProfileResponse,
			opts...,
		).Endpoint()
	}

	return Endpoints{
		GetDEPAccountInfoEndpoint: getDEPAccountInfoEndpoint,
		GetDEPDeviceEndpoint:      getDEPDeviceDetailsEndpoint,
		GetDEPProfileEndpoint:     getDEPProfilesEndpoint,
	}, nil
}

func encodeRequestWithToken(token string, next httptransport.EncodeRequestFunc) httptransport.EncodeRequestFunc {
	return func(ctx context.Context, r *http.Request, request interface{}) error {
		r.SetBasicAuth("micromdm", token)
		return next(ctx, r, request)
	}
}
func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}

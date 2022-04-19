package stuartclient

import (
	"context"
	"fmt"
	"github.com/carlmjohnson/requests"
	creds "golang.org/x/oauth2/clientcredentials"
	"net/http"
)

type ClientWrapper struct {
	http.Client
	envEndpoint EnvEndpointType
}

func StuartResponseHandler(resp *http.Response) error {
	type errorResponse struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	}

	statusErr := requests.CheckStatus(
		http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
		http.StatusNonAuthoritativeInfo,
		http.StatusNoContent,
	)(resp)
	if statusErr == nil {
		return nil
	}
	var placeholder errorResponse
	statusErr = requests.ToJSON(&placeholder)(resp)
	return fmt.Errorf("%s: %s", placeholder.Error, placeholder.Message)
}

func (c ClientWrapper) newRequest(suffix string) *requests.Builder {
	return requests.
		URL(c.envEndpoint.GetURL(suffix)).
		Client(&c.Client).
		AddValidator(StuartResponseHandler)
}

func NewClient(ctx context.Context, envType EnvEndpointType, apiClientId string, apiClientSecret string) (client ClientInterface) {
	config := new(creds.Config)
	config.ClientID = apiClientId
	config.ClientSecret = apiClientSecret
	config.TokenURL = envType.GetOAuthURL()

	httpClient := config.Client(ctx)
	return ClientWrapper{
		Client:      *httpClient,
		envEndpoint: envType,
	}
}

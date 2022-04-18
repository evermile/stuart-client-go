package stuartclient

type EnvEndpointType string

const (
	SandboxEnv     EnvEndpointType = "https://api.sandbox.stuart.com"
	ProdEnv        EnvEndpointType = "https://api.stuart.com"
	TokenURLSuffix                 = "/oauth/token"
)

func (endpoint EnvEndpointType) GetBaseURL() string {
	return string(endpoint)
}

func (endpoint EnvEndpointType) GetURL(suffix string) string {
	return endpoint.GetBaseURL() + suffix
}

func (endpoint EnvEndpointType) GetOAuthURL() string {
	return endpoint.GetURL(TokenURLSuffix)
}

//go:build integration
// +build integration

package pfsense

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestPfsenseClient(t *testing.T) {
	// custom HTTP client
	hc := http.Client{
		Transport: &http.Transport{
			// InsecureSkipVerify is set to true to allow self-signed certificates
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	err := godotenv.Load(".env")
	assert.NoError(t, err, "failed to load .env file")
	apiKey := os.Getenv(EnvPfsenseAPIKey)
	apiUri := os.Getenv(EnvPfsenseURI)

	apiKeyProvider := NewApiKeyAuthProvider(apiKey)

	c, err := NewClientWithResponses(apiUri, WithHTTPClient(&hc), WithRequestEditorFn(apiKeyProvider.Intercept))
	assert.NoError(t, err, "failed to create client")
	var pageSize = 10
	var offset = 0

	resp, err := c.GetServicesDNSResolverHostOverridesEndpointWithResponse(context.Background(), &GetServicesDNSResolverHostOverridesEndpointParams{
		Limit:  &pageSize,
		Offset: &offset,
	})

	assert.NoError(t, err, "failed to get services DNS resolver host overrides")
	// assert that the response is not nil
	assert.NotNil(t, resp, "response is nil")
	assert.Equal(t, http.StatusOK, resp.StatusCode(), "response status code is not 200")
}

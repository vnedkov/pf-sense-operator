package pfsense

import (
	"context"
	"net/http"
)

type ApiKeyAuthProvider struct {
	token string
}

func NewApiKeyAuthProvider(token string) *ApiKeyAuthProvider {
	return &ApiKeyAuthProvider{
		token: token,
	}
}

func (s *ApiKeyAuthProvider) Intercept(ctx context.Context, req *http.Request) error {
	req.Header.Set("X-API-Key", s.token)
	return nil
}

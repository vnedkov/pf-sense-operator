package pfsense

import (
	"context"
	"net/http"
)

type BearerTokenAuthProvider struct {
	token string
}

func NewBearerTokenAuthProvider(token string) *BearerTokenAuthProvider {
	return &BearerTokenAuthProvider{
		token: token,
	}
}

func (s *BearerTokenAuthProvider) Intercept(ctx context.Context, req *http.Request) error {
	req.Header.Set("Authorization", "Bearer "+s.token)
	return nil
}

package pfsense

import (
	"context"
	"net/http"
)

type BasicAuthProvider struct {
	userName string
	password string
}

func NewBasicAuthProvider(userName, password string) *BasicAuthProvider {
	return &BasicAuthProvider{
		userName: userName,
		password: password,
	}
}

func (s *BasicAuthProvider) Intercept(ctx context.Context, req *http.Request) error {
	req.SetBasicAuth(s.userName, s.password)
	return nil
}

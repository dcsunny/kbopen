package http

import "encoding/json"

type CommonResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CommonResultWithData struct {
	CommonResult
	Data json.RawMessage `json:"data"`
}

type AuthorizerUser struct {
	AuthorizerUserId string `json:"authorizerUserId"`
}

func (a *AuthorizerUser) SetAuthorizerUserId(authorizerUserId string) {
	if a == nil {
		a = &AuthorizerUser{}
	}
	a.AuthorizerUserId = authorizerUserId
}

func (a *AuthorizerUser) GetAuthorizerUserId() string {
	return a.AuthorizerUserId
}

type RequestAuthorizer interface {
	SetAuthorizerUserId(authorizerUserId string)
	GetAuthorizerUserId() string
}

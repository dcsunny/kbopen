package auth

import (
	"fmt"

	"github.com/dcsunny/kbopen/conf"
	"github.com/dcsunny/kbopen/context"
)

const (
	getAuthorizerUrl = "/api/get_authorizer_user_id"
)

type Authorizer struct {
	ctx *context.Context
}

func NewAuthorizer(cfg *conf.Config) *Authorizer {
	ctx := context.NewContext(cfg)
	return &Authorizer{
		ctx: ctx,
	}
}

type GetAuthorizerUserIDReq struct {
	Code string
}

type GetAuthorizerUserIDData struct {
	AuthorizerUserID string `json:"authorizerUserID"`
}

func (a *Authorizer) GetAuthorizerUserID(code string) (string, error) {
	req := GetAuthorizerUserIDReq{Code: code}
	var result GetAuthorizerUserIDData
	_, err := a.ctx.HttpClient.HttpPostJson(getAuthorizerUrl, req, &result, a.ctx.GetAccessToken)
	if err != nil {
		return "", err
	}
	return result.AuthorizerUserID, nil
}

func GetHtsyAuthLink(appid string, callback string) string {
	return fmt.Sprintf("https://htsy.cpshelp.cn/auth/third/part?appid=%s&redirect_uri=%s", appid, callback)
}

func GetCpsAuthLink(appid string, callback string) string {
	return fmt.Sprintf("https://mp.cpshelp.cn/auth/third/part?appid=%s&redirect_uri=%s", appid, callback)
}

func GetKqAuthLink(appid string, callback string) string {
	return fmt.Sprintf("https://kq.cpshelp.cn/auth/third/part?appid=%s&redirect_uri=%s", appid, callback)
}

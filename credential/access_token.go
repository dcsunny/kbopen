package credential

import (
	"fmt"
	"time"

	"github.com/dcsunny/kbopen/cache"
	"github.com/dcsunny/kbopen/http"
)

type AccessTokenHandle interface {
	GetAccessToken() (accessToken string, err error)
}

const (
	getAccessTokenUrl = "/api/get_access_token"
)

type DefaultAccessToken struct {
	appID      string
	appSecret  string
	cache      cache.Cache
	httpClient *http.HttpClient
}

func NewDefaultAccessToken(appID string,
	appSecret string,
	_cache cache.Cache, httpClient *http.HttpClient) *DefaultAccessToken {
	return &DefaultAccessToken{
		appID:      appID,
		appSecret:  appSecret,
		cache:      _cache,
		httpClient: httpClient,
	}
}

type GetServerTokenReq struct {
	Appid     string `json:"appid"`
	AppSecret string `json:"appSecret"`
}

type GetServerTokenResp struct {
	AccessToken string `json:"accessToken"`
	ExpireTime  int64  `json:"expireTime"`
}

func (a *DefaultAccessToken) GetRemoteServerToken() (*GetServerTokenResp, error) {
	req := GetServerTokenReq{
		Appid:     a.appID,
		AppSecret: a.appSecret,
	}
	var result GetServerTokenResp
	_, err := a.httpClient.HttpPostJson(getAccessTokenUrl, req, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *DefaultAccessToken) GetAccessToken() (string, error) {

	accessTokenCacheKey := fmt.Sprintf("kbopen_access_token_%s", a.appID)
	val := a.cache.Get(accessTokenCacheKey)
	if val != nil {
		return val.(string), nil
	}

	//cache失效，从微信服务器获取
	r, err := a.GetRemoteServerToken()
	if err != nil {
		return "", err
	}

	//留10分钟
	expires := r.ExpireTime/1000 - time.Now().Unix() - 600

	err = a.cache.Set(accessTokenCacheKey, r.AccessToken, time.Duration(expires)*time.Second)
	if err != nil {
		return "", nil
	}
	return r.AccessToken, nil
}

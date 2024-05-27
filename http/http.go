package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/dcsunny/kbopen/conf"
)

type HttpClient struct {
	cfg            *conf.Config
	HttpLastResult HttpLastResult
}

type HttpLastResult struct {
	Body []byte
	Err  error
}

func NewHttpClient(cfg *conf.Config) *HttpClient {
	return &HttpClient{
		cfg: cfg,
	}
}

func (c *HttpClient) HttpPostJson(link string, req interface{}, resp interface{}, tokenFunc func() (string, error)) (CommonResult, interface{}, error) {
	var err error
	var result CommonResultWithData
	link, err = url.JoinPath(c.cfg.Endpoint, link)
	if err != nil {
		return result.CommonResult, resp, err
	}

	client := http.Client{
		Timeout: time.Second * time.Duration(c.cfg.HttpTimeout),
	}
	var reqReader io.Reader
	if req != nil {
		j, _ := json.Marshal(req)
		reqReader = bytes.NewReader(j)
	}

	httpReq, err := http.NewRequest(http.MethodPost, link, reqReader)
	if err != nil {
		return result.CommonResult, resp, err
	}
	httpReq.Header.Set("content-type", "application/json")
	if tokenFunc != nil {
		var token string
		token, err = tokenFunc()
		if err != nil {
			return CommonResult{}, resp, err
		}
		httpReq.Header.Set("Authorization", "Bearer "+token)
	}
	var r *http.Response
	r, err = client.Do(httpReq)
	defer func() {
		c.HttpLastResult.Err = err
		if r.Body != nil {
			r.Body.Close()
		}
	}()
	if err != nil {
		return result.CommonResult, resp, err
	}
	var body []byte
	body, err = io.ReadAll(r.Body)
	if err != nil {
		return result.CommonResult, resp, err
	}
	c.HttpLastResult.Body = body
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result.CommonResult, resp, err
	}
	if result.CommonResult.Code != 0 {
		return result.CommonResult, resp, errors.New(result.CommonResult.Message)
	}
	if resp == nil {
		return result.CommonResult, resp, nil
	}
	err = json.Unmarshal(result.Data, resp)
	if err != nil {
		return result.CommonResult, resp, err
	}
	return result.CommonResult, resp, nil
}

func (c *HttpClient) HttpPostJsonWithAuthorizer(link string, req RequestAuthorizer, resp interface{}, tokenFunc func() (string, error)) (CommonResult, interface{}, error) {
	if req == nil {
		req = &AuthorizerUser{}
	}
	if reflect.ValueOf(req).IsNil() {
		req = &AuthorizerUser{}
	}
	if req.GetAuthorizerUserId() == "" {
		req.SetAuthorizerUserId(c.cfg.AuthorizerUserId)
	}
	if req.GetAuthorizerUserId() == "" {
		return CommonResult{}, resp, errors.New("authorizerUserId is empty")
	}
	return c.HttpPostJson(link, req, resp, tokenFunc)
}

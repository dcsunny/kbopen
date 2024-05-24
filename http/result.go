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

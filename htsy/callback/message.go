package callback

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/dcsunny/kbopen/context"
)

type Callback struct {
	ctx *context.Context

	Writer http.ResponseWriter

	UrlQuery url.Values

	RequestBody []byte

	messageHandler MessageHandler
}

type MessageHandler func(message *Message)

func NewCallback(ctx *context.Context) *Callback {
	return &Callback{
		ctx: ctx,
	}
}

func (c *Callback) Validate(msg Message) (string, error) {

	var result ValidateResp
	result.Handshake = msg.Handshake
	j, _ := json.Marshal(result)
	if c.Writer != nil {
		c.Writer.Header().Set("content-type", "application/json")
		c.Writer.WriteHeader(200)
		c.Writer.Write(j)
	}
	return string(j), nil
}

func (c *Callback) Serve() (string, error) {
	if len(c.RequestBody) == 0 {
		return "", nil
	}
	var msg Message
	err := json.Unmarshal(c.RequestBody, &msg)
	if err != nil {
		return "", nil
	}
	if msg.Type == TypeHandshake || msg.Handshake != "" {
		//校验
		_, err = c.Validate(msg)
		if err != nil {
			return "", err
		}

		if msg.Type == TypeHandshake {
			return "", nil
		}
	}

	if c.messageHandler != nil {
		c.messageHandler(&msg)
	}

	return "", nil
}

func (c *Callback) SetHandler(h MessageHandler) {
	c.messageHandler = h
}

type ValidateReq struct {
	Type      int    `json:"type"`
	Handshake string `json:"handshake"`
}

type ValidateResp struct {
	Handshake string `json:"handshake"`
}

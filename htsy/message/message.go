package message

import (
	"github.com/dcsunny/kbopen/context"
	"github.com/dcsunny/kbopen/http"
)

const (
	sendUrl = "/api/htsy/msg/send"
)

type Msg struct {
	ctx *context.Context
}

func NewMsg(ctx *context.Context) *Msg {
	return &Msg{
		ctx: ctx,
	}
}

type SendReq struct {
	http.AuthorizerUser
	Wxid    string  `json:"wxid"`    //接收者
	BotWxid string  `json:"botWxid"` //发送的微号
	BotType int     `json:"botType"` //微号类型
	Msg     Message `json:"msg"`
}

type SendData struct {
	TraceId string `json:"traceId"`
}

func (m *Msg) Send(req *SendReq) (string, error) {
	_, result, err := m.ctx.HttpClient.HttpPostJsonWithAuthorizer(sendUrl, req, &SendData{}, m.ctx.GetAccessToken)
	if err != nil {
		return "", err
	}
	return result.(*SendData).TraceId, nil
}

func BuildMessageText(content string) Message {
	return Message{
		MsgType: MsgTypeText,
		Text: Text{
			Content: content,
		},
	}
}

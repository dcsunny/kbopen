package room

import (
	"github.com/dcsunny/kbopen/context"
	"github.com/dcsunny/kbopen/http"
)

const (
	listUrl = "/api/htsy/group/list"
)

type Room struct {
	ctx *context.Context
}

func NewRoom(ctx *context.Context) *Room {
	return &Room{
		ctx: ctx,
	}
}

type ListReq struct {
	http.AuthorizerUser
	BotType   int    `json:"botType"`
	Page      int    `json:"page,omitempty"`
	PageSize  int    `json:"pageSize,omitempty"`
	GroupId   int    `json:"groupId,omitempty"`
	GroupName string `json:"groupName,omitempty"`
	GroupType string `json:"groupType,omitempty"`
	BotId     int    `json:"botId,omitempty"`
	CorpId    string `json:"corpId,omitempty"`
	IsManager bool   `json:"isManager,omitempty"`
}

type ListData struct {
	List     []ListItem `json:"list"`
	Page     int        `json:"page"`
	PageSize int        `json:"pageSize"`
	Total    int        `json:"total"`
	Last     bool       `json:"last"`
}

type ListItem struct {
	Id          int    `json:"id"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	TotalMember int    `json:"totalMember"`
	Manager     struct {
		MemberId int    `json:"memberId"`
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
		BotId    int    `json:"botId"`
	} `json:"manager"`
	Wxid      string `json:"wxid"`
	GroupType string `json:"groupType"`
	Status    int    `json:"status"`
}

func (r *Room) List(req *ListReq) (*ListData, error) {
	_, result, err := r.ctx.HttpClient.HttpPostJsonWithAuthorizer(listUrl, req, &ListData{}, r.ctx.GetAccessToken)
	if err != nil {
		return nil, err
	}
	return result.(*ListData), nil
}

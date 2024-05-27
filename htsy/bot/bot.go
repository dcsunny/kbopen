package bot

import (
	"github.com/dcsunny/kbopen/context"
	"github.com/dcsunny/kbopen/http"
)

const (
	infoUrl = "/api/htys/bot/info"
	listUrl = "/api/htsy/bot/list"
)

type Bot struct {
	ctx *context.Context
}

func NewBot(ctx *context.Context) *Bot {
	return &Bot{
		ctx: ctx,
	}
}

type InfoReq struct {
	http.AuthorizerUser
	Wxid    string `json:"wxid"`
	BotType int    `json:"botType"`
}

type InfoData struct {
	Id         int    `json:"id"`
	Status     int    `json:"status"`
	Version    string `json:"version"`
	ClientType string `json:"clientType"`
	BotKey     string `json:"botKey"`
	Wxid       string `json:"wxid"`
	Ip         string `json:"ip"`
	Port       int    `json:"port"`
	CodeType   int    `json:"codeType"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
}

func (b *Bot) Info(req *InfoReq) (*InfoData, error) {
	_, result, err := b.ctx.HttpClient.HttpPostJsonWithAuthorizer(infoUrl, req, &InfoData{}, b.ctx.GetAccessToken)
	if err != nil {
		return nil, err
	}
	return result.(*InfoData), nil
}

type ListReq struct {
	http.AuthorizerUser
	Page       int     `json:"page"`
	PageSize   int     `json:"pageSize"`
	Ids        []int64 `json:"ids"`
	Name       string  `json:"name"`
	StatusList []int   `json:"statusList"`
	GroupId    int     `json:"groupId"`
	BotType    int     `json:"botType"`
	CorpId     string  `json:"corpId"`
}

type ListData struct {
	List     []ListItem `json:"list"`
	Page     int        `json:"page"`
	PageSize int        `json:"pageSize"`
	Total    int        `json:"total"`
	Last     bool       `json:"last"`
}

type ListItem struct {
	Id             int64  `json:"id"`
	Status         int    `json:"status"`
	Version        string `json:"version"`
	ClientType     string `json:"clientType"`
	BotKey         string `json:"botKey"`
	ExpireTime     int    `json:"expireTime"`
	Ip             string `json:"ip"`
	Port           int    `json:"port"`
	Nickname       string `json:"nickname"`
	Avatar         string `json:"avatar"`
	Wxid           string `json:"wxid"`
	WxAccount      string `json:"wxAccount"`
	Remark         string `json:"remark"`
	CorpName       string `json:"corpName"`
	CodeType       int    `json:"codeType"`
	OfflineTime    int    `json:"offlineTime"`
	Name           string `json:"name"`
	BotStatus      int    `json:"botStatus"`
	CorpShortName  string `json:"corpShortName"`
	CorpId         string `json:"corpId"`
	RelateFriendId int    `json:"relateFriendId"`
	WxVersion      string `json:"wxVersion"`
	LoginPlatform  int    `json:"loginPlatform"`
}

func (b *Bot) List(req *ListReq) (*ListData, error) {
	_, result, err := b.ctx.HttpClient.HttpPostJsonWithAuthorizer(listUrl, req, &ListData{}, b.ctx.GetAccessToken)
	if err != nil {
		return nil, err
	}
	return result.(*ListData), nil
}

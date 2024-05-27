package message

type Message struct {
	MsgType      string      `json:"msgType"`
	Text         Text        `json:"text,omitempty"`
	Image        Image       `json:"image,omitempty"`
	Sticker      Sticker     `json:"sticker,omitempty"`
	Video        Video       `json:"video,omitempty"`
	MiniProgram  MiniProgram `json:"miniProgram,omitempty"`
	Link         Link        `json:"link,omitempty"`
	Audio        Audio       `json:"audio,omitempty"`
	FreeVideo    FreeVideo   `json:"freeVideo,omitempty"`
	Group        Group       `json:"group,omitempty"`
	InviteGroup  InviteGroup `json:"inviteGroup,omitempty"`
	MaterialId   int64       `json:"materialId,omitempty"`
	TraceId      string      `json:"traceId,omitempty"`
	MaterialName string      `json:"materialName,omitempty"`
}

type Text struct {
	Content string   `json:"content"`
	AtList  []string `json:"atList"`
}

type Image struct {
	Link string `json:"link"`
}

type Sticker struct {
	Link string `json:"link"`
}

type Video struct {
	Link string `json:"link"`
}

type MiniProgram struct {
	Appid    string `json:"appid"`
	PagePath string `json:"pagePath"`
	Title    string `json:"title"`
	ImgUrl   string `json:"imgUrl"`
	AppIcon  string `json:"appIcon"`
	AppName  string `json:"appName"`
	Username string `json:"username"`
}

type Link struct {
	Title  string `json:"title"`
	Link   string `json:"link"`
	ImgUrl string `json:"imgUrl"`
	Desc   string `json:"desc"`
}

type Audio struct {
	Link string `json:"link"`
}

type FreeVideo struct {
	ObjectId      string `json:"objectId"`
	ObjectNonceId string `json:"objectNonceId"`
	UserName      string `json:"userName"`
	NickName      string `json:"nickName"`
	Avatar        string `json:"avatar"`
	Desc          string `json:"desc"`
	ThumbUrl      string `json:"thumbUrl"`
	Url           string `json:"url"`
	FeedType      int    `json:"feedType"`
	CoverUrl      string `json:"coverUrl"`
	Extras        string `json:"extras"`
}

type Group struct {
	Id                     string    `json:"id"` //群id
	PromotionalMsgList     []Message `json:"promotionalMsgList"`
	PromotionalMsgInterval int       `json:"promotionalMsgInterval"`
}

type InviteGroup struct {
	Id int64 `json:"id"` //群邀请方案的id
}

package common

// QwWxidByFriend 需要发送消息的时候，需要转成R：开头的id,这个是好友的
func QwWxidByFriend(wxid string) string {
	return "S:" + wxid
}

// QwWxidByRoom 需要发送消息的时候，需要转成R：开头的id,这个是群的
func QwWxidByRoom(wxid string) string {
	return "R:" + wxid
}

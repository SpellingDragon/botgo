package dto

// GroupRichMediaUploadReq v2 群富媒体上传请求
// POST /v2/groups/{group_openid}/files
// file_type: 1图片 2视频 3语音 4文件; url 需平台可拉取的公网地址
type GroupRichMediaUploadReq struct {
	FileType   int    `json:"file_type"`
	URL        string `json:"url"`
	SRVSendMsg bool   `json:"srv_send_msg,omitempty"`
}

// GroupRichMediaUploadResp 上传响应，file_info 有效期有限需尽快引用发送
type GroupRichMediaUploadResp struct {
	FileUUID string `json:"file_uuid"`
	FileInfo string `json:"file_info"`
	TTL      int    `json:"ttl"`
}

// GroupMediaMessage v2 群富媒体消息（msg_type=7 时 Media 必填）
type GroupMediaMessage struct {
	MsgType int                `json:"msg_type"` // 7 = 富媒体
	Content string             `json:"content"`  // 文本说明，可为空串
	Media   *GroupMediaPayload `json:"media"`
	MsgID   string             `json:"msg_id,omitempty"` // 被动回复时传
	MsgSeq  int                `json:"msg_seq,omitempty"`
}

type GroupMediaPayload struct {
	FileInfo string `json:"file_info"`
}

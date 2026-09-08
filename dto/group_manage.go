package dto

// GroupRestrictChatReq v2 群成员禁言（限制发言）
// POST /v2/groups/{group_openid}/restrict_chat_setting
type GroupRestrictChatReq struct {
	ActionType string `json:"action_type"` // add / del
	ExpireAt   int64  `json:"mute_expire_at,omitempty"`
}

// GroupRestrictChatResp 禁言设置响应
type GroupRestrictChatResp struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    []GroupRestrictChatItem `json:"data,omitempty"`
}

type GroupRestrictChatItem struct {
	MemberOpenID string `json:"member_openid"`
	ExpireAt     int64  `json:"mute_expire_at"`
}

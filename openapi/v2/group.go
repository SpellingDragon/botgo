// Package v2 提供官方 SDK 尚未封装的 v2 群场景接口（富媒体/管理）。
// 实现策略：薄封装在 OpenAPI.Transport 透传通道之上，
// 复用 v1 客户端的 token 刷新、过滤器与日志链，零额外鉴权代码。
package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/openapi"
)

const (
	groupRichMediaUploadURI = "/v2/groups/%s/files"
	groupRichMediaMsgURI    = "/v2/groups/%s/messages"
	groupRestrictChatURI    = "/v2/groups/%s/members/%s/restrict_chat_setting"
)

// V2 v2 群场景能力集
type V2 struct {
	api openapi.OpenAPI // 复用 v1 透传通道（含 token 链）
}

// New 基于 v1 OpenAPI 实例构造 v2 能力集
func New(api openapi.OpenAPI) *V2 {
	return &V2{api: api}
}

// GroupRichMediaUpload 上传富媒体（file_type 1图/2视频/3语音/4文件）
// 返回 file_info 供引用发送，注意平台侧 TTL，取得后尽快发送。
func (v *V2) GroupRichMediaUpload(ctx context.Context, groupOpenID string,
	req *dto.GroupRichMediaUploadReq) (*dto.GroupRichMediaUploadResp, error) {
	if req == nil || req.URL == "" {
		return nil, fmt.Errorf("v2: rich media upload requires url")
	}
	body, err := v.api.Transport(ctx, http.MethodPost,
		fmt.Sprintf(groupRichMediaUploadURI, groupOpenID), req)
	if err != nil {
		return nil, err
	}
	out := &dto.GroupRichMediaUploadResp{}
	if err := json.Unmarshal(body, out); err != nil {
		return nil, fmt.Errorf("v2: decode upload resp: %w (%s)", err, string(body))
	}
	return out, nil
}

// GroupRichMediaMessage 发送富媒体消息（msg_type=7, media.file_info 引用上传结果）
// content 可为空串或附说明文字。
func (v *V2) GroupRichMediaMessage(ctx context.Context, groupOpenID string,
	msg *dto.GroupMediaMessage) error {
	if msg == nil || msg.Media == nil || msg.Media.FileInfo == "" {
		return fmt.Errorf("v2: media message requires file_info")
	}
	msg.MsgType = 7
	_, err := v.api.Transport(ctx, http.MethodPost,
		fmt.Sprintf(groupRichMediaMsgURI, groupOpenID), msg)
	return err
}

// GroupRestrictChat 群成员禁言/解除（action_type add/del, mute_expire_at 秒级时间戳）
// 官方契约: POST /v2/groups/{group_openid}/members/{member_openid}/restrict_chat_setting
func (v *V2) GroupRestrictChat(ctx context.Context, groupOpenID, memberOpenID string,
	req *dto.GroupRestrictChatReq) (*dto.GroupRestrictChatResp, error) {
	if memberOpenID == "" {
		return nil, fmt.Errorf("v2: restrict_chat requires member_openid")
	}
	if req == nil || (req.ActionType != "add" && req.ActionType != "del") {
		return nil, fmt.Errorf("v2: restrict_chat requires action_type add|del")
	}
	body, err := v.api.Transport(ctx, http.MethodPost,
		fmt.Sprintf(groupRestrictChatURI, groupOpenID, memberOpenID), req)
	if err != nil {
		return nil, err
	}
	out := &dto.GroupRestrictChatResp{}
	if err := json.Unmarshal(body, out); err != nil {
		return nil, fmt.Errorf("v2: decode restrict resp: %w (%s)", err, string(body))
	}
	return out, nil
}

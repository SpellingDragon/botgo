package v1

import (
	"context"
	"encoding/json"

	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/log"
	"github.com/tidwall/gjson"
)

// PostMessage 发消息
func (o *openAPI) PutThread(ctx context.Context, channelID string, msg *dto.ThreadToCreate) (*dto.ForumAuditResult,
	error) {
	resp, err := o.request(ctx).
		SetResult(dto.ForumAuditResult{}).
		SetPathParam("channel_id", channelID).
		SetBody(msg).
		Put(o.getURL(threadsURI))
	if err != nil {
		return nil, err
	}

	log.Infof("put thread response:%v", resp)
	return resp.Result().(*dto.ForumAuditResult), nil
}

// Message 拉取单条消息
func (o *openAPI) GetThread(ctx context.Context, channelID string, messageID string) (*dto.Message, error) {
	resp, err := o.request(ctx).
		SetResult(dto.Message{}).
		SetPathParam("channel_id", channelID).
		SetPathParam("message_id", messageID).
		Get(o.getURL(threadsURI))
	if err != nil {
		return nil, err
	}

	// 兼容处理
	result := resp.Result().(*dto.Message)
	if result.ID == "" {
		body := gjson.Get(resp.String(), "message")
		if err := json.Unmarshal([]byte(body.String()), result); err != nil {
			return nil, err
		}
	}
	return result, nil
}

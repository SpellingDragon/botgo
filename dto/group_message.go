package dto

import "time"

// GroupMessage 消息结构体定义
type GroupMessage struct {
	Author      Author    `json:"author"`
	Content     string    `json:"content"`
	GroupOpenID string    `json:"group_openid"`
	ID          string    `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	// MessageReference 引用消息元数据 (官方群帧字段): prompt 为被引消息内容摘要
	MessageReference struct {
		MessageID string `json:"message_id"`
		Prompt    string `json:"prompt"`
	} `json:"message_reference"`
}

// Author represents the author part of the data.
type Author struct {
	MemberOpenID string `json:"member_openid"`
	Username     string `json:"username"` // QQ 用户昵称 (群昵称需 V2 members API 反查)
}

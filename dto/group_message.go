package dto

import "time"

// GroupMessage 消息结构体定义
type GroupMessage struct {
	Author      Author    `json:"author"`
	Content     string    `json:"content"`
	GroupOpenID string    `json:"group_openid"`
	ID          string    `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
}

// Author represents the author part of the data.
type Author struct {
	MemberOpenID string `json:"member_openid"`
}

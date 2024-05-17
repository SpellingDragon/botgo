package dto

import "time"

// GroupMessage 消息结构体定义
type GroupMessage struct {
	Op int    `json:"op"`
	S  int    `json:"s"`
	T  string `json:"t"`
	ID string `json:"id"`
	D  Data   `json:"d"`
}

// Data represents the data part of the websocket message.
type Data struct {
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

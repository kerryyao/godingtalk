package message

import "github.com/kerryyao/godingtalk"

// MessageResponse is
type MessageResponse struct {
	godingtalk.OAPIResponse
	MessageID string `json:"messageId"`
}

// MessageResponse is
type MessageReadListResponse struct {
	godingtalk.OAPIResponse
	NextCursor     int64    `json:"next_cursor"`
	ReadUserIdList []string `json:"readUserIdList"`
}

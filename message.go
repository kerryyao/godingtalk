package godingtalk

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// SendAppMessage is 发送企业会话消息
func SendAppMessage(agentID string, touser string, msg string) (*OAPIResponse, error) {
	if agentID == "" {
		agentID = conf.AgentID
	}
	var data OAPIResponse
	request := map[string]interface{}{
		"touser":  touser,
		"agentid": agentID,
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": msg,
		},
	}
	payload, err := httpRequest("message/send", nil, request)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// SendAppOAMessage is 发送OA消息
func SendAppOAMessage(agentID string, touser string, msg OAMessage) (*OAPIResponse, error) {
	if agentID == "" {
		agentID = conf.AgentID
	}
	var data OAPIResponse
	request := map[string]interface{}{
		"touser":  touser,
		"agentid": agentID,
		"msgtype": "oa",
		"oa":      msg,
	}
	payload, err := httpRequest("message/send", nil, request)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// ActionCardMessage
func SendOverAllActionCardMessage(agentID string, touser string, msg OverAllActionCardMessage) (*OAPIResponse, error) {
	if agentID == "" {
		agentID = conf.AgentID
	}
	var data OAPIResponse
	request := map[string]interface{}{
		"touser":      touser,
		"agentid":     agentID,
		"msgtype":     "action_card",
		"action_card": msg,
	}
	payload, err := httpRequest("message/send", nil, request)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func SendIndependentActionCardMessage(agentID string, touser string, msg IndependentActionCardMessage) (*OAPIResponse, error) {
	if agentID == "" {
		agentID = conf.AgentID
	}
	var data OAPIResponse
	request := map[string]interface{}{
		"touser":      touser,
		"agentid":     agentID,
		"msgtype":     "action_card",
		"action_card": msg,
	}
	payload, err := httpRequest("message/send", nil, request)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// SendAppLinkMessage is 发送企业会话链接消息
func SendAppLinkMessage(agentID, touser string, title, text string, picUrl, url string) (*OAPIResponse, error) {
	if agentID == "" {
		agentID = conf.AgentID
	}
	var data OAPIResponse
	request := map[string]interface{}{
		"touser":  touser,
		"agentid": agentID,
		"msgtype": "link",
		"link": map[string]string{
			"messageUrl": url,
			"picUrl":     picUrl,
			"title":      title,
			"text":       text,
		},
	}
	payload, err := httpRequest("message/send", nil, request)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// SendTextMessage is 发送普通文本消息
func SendTextMessage(sender string, cid string, msg string) (*MessageResponse, error) {
	var data MessageResponse
	request := map[string]interface{}{
		"chatid":  cid,
		"sender":  sender,
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": msg,
		},
	}
	payload, err := httpRequest("chat/send", nil, request)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// SendImageMessage is 发送图片消息
func SendImageMessage(sender string, cid string, mediaID string) (*MessageResponse, error) {
	var data MessageResponse
	request := map[string]interface{}{
		"chatid":  cid,
		"sender":  sender,
		"msgtype": "image",
		"image": map[string]string{
			"media_id": mediaID,
		},
	}
	payload, err := httpRequest("chat/send", nil, request)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// SendVoiceMessage is 发送语音消息
func SendVoiceMessage(sender string, cid string, mediaID string, duration string) (*MessageResponse, error) {
	var data MessageResponse
	request := map[string]interface{}{
		"chatid":  cid,
		"sender":  sender,
		"msgtype": "voice",
		"voice": map[string]string{
			"media_id": mediaID,
			"duration": duration,
		},
	}
	payload, err := httpRequest("chat/send", nil, request)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// SendFileMessage is 发送文件消息
func SendFileMessage(sender string, cid string, mediaID string) (*MessageResponse, error) {
	var data MessageResponse
	request := map[string]interface{}{
		"chatid":  cid,
		"sender":  sender,
		"msgtype": "file",
		"file": map[string]string{
			"media_id": mediaID,
		},
	}
	payload, err := httpRequest("chat/send", nil, request)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// SendLinkMessage is 发送链接消息
func SendLinkMessage(sender string, cid string, mediaID string, url string, title string, text string) (*MessageResponse, error) {
	var data MessageResponse
	request := map[string]interface{}{
		"chatid":  cid,
		"sender":  sender,
		"msgtype": "link",
		"link": map[string]string{
			"messageUrl": url,
			"picUrl":     mediaID,
			"title":      title,
			"text":       text,
		},
	}
	payload, err := httpRequest("chat/send", nil, request)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// OverAllActionCardMessage 整体跳转ActionCard
type OverAllActionCardMessage struct {
	Title       string `json:"title"`
	MarkDown    string `json:"markdown"`
	SingleTitle string `json:"single_title"`
	SingleUrl   string `json:"single_url"`
}

// IndependentActionCardMessage 独立跳转ActionCard
type IndependentActionCardMessage struct {
	Title          string                     `json:"title"`
	MarkDown       string                     `json:"markdown"`
	BtnOrientation string                     `json:"btn_orientation"`
	BtnJsonList    []ActionCardMessageBtnList `json:"btn_json_list"`
}

type ActionCardMessageBtnList struct {
	Title     string `json:"title,omitempty"`
	ActionUrl string `json:"action_url,omitempty"`
}

func (m *IndependentActionCardMessage) AppendBtnItem(title string, action_url string) {
	f := ActionCardMessageBtnList{Title: title, ActionUrl: action_url}

	if m.BtnJsonList == nil {
		m.BtnJsonList = []ActionCardMessageBtnList{}
	}

	m.BtnJsonList = append(m.BtnJsonList, f)
}

// OAMessage is the Message for OA
type OAMessage struct {
	URL   string `json:"message_url"`
	PcURL string `json:"pc_message_url"`
	Head  struct {
		BgColor string `json:"bgcolor,omitempty"`
		Text    string `json:"text,omitempty"`
	} `json:"head,omitempty"`
	Body struct {
		Title     string          `json:"title,omitempty"`
		Form      []OAMessageForm `json:"form,omitempty"`
		Rich      OAMessageRich   `json:"rich,omitempty"`
		Content   string          `json:"content,omitempty"`
		Image     string          `json:"image,omitempty"`
		FileCount int             `json:"file_count,omitempty"`
		Author    string          `json:"author,omitempty"`
	} `json:"body,omitempty"`
}

type OAMessageForm struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type OAMessageRich struct {
	Num  string `json:"num,omitempty"`
	Unit string `json:"body,omitempty"`
}

func (m *OAMessage) AppendFormItem(key string, value string) {
	f := OAMessageForm{Key: key, Value: value}

	if m.Body.Form == nil {
		m.Body.Form = []OAMessageForm{}
	}

	m.Body.Form = append(m.Body.Form, f)
}

// SendOAMessage is 发送OA消息
func SendOAMessage(sender string, cid string, msg OAMessage) (*MessageResponse, error) {
	var data MessageResponse
	request := map[string]interface{}{
		"chatid":  cid,
		"sender":  sender,
		"msgtype": "oa",
		"oa":      msg,
	}
	payload, err := httpRequest("chat/send", nil, request)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMessageReadList is 获取已读列表
func GetMessageReadList(messageID string, cursor int, size int) (*MessageReadListResponse, error) {
	var data MessageReadListResponse
	params := url.Values{}
	params.Add("messageId", messageID)
	params.Add("cursor", strconv.Itoa(cursor))
	params.Add("size", strconv.Itoa(size))
	payload, err := httpRequest("chat/getReadList", params, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

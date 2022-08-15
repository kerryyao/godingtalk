package godingtalk

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

var (
	conf  *Config
	cache *gocache.Cache
)

func Init(config *Config) {
	conf = config
	if conf == nil {
		conf = &Config{}
	}
	if conf.BaseURL == "" {
		conf.BaseURL = "oapi.dingtalk.com"
	}

	cache = gocache.New(7200*time.Second, 9000*time.Second)
}

// OAPIResponse is
type OAPIResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// MessageResponse is
type MessageResponse struct {
	OAPIResponse
	MessageID string `json:"messageId"`
}

// MessageResponse is
type MessageReadListResponse struct {
	OAPIResponse
	NextCursor     int64    `json:"next_cursor"`
	ReadUserIdList []string `json:"readUserIdList"`
}

// AccessTokenResponse is
type AccessTokenResponse struct {
	OAPIResponse
	AccessToken string `json:"access_token"`
	Expires     int    `json:"expires_in"`
}

// JsAPITicketResponse is
type JsAPITicketResponse struct {
	OAPIResponse
	Ticket  string
	Expires int `json:"expires_in"`
}

// RefreshAccessToken is to get a valid access token
func RefreshAccessToken(p ...interface{}) (*string, error) {
	var data *AccessTokenResponse
	if cacheData, found := cache.Get("auth"); found {
		data = cacheData.(*AccessTokenResponse)
		return &data.AccessToken, nil
	}

	params := url.Values{}
	params.Add("appkey", conf.AppKey)
	params.Add("appsecret", conf.AppSecret)

	payload, err := httpRequest("gettoken", params, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}

	cache.Set("auth", data, time.Duration(data.Expires)*time.Second)

	return &data.AccessToken, nil
}

// GetJsAPITicket is to get a valid ticket for JS API
func GetJsAPITicket(nonceStr, timestamp, url string) (string, error) {
	var data *JsAPITicketResponse

	if cacheData, found := cache.Get("jsapi"); found {
		data = cacheData.(*JsAPITicketResponse)
	}

	if data == nil {
		payload, err := httpRequest("get_jsapi_ticket", nil, nil)
		if err == nil {
			return "", err
		}
		if err := json.Unmarshal(*payload, data); err != nil {
			return "", err
		}

		cache.Set("jsapi", data, time.Duration(data.Expires)*time.Second)
	}

	config := map[string]string{
		"url":       url,
		"nonceStr":  nonceStr,
		"agentId":   conf.AgentID,
		"timeStamp": timestamp,
		"corpId":    conf.CorpID,
		"ticket":    data.Ticket,
		"signature": Sign(data.Ticket, nonceStr, timestamp, url),
	}
	bytes, _ := json.Marshal(&config)
	return string(bytes), nil
}

// Sign is 签名
func Sign(ticket string, nonceStr string, timeStamp string, url string) string {
	s := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", ticket, nonceStr, timeStamp, url)
	return sha1Sign(s)
}

package oauth

import (
	"encoding/json"
	"net/url"

	"github.com/kerryyao/godingtalk"
	"github.com/kerryyao/godingtalk/contact"
)

// OauthUserInfoByCode 通过免登码获取用户信息
func OauthUserInfoByCode(code string) (*contact.User, error) {
	var data contact.User
	params := url.Values{}
	body := &struct {
		Code string `json:"code"`
	}{
		Code: code,
	}
	payload, err := godingtalk.HttpRequest("topapi/v2/user/getuserinfo", params, body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

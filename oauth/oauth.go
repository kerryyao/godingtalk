package oauth

import (
	"encoding/json"
	"net/url"

	"github.com/kerryyao/godingtalk"
	"github.com/kerryyao/godingtalk/contact"
)

// OauthUserInfoByCode 校验免登录码并换取用户身份
func OauthUserInfoByCode(code string) (*contact.User, error) {
	var data contact.User
	params := url.Values{}
	params.Add("code", code)
	payload, err := godingtalk.HttpRequestWithToken("user/getuserinfo", params, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

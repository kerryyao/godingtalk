package oauth

import (
	"encoding/json"
	"net/url"

	"github.com/kerryyao/godingtalk"
)

// UserInfoByCode 校验免登录码并换取用户身份
func UserInfoByCode(code string) (*godingtalk.User, error) {
	var data godingtalk.User
	params := url.Values{}
	params.Add("code", code)
	payload, err := godingtalk.HttpRequest("user/getuserinfo", params, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

package oauth

import (
	"encoding/json"
	"net/url"

	"github.com/kerryyao/godingtalk"
)

// OauthUserInfoByCode 通过免登码获取用户信息
func OauthUserInfoByCode(code string) (*UserResponse, error) {
	var data UserResponse
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

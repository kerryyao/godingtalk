package oauth

import "github.com/kerryyao/godingtalk"

type UserResponse struct {
	godingtalk.OAPIResponse
	Result struct {
		DeviceID string `json:"device_id"`
		Name     string `json:"name"`
		Sys      bool   `json:"sys"`
		SysLevel int    `json:"sys_level"`
		Unionid  string `json:"unionid"`
		Userid   string `json:"userid"`
	} `json:"result"`
}

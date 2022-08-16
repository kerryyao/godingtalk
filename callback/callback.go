package callback

import (
	"encoding/json"

	"github.com/kerryyao/godingtalk"
)

// Register is 注册事件回调接口
func CallbackRegister(callbacks []string, token string, aes_key string, callbackURL string) (*godingtalk.OAPIResponse, error) {
	data := &godingtalk.OAPIResponse{}
	request := map[string]interface{}{
		"call_back_tag": callbacks,
		"token":         token,
		"aes_key":       aes_key,
		"url":           callbackURL,
	}
	payload, err := godingtalk.HttpRequest("call_back/register_call_back", nil, request)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, data); err != nil {
		return nil, err
	}
	return data, err
}

// Update 更新事件回调接口
func CallbackUpdate(callbacks []string, token string, aes_key string, callbackURL string) (*godingtalk.OAPIResponse, error) {
	data := &godingtalk.OAPIResponse{}
	request := map[string]interface{}{
		"call_back_tag": callbacks,
		"token":         token,
		"aes_key":       aes_key,
		"url":           callbackURL,
	}
	payload, err := godingtalk.HttpRequest("call_back/update_call_back", nil, request)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, data); err != nil {
		return nil, err
	}
	return data, err
}

// Delete is 删除事件回调接口
func CallbackDelete() (*godingtalk.OAPIResponse, error) {
	data := &godingtalk.OAPIResponse{}
	payload, err := godingtalk.HttpRequest("call_back/delete_call_back", nil, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, data); err != nil {
		return nil, err
	}
	return data, err
}

// List is 查询事件回调接口
func CallbackList() (*Callback, error) {
	data := &Callback{}
	payload, err := godingtalk.HttpRequest("call_back/get_call_back", nil, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return data, err
}

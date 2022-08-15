package godingtalk

import "encoding/json"

type Callback struct {
	OAPIResponse
	Token     string
	AES_KEY   string `json:"aes_key"`
	URL       string
	Callbacks []string `json:"call_back_tag"`
}

// Register is 注册事件回调接口
func CallbackRegister(callbacks []string, token string, aes_key string, callbackURL string) (*OAPIResponse, error) {
	data := &OAPIResponse{}
	request := map[string]interface{}{
		"call_back_tag": callbacks,
		"token":         token,
		"aes_key":       aes_key,
		"url":           callbackURL,
	}
	payload, err := httpRequest("call_back/register_call_back", nil, request)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, data); err != nil {
		return nil, err
	}
	return data, err
}

// Update 更新事件回调接口
func CallbackUpdate(callbacks []string, token string, aes_key string, callbackURL string) (*OAPIResponse, error) {
	data := &OAPIResponse{}
	request := map[string]interface{}{
		"call_back_tag": callbacks,
		"token":         token,
		"aes_key":       aes_key,
		"url":           callbackURL,
	}
	payload, err := httpRequest("call_back/update_call_back", nil, request)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, data); err != nil {
		return nil, err
	}
	return data, err
}

// Delete is 删除事件回调接口
func CallbackDelete() (*OAPIResponse, error) {
	data := &OAPIResponse{}
	payload, err := httpRequest("call_back/delete_call_back", nil, nil)
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
	payload, err := httpRequest("call_back/get_call_back", nil, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(*payload, &data); err != nil {
		return nil, err
	}
	return data, err
}

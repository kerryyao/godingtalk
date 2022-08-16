package godingtalk

import "encoding/json"

// DataMessage 服务端加密、解密消息
type DataMessage struct {
	OAPIResponse
	Data string
}

// Encrypt is 服务端加密
func Encrypt(str string) (string, error) {
	var data DataMessage
	request := map[string]interface{}{
		"data": str,
	}
	payload, err := HttpRequest("encryption/encrypt", nil, request)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return "", err
	}
	return data.Data, nil
}

// Decrypt is 服务端解密
func Decrypt(str string) (string, error) {
	var data DataMessage
	request := map[string]interface{}{
		"data": str,
	}
	payload, err := HttpRequest("encryption/decrypt", nil, request)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return "", err
	}
	return data.Data, nil
}

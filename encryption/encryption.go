package encryption

import (
	"encoding/json"

	"github.com/kerryyao/godingtalk"
)

// Encrypt is 服务端加密
func Encrypt(str string) (string, error) {
	var data DataMessage
	request := map[string]interface{}{
		"data": str,
	}
	payload, err := godingtalk.HttpRequestWithToken("encryption/encrypt", nil, request)
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
	payload, err := godingtalk.HttpRequestWithToken("encryption/decrypt", nil, request)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(*payload, &data); err != nil {
		return "", err
	}
	return data.Data, nil
}

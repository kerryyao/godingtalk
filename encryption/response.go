package encryption

import "github.com/kerryyao/godingtalk"

// DataMessage 服务端加密、解密消息
type DataMessage struct {
	godingtalk.OAPIResponse
	Data string
}

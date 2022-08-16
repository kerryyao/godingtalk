package callback

import "github.com/kerryyao/godingtalk"

type Callback struct {
	godingtalk.OAPIResponse
	Token     string
	AES_KEY   string `json:"aes_key"`
	URL       string
	Callbacks []string `json:"call_back_tag"`
}

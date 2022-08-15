package godingtalk

var (
	conf *Config
)

type OAPIResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func Init(config *Config) {
	conf = config
	if conf == nil {
		conf = &Config{}
	}
	if conf.BaseURL == "" {
		conf.BaseURL = "oapi.dingtalk.com"
	}
}

package godingtalk

var (
	conf *Config
)

func Init(config *Config) {
	conf = config
	if conf == nil {
		conf = &Config{}
	}
	if conf.BaseURL == "" {
		conf.BaseURL = "oapi.dingtalk.com"
	}
}

package godingtalk

type Config struct {
	BaseURL  string
	CorpID   string
	ApiToken string

	AgentID   string
	AppKey    string
	AppSecret string

	//社交相关的属性
	SnsAppID       string
	SnsAppSecret   string
	SnsAccessToken string
}

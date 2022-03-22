package common

/*
ApiConfig
api 配置信息
*/
type ApiConfig struct {
	QyWebhook string // 企微 webhook api
}

/*
ThirdPartyToken
第三方 token
*/
type ThirdPartyToken struct {
	GithubToken string // github token
}

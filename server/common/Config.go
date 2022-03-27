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

/*
AntioxidantConfig
配置文件
*/
type AntioxidantConfig struct {
	QyWechat    string `yaml:"QyWechat"`    // 企微 webhook api
	GithubToken string `yaml:"GithubToken"` // github token
}

package model

var Config TomlConfig

type TomlConfig struct {
	Proxy    string   `toml:"proxy"`
	Port     string   `toml:"port"`
	Slack    Slack    `toml:"slack"`
	Mail     Mail     `toml:"mail"`
	Lark     Lark     `toml:"lark"`
	Dingding Dingding `toml:"dingding"`
	Wechat   Wechat   `toml:"wechat"`
}

type Slack struct {
	Open int    `toml:"openSlack"`
	Url  string `toml:"slackUrl"`
}

type Mail struct {
	Open     int    `toml:"openEmail"`
	Host     string `toml:"emailHost"`
	Port     int    `toml:"emailPort"`
	User     string `toml:"emailUser"`
	Password string `toml:"emailPassword"`
	Title    string `toml:"emailTitle"`
	Emails   string `toml:"defaultEmails"`
}

type Lark struct {
	Open int    `toml:"openLark"`
	Url  string `toml:"larkUrl"`
}

type Dingding struct {
	Open int    `toml:"openDingding"`
	Url  string `toml:"ddUrl"`
	All  int    `toml:"ddIsAtAll"`
}

type Wechat struct {
	Open int    `toml:"openWechat"`
	Url  string `toml:"wechatUrl"`
}

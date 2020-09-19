package model

var Config TomlConfig


type TomlConfig struct {
	Proxy string `toml:"proxy"`
	Slack Slack `toml:"slack"`
	Mail Mail `toml:"mail"`
	Lark Lark `toml:"lark"`
	Dingding Dingding `toml:"dingding"`
	Wechat Wechat `toml:"wechat"`
}

type Slack struct {
	Open int `toml:"open-slack"`
	Url string `toml:"slackurl"`
}

type Mail struct {
	Open int `toml:"open-email"`
	Host string `toml:"Email_host"`
	Port int `toml:"Email_port"`
	User string `toml:"Email_user"`
	Password string `toml:"Email_password"`
	Title string `toml:"Email_title"`
	Emails string `toml:"Default_emails"`
}

type Lark struct {
	Open int `toml:"open-lark"`
	Url string `toml:"larkurl"`
}

type Dingding struct {
	Open int `toml:"open-dingding"`
	Url string `toml:"ddurl"`
	All int `toml:"dd_isatall"`

}

type Wechat struct {
	Open int `toml:"open-wechat"`
	Url string `toml:"wechaturl"`
}

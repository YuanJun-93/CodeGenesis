package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	Redis struct {
		Host string
		Type string
		Pass string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Log LogConf
}

type LogConf struct {
	Level    string `json:",default=info"`
	Encoding string `json:",default=json"`
}

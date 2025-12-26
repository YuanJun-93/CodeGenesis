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
	Ai struct {
		ApiKey   string
		Provider string `json:",optional"`
		BaseUrl  string `json:",optional"`
		Model    string `json:",optional"`
	}
	ZapLog    LogConf
	MachineId int64 `json:",default=1"`
}

type LogConf struct {
	Level    string `json:",default=info"`
	Encoding string `json:",default=json"`
}

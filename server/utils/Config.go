package utils

import (
	"github.com/go-ini/ini"
)

type Conf struct {
	Port                    string
	IP                      string
	DomainName              string
	DataSourceName          string
	RedisSourceName         string
	RedisPassword           string
	RedisDB                 int
	RegisterCodeFailureTime int
	EmailAccount            string
	EmailPassword           string
	EmailAuthorizationCode  string
	JWTKey                  string
	ContinuousTime          int
}

var C *Conf

func GetConf() *Conf {
	C = new(Conf)
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		panic(err.Error())
	}
	C.Port = cfg.Section("main").Key("Port").Value()
	C.IP = cfg.Section("main").Key("Ip").Value()
	C.DomainName = cfg.Section("main").Key("Domain").Value()
	C.DataSourceName = cfg.Section("mysql").Key("DataSourceName").Value()
	C.RedisSourceName = cfg.Section("redis").Key("RedisSourceName").Value()
	C.RedisPassword = cfg.Section("redis").Key("RedisPassword").Value()
	C.RedisDB, _ = cfg.Section("redis").Key("RedisPassword").Int()
	C.RegisterCodeFailureTime, _ = cfg.Section("main").Key("RegisterCodeFailureTime").Int()
	C.EmailAccount = cfg.Section("email").Key("EmailAccount").Value()
	C.EmailPassword = cfg.Section("email").Key("EmailPassword").Value()
	C.EmailAuthorizationCode = cfg.Section("email").Key("EmailAuthorizationCode").Value()
	C.JWTKey = cfg.Section("email").Key("JWTKey").Value()
	C.ContinuousTime, _ = cfg.Section("main").Key("ContinuousTime").Int()
	return C
}

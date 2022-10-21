package utils

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Port            string `yaml:"Port"`
	IP              string `yaml:"Ip"`
	DomainName      string `yaml:"DomainName"`
	DataSourceName  string `yaml:"DataSourceName"`
	RedisSourceName string `yaml:"RedisSourceName"`
	RedisPassword   string `yaml:"RedisPassword"`
	RedisDB         int    `yaml:"RedisDB"`
}

var C *Conf

func GetConf() *Conf {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &C)
	if err != nil {
		fmt.Println(err.Error())
	}
	return C
}

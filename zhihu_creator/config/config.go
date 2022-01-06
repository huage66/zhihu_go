package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Settings struct {
	Application Application `json:"application"`
	Database    Database    `json:"database"`
	Redis       Redis       `json:"redis"`
	Kafka       Kafka       `json:"kafka"`
}

type Application struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	Module string `json:"module"`
}

type Database struct {
	Tidb string // tidb的访问入口可以通过nginx等负载均衡组件来合理分配
}

type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

// 这个可能是通知系统
type Kafka struct {
	Addrs []string `json:"addrs"`
}

var (
	Setting Settings
)

func Setup(name string) {
	if len(name) < 1 {
		name = "config.yml"
	}
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(fmt.Sprintf("./config/%s", name))
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&Setting); err != nil {
		panic(err)
	}
}

package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Configs struct {
	Server    ConfigServer    `yaml:"server"`
	Typesense ConfigTypesense `yaml:"typesense"`
}

type ConfigServer struct {
	AppName string `yaml:"appname"`
	Prefork bool   `yaml:"prefork"`
	Port    string `yaml:"port"`
}

type ConfigTypesense struct {
	ApiKey string                 `yaml:"apikey"`
	Nodes  []configTypesenseNodes `yaml:"nodes"`
}

type configTypesenseNodes struct {
	Protocol string `yaml:"protocol"`
	Hostname string `yaml:"hostname"`
	Port     string `yaml:"port"`
}

var Cfg *Configs

func Get() *Configs {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs/.")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	Cfg = &Configs{}
	err = viper.Unmarshal(Cfg)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return Cfg
}

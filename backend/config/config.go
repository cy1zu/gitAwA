package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name                  string `mapstructure:"name"`
	Mode                  string `mapstructure:"mode"`
	Version               string `mapstructure:"version"`
	Port                  string `mapstructure:"port"`
	StartTime             string `mapstructure:"start_time"`
	MachineID             int64  `mapstructure:"machine_id"`
	TestGithubAccessToken string `mapstructure:"test_github_access_token"`
	LLMAppId              string `mapstructure:"llm_app_id"`
	LLMAccessToken        string `mapstructure:"llm_access_token"`
	*LogConfig            `mapstructure:"log"`
	*PostgresConfig       `mapstructure:"Postgres"`
}

type LogConfig struct {
	Mode       string `mapstructure:"mode"`
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
}

type PostgresConfig struct {
	SSLMode  string `mapstructure:"ssl_mode"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	MaxConns int    `mapstructure:"max_conns"`
	MaxIdle  int    `mapstructure:"max_idle"`
}

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return
	}
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal() failed, err:%v\n", err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal() failed, err:%v\n", err)
		}
	})
	return err
}

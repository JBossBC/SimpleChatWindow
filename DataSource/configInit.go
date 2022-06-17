package DataSource

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	MysqlConfig MysqlConfig `mapstructure:"mysql"`
	RedisConfig RedisConfig `mapstructure:"redis"`
}

var config *Config

func init() {
	config = &Config{
		MysqlConfig: MysqlConfig{},
		RedisConfig: RedisConfig{},
	}
	getwd, _ := os.Getwd()
	viper.SetConfigName("dataSource")
	viper.SetConfigFile(getwd + "\\DataSource\\dataSource.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("配置文件读取错误")
	}
	err = viper.Unmarshal(config)
	if err != nil {
		panic("viper 解析错误")
	}
	MysqlInit()
	RedisInit()
}
func GetMysqlConfig() *MysqlConfig {
	return &config.MysqlConfig
}
func GetRedisConfig() *RedisConfig {
	return &config.RedisConfig
}

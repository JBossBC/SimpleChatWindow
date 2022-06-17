package DataSource

import (
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
	"strings"
)

type MysqlConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
}

var mysqlConnection *gorm.DB

func MysqlInit() {
	mysqlConfig := GetMysqlConfig()
	mysqlstr := strings.Builder{}
	mysqlstr.WriteString(mysqlConfig.Username)
	mysqlstr.WriteString(":")
	mysqlstr.WriteString(mysqlConfig.Password)
	mysqlstr.WriteString("@tcp(192.168.43.26:3306)/")
	mysqlstr.WriteString(mysqlConfig.DB)
	mysqlstr.WriteString("?charset=utf8&parseTime=True&loc=Local")
	mysql := mysqlstr.String()
	open, err := gorm.Open("mysql", mysql)
	if err != nil {
		panic("mysql初始化错误")
	}
	mysqlConnection = open
}

func GetMysqlConnection() *gorm.DB {
	return mysqlConnection
}

package Dao

import (
	"chat/DataSource"
	"chat/repository"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strings"
	"sync"
)

type User struct {
}

var (
	userDao  *User
	userOnce sync.Once
)

func init() {
	DataSource.GetMysqlConnection().AutoMigrate(&repository.User{})
}
func NewUserDao() *User {
	userOnce.Do(func() {
		userDao = &User{}
	})
	return userDao
}
func (userDao *User) FindUserByUsername(username string) bool {
	connection := DataSource.GetMysqlConnection()
	user := repository.User{}
	if err := connection.Model(&repository.User{}).Where("username= ?", username).Find(&user).Error.Error(); strings.Compare(err, "record not found") == 0 {
		return false
	}
	return true
}
func (userDao *User) FindUserByUser(user *repository.User) bool {
	connection := DataSource.GetMysqlConnection()
	err := connection.Model(&repository.User{}).Where("username= ? and password= ?", user.Username, user.Password).First(&repository.User{}).Error
	if err != nil {
		return false
	}
	return true
}
func (userDao *User) CreateUser(user *repository.User) bool {
	connection := DataSource.GetMysqlConnection()
	if err := connection.Create(user).Error; err != nil {
		return false
	}
	return true
}

func (userDao *User) ChatOutput() {
	connection := DataSource.GetRedisConnection()
	pubSubConn := redis.PubSubConn{Conn: connection}
	pubSubConn.Subscribe("goclass")
	for {
		switch v := pubSubConn.Receive().(type) {
		case redis.Message:
			fmt.Println(string(v.Data))
		case redis.Subscription:
			fmt.Println(v.Channel)
		case error:
			panic("error")
		}
	}
}
func (userDao *User) ChatInput(user *repository.User, context string) bool {
	connection := DataSource.GetRedisConnection()
	_, err := connection.Do("Publish", "goclass", user.Username+"说:"+context)
	if err != nil {
		return false
	}
	return true
}

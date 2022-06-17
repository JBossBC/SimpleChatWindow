package main

import (
	"chat/Dao"
	_ "chat/DataSource"
	"chat/repository"
	"fmt"
	"github.com/jinzhu/gorm"
	"os/exec"
	"strings"
	"sync"
)

func main() {
	var LogAndRegister string
	fmt.Println("登录还是注册:")
	fmt.Scan(&LogAndRegister)
	if strings.Compare(LogAndRegister, "注册") == 0 {
		registerServer()
	}
	var isLog = false
	var isLogStr string
	var user *repository.User
	for !isLog {
		user, isLog = log()
		if !isLog {
			fmt.Println("账户或者密码错误，是否选择重新登录(是或否)")
			fmt.Scan(&isLogStr)
			if strings.Compare(isLogStr, "否") == 0 {
				panic("退出程序")
			} else {
				continue
			}
		}
		fmt.Println("恭喜你登录成功")
	}
	Chat(user)
}
func Chat(user *repository.User) {
	exec.Command("cmd", "/c", "start", "chatOutPut.exe").Start()

	for {
		var context string
		fmt.Println("您想发送的消息:")
		fmt.Scan(&context)
		Dao.NewUserDao().ChatInput(user, context)

	}
}

func registerServer() {
	fmt.Println("下面进行账户注册")
	var registerWaitGroup = &sync.WaitGroup{}
	var name string
	var username string
	var password string
	var rePassword string
	var flag bool = true
	var userMethod string
	for flag {
		registerWaitGroup.Add(2)
		fmt.Println("输入你的真实名字")
		fmt.Scan(&name)
		fmt.Println("请输入你需要注册的用户名")
		fmt.Scan(&username)
		fmt.Println("请输入密码")
		fmt.Scan(&password)
		fmt.Println("请再次输入你的密码")
		fmt.Scan(&rePassword)
		var ack = make(chan bool, 2)
		go func() {
			if strings.Compare(password, rePassword) != 0 {
				fmt.Println("第二次输入的密码与第一次不匹配,是否选择重新输入(是或否):")
				fmt.Scan(&userMethod)
				if strings.Compare(userMethod, "否") == 0 {
					panic("退出程序")
				}
				ack <- true
				registerWaitGroup.Done()
				return
			}
			ack <- false
			registerWaitGroup.Done()
		}()
		go func() {
			IsUsername := Dao.NewUserDao().FindUserByUsername(username)
			iscontinue := <-ack
			if iscontinue {
				registerWaitGroup.Done()
				return
			}
			if IsUsername {
				fmt.Println("用户名已经存在，是否选择重新输入(是或否):")
				fmt.Scan(&userMethod)
				if strings.Compare(userMethod, "否") == 0 {
					panic("退出程序")
				}
				ack <- true
				registerWaitGroup.Done()
				return
			}
			ack <- false
			registerWaitGroup.Done()
		}()
		registerWaitGroup.Wait()
		flag = <-ack
	}
	user := repository.User{
		Model:    gorm.Model{},
		Username: username,
		Password: password,
		Name:     name,
	}
	CreateSuccessful := Dao.NewUserDao().CreateUser(&user)
	if CreateSuccessful {
		fmt.Println("恭喜你创建账户成功，下面进行登录")
	}
}

func log() (*repository.User, bool) {
	fmt.Println("下面进行账户登录")
	var username string
	var password string
	var name string
	fmt.Println("请输入你的账号")
	fmt.Scan(&username)
	fmt.Println("请输入你的密码")
	fmt.Scan(&password)
	user := &repository.User{
		Username: username,
		Password: password,
		Name:     name,
	}
	return user, Dao.NewUserDao().FindUserByUser(user)

}

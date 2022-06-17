package main

import (
	"chat/Dao"
	"chat/repository"
	"testing"
)

func TestName(t *testing.T) {
	u := &repository.User{
		Name:     "析洋",
		Username: "123",
		Password: "123",
	}
	Dao.NewUserDao().ChatInput(u, "hello")
}

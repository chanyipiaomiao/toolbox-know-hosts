package common

import (
	"time"
	"fmt"
	"os/user"
	"os"
)

// 获取当前时间
func GetNowTime() string {
	t := time.Now()
	return fmt.Sprintf("%d%d%d%d%d%d", t.Year(),t.Month(),t.Day(), t.Hour(), t.Minute(), t.Second())
}

// 获取用户的家目录
func GetUserHome()(string, error)  {
	currentUser, err := user.Current()
	if err == nil {
		return currentUser.HomeDir, nil
	} else {
		home := os.Getenv("HOME")
		if home == ""{
			fmt.Println("User < HOME > Env Not Found")
			os.Exit(1)
		}
		return home, nil
	}
}

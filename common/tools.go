package common

import (
	"time"
	"fmt"
	"os/user"
	"os"
)

func GetNowTime() string {
	t := time.Now()
	return fmt.Sprintf("%d%d%d%d%d%d", t.Year(),t.Month(),t.Day(), t.Hour(), t.Minute(), t.Second())
}

func GetUserHome()(string, error)  {
	currentUser, err := user.Current()
	if err == nil {
		return currentUser.HomeDir, nil
	} else {
		home := os.Getenv("HOME")
		if home == ""{
			fmt.Println("User Home Env Not Found")
			os.Exit(1)
		}
		return home, nil
	}
}

package main

import (
	"./common"
)

const (
	KnownHostsPath string = "/.ssh/known_hosts"
)


func main() {
	home, _ := common.GetUserHome()
	path := home + KnownHostsPath
	knowHost := &common.KnowHost{SrcName: path, DstName:path + common.GetNowTime()}
	knowHost.BackupFile()
	lines := knowHost.FileHandler(knowHost.ReadFile())
	knowHost.WriteFile(lines)
}

package main

import (
	"toolbox-know-hosts/common"
	"toolbox-know-hosts/hosts"
	"strings"
	"fmt"
)

const (
	KnownHostsPath string = "/.ssh/known_hosts"
)


func main() {
	home, _ := common.GetUserHome()
	path := home + KnownHostsPath
	knowHost := &hosts.KnowHost{SrcName: path, DstName: strings.Join([]string{path, "." ,common.GetNowTime()}, "")}
	knowHost.BackupFile()
	lines := knowHost.FileHandler(knowHost.ReadFile())
	knowHost.WriteFile(lines)
	fmt.Println("success.")
}

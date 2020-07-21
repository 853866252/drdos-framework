package plugins

import (
	"drdos/config"
	"drdos/utils"
	"os"
)

// [*] Blacklist防止你输入错误，别误伤自己
var (
	Blacklist []string
)

func init() {
	getFromfile()
}

func getFromfile() error {
	dir, _ := os.Getwd()
	iplist, err := utils.FileLoads(dir + config.Blacklists)
	if err != nil {
		return err
	}
	Blacklist = append(Blacklist, iplist...)
	return nil
}

// [*] You can set your own way

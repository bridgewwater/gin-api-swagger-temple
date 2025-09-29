package sys

import (
	"errors"
	"net"
	"strings"
)

const (
	NetMask4Str  = "4"
	NetMask4     = "ff000000"
	NetMask8Str  = "8"
	NetMask8     = "ff000000"
	NetMask16Str = "16"
	NetMask16    = "ffff0000"
	NetMask24Str = "24"
	NetMask24    = "ffffff00"
	NetMask32Str = "32"
	NetMask32    = "ffffffff"
)

func NetMaskStr2Hex(mask string) (hex string) {
	switch mask {
	case NetMask4Str:
		return NetMask4
	case NetMask8Str:
		return NetMask8
	case NetMask16Str:
		return NetMask16
	case NetMask24Str:
		return NetMask24
	case NetMask32Str:
		return NetMask32
	}

	return
}

// NetworkLocalIP
// get local IPv4 addr with mask NetMask24.
// return: address ipv4 only
func NetworkLocalIP() (ipv4 string, err error) {
	return NetworkLocalIPByMask(NetMask24)
}

// NetworkLocalIPByMask
//
//	get local IP addr
//	return: address ipv4 only
func NetworkLocalIPByMask(netMask string) (ipv4 string, err error) {
	ips, err := getPhysicalInterfaceIPsByMask(netMask)
	if err != nil {
		return "", err
	}

	if len(ips) == 0 {
		return "", errors.New("no valid ipv4 address found")
	}

	return ips[0], nil
}

func getPhysicalInterfaceIPsByMask(netMask string) ([]string, error) {
	var ips []string

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range interfaces {
		// 跳过虚拟接口、回环接口、非启用接口
		if iface.Flags&net.FlagUp == 0 ||
			iface.Flags&net.FlagLoopback != 0 ||
			isVirtualInterface(iface.Name) {
			continue
		}

		addrs, errAddr := iface.Addrs()
		if errAddr != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip.DefaultMask().String() != netMask {
				continue
			}

			// 只获取 IPv4 地址
			if ip != nil && ip.To4() != nil {
				ips = append(ips, ip.String())
			}
		}
	}

	return ips, nil
}

// 判断是否为虚拟接口.
func isVirtualInterface(name string) bool {
	virtualPrefixes := []string{
		"docker", "veth", "br-", "virbr", "cali",
		"flannel", "cni", "lo", "tun", "tap",
	}

	for _, prefix := range virtualPrefixes {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}

	return false
}

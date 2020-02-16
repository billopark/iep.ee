package record

import (
	"github.com/billopark/iep.ee/config"
	"github.com/billopark/iep.ee/internal/util"
	"net"
	"strings"
)

var delims = []rune{'-', '.'}
var domainIP = net.ParseIP(config.Get().DomainIP)
var defaultIPv4 = net.ParseIP(config.Get().DefaultIPv4).To4()
var defaultIPv6 = net.ParseIP(config.Get().DefaultIPv6).To16()

func dupIP(ip net.IP) net.IP {
	if ip == nil {
		return nil
	}

	dup := make(net.IP, len(ip))
	copy(dup, ip)
	return dup
}

func mapIPv4(query string) net.IP {
	fields := util.Split(query, delims...)
	fieldsLen := len(fields)
	if fieldsLen < 4 {
		return dupIP(defaultIPv4)
	}

	ip := net.ParseIP(strings.Join(fields[fieldsLen-4:], "."))
	if ip != nil {
		return ip
	}
	return dupIP(defaultIPv4)
}

// TODO(billopark): Fill
func mapIPv6(query string) net.IP {
	return dupIP(defaultIPv6)
}

func buildIPv4(query string) net.IP {
	domain := config.Get().Domain

	if query == domain {
		if domainIPv4 := domainIP.To4(); domainIPv4 != nil {
			return dupIP(domainIPv4)
		} else {
			return dupIP(defaultIPv4)
		}
	}

	if ip, ok := config.Get().Nss[query]; ok {
		if ipv4 := net.ParseIP(ip).To4(); ipv4 != nil {
			return ipv4
		}
	}

	if !strings.HasSuffix(query, domain) {
		return dupIP(defaultIPv4)
	}

	trimmedQuery := strings.TrimSuffix(query, domain)
	return mapIPv4(trimmedQuery)
}

func buildIPv6(query string) net.IP {
	domain := config.Get().Domain

	if query == domain {
		if domainIPv6 := domainIP.To16(); domainIPv6 != nil {
			return dupIP(domainIPv6)
		} else {
			return dupIP(defaultIPv6)
		}
	}

	if ip, ok := config.Get().Nss[query]; ok {
		if ipv6 := net.ParseIP(ip).To16(); ipv6 != nil {
			return ipv6
		}
	}

	if !strings.HasSuffix(query, domain) {
		return dupIP(defaultIPv4)
	}

	trimmedQuery := strings.TrimSuffix(query, domain)
	return mapIPv6(trimmedQuery)
}

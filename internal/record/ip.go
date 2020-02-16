package record

import (
	"github.com/billopark/iep.ee/config"
	"github.com/billopark/iep.ee/internal/util"
	"net"
	"strings"
)

var delims = []rune{'-', '.'}
var defaultIP = net.ParseIP(config.Get().DefaultIP)

func dupIP(ip net.IP) net.IP {
	dup := make(net.IP, len(ip))
	copy(dup, ip)
	return dup
}

func mapIP(query string) net.IP {
	println(query)
	fields := util.Split(query, delims...)
	fieldsLen := len(fields)
	if fieldsLen < 4 {
		return dupIP(defaultIP)
	}

	ip := net.ParseIP(strings.Join(fields[fieldsLen-4:], "."))
	if ip != nil {
		return ip
	}
	return dupIP(defaultIP)
}

func buildIP(query string) net.IP {
	if ip, ok := config.Get().Nss[query]; ok {
		return net.ParseIP(ip)
	}

	domain := config.Get().Domain
	if !strings.HasSuffix(query, domain) {
		return dupIP(defaultIP)
	}

	trimmedQuery := strings.TrimSuffix(query, domain)
	return mapIP(trimmedQuery)
}

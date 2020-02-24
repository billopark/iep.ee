package dnsserver

import (
	"github.com/billopark/iep.ee/config"
	"github.com/billopark/iep.ee/internal/record"
	"github.com/miekg/dns"
	"log"
	"net"
)

func handleA(m *dns.Msg, req *dns.Msg) {
	rr, err := record.BuildA(req.Question[0].Name, nil)
	if err != nil {
		_ = handleNX(m)
	} else {
		m.Answer = append(m.Answer, rr)
	}
}

func handleAAAA(m *dns.Msg, req *dns.Msg) {
	rr, err := record.BuildAAAA(req.Question[0].Name, nil)
	if err != nil {
		_ = handleNX(m)
	} else {
		m.Answer = append(m.Answer, rr)
	}
}

func handleTXT(m *dns.Msg) {
	m.Answer = append(m.Answer, record.BuildTXT())
}

func handleNS(m *dns.Msg) {
	for ns, ip := range config.Get().Nss {
		m.Answer = append(m.Answer, record.BuildNS(ns))
		r, _ := record.BuildA(ns, net.ParseIP(ip)) // TODO(billopark): error handle
		m.Extra = append(m.Extra, r)
	}
}

func handleSOA(m *dns.Msg) {
	soa, err := record.BuildSOA()
	if err != nil {
		return
	}
	m.Answer = append(m.Answer, soa)
}

func handleNX(m *dns.Msg) error {
	rr, err := record.BuildSOA()
	m.Ns = append(m.Ns, rr)
	m.Rcode = dns.RcodeNameError
	return err
}

func handler(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = false

	switch r.Question[0].Qtype {
	case dns.TypeA, dns.TypeNone:
		handleA(m, r)

	case dns.TypeAAAA:
		handleAAAA(m, r)

	case dns.TypeTXT:
		handleTXT(m)

	case dns.TypeNS:
		handleNS(m)

	default:
		fallthrough
	case dns.TypeSOA:
		handleSOA(m)
	}

	log.Println(m)
	err := w.WriteMsg(m)
	if err != nil {
		log.Println(err)
	}
}

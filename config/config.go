package config

import (
	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	Domain      string            `mapstructure:"DOMAIN"`
	DomainIP    string            `mapstructure:"DOMAIN_IP"`
	DnsPort     uint16            `mapstructure:"DNS_PORT"`
	DefaultIPv4 string            `mapstructure:"DEFAULT_IPV4"`
	DefaultIPv6 string            `mapstructure:"DEFAULT_IPV6"`
	SOA         SOAConfig         `mapstructure:"SOA"`
	Nss         map[string]string `mapstructure:"NSS"`
	WebPort     uint16            `mapstructure:"WEB_PORT"`
}

type SOAConfig struct {
	RName  string `mapstructure:"RNAME"`
	MName  string `mapstructure:"MNAME"`
	Serial string `mapstructure:"SERIAL"`
}

var config *Config
var once sync.Once

func normalize() {
	// FDQNify domain
	if config.Domain != "" {
		config.Domain = dns.Fqdn(config.Domain)
	}

	nss := make(map[string]string)
	for key, value := range config.Nss {
		nss[dns.Fqdn(key)] = value
	}

	config.Nss = nss

	// port
	if config.DnsPort == 0 {
		config.DnsPort = 53
	}

	if config.WebPort == 0 {
		config.WebPort = 80
	}
}

func Get() *Config {
	once.Do(func() {
		config = &Config{}

		// Load config file
		viper.AddConfigPath("./config")
		viper.SetConfigType("yaml")
		err := viper.ReadInConfig() // FindAll and read the config file
		if err != nil {             // Handle errors reading the config file
			log.Fatalf("Fatal error config file: %s \n", err)
		}

		err = viper.Unmarshal(config)
		if err != nil {
			log.Fatalf("unable to decode config into struct, %v", err)
		}

		normalize()
	})

	return config
}

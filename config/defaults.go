package config

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

// Default values below

// NetworkDaemonDefaults Returns the network daemon's default config.
func NetworkDaemonDefaults() map[string]string {
	m := make(map[string]string)
	base, err := GetGladiusBase()
	if err != nil {
		log.Fatal(err)
	}
	m["ContentDirectory"] = filepath.Join(base, "content")
	m["P2PSeedNodeAddress"] = "165.227.16.209"
	m["ControldHostname"] = "localhost"
	m["ControldPort"] = "3001"
	m["ControldProtocol"] = "http"

	return m
}

// CLIDefaults - Returns the CLI's default config.
func CLIDefaults() map[string]string {
	m := make(map[string]string)
	base, err := GetGladiusBase()
	if err != nil {
		log.Fatal(err)
	}
	m["DirLogs"] = filepath.Join(base, "logs")

	return m
}

// MasterNodeDefaults - Returns the Masternode default config.
func MasterNodeDefaults() map[string]string {
	m := make(map[string]string)

	viper.SetEnvPrefix("GLADIUSMN")
	viper.BindEnv("ROUND_ROBIN")
	viper.BindEnv("MININET_CONFIG")

	base, err := GetGladiusBase()
	if err != nil {
		log.Fatal(err)
	}

	m["PoolEthAddress"] = "0xDAcd582c3Ba1A90567Da0fC3f1dBB638D9438e06"

	m["ControldHostname"] = "localhost"
	m["ControldPort"] = "3001"
	m["ControldProtocol"] = "http"

	m["GeoIPDatabaseDir"] = filepath.Join(base, "geoip")
	m["GeoIPDatabaseFile"] = "GeoLite2-City.mmdb"
	m["GeoIPDatabaseMD5URL"] = "http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz.md5"
	m["GeoIPDatabaseURL"] = "http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz"

	return m
}

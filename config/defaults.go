package config

import (
	"log"
	"path/filepath"
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

	return m
}

// ControlDaemonDefaults returns the default config for the control daemon
func ControlDaemonDefaults() map[string]string {
	m := make(map[string]string)
	base, err := GetGladiusBase()
	if err != nil {
		log.Fatal(err)
	}
	m["DirWallet"] = filepath.Join(base, ".gladius/wallet")
	m["DirKeys"] = filepath.Join(base, ".gladius/keys")

	m["BlockchainMarketAddress"] = "0xc4dfb5c9e861eeae844795cfb8d30b77b78bbc38"
	m["BlockchainNodeFactoryAddress"] = "0x85f0129d0b40b0ed15d97b657872b55cf91ae7de"
	m["BlockchainProvider"] = "https://ropsten.infura.io/tjqLYxxGIUp0NylVCiWw"

	return m
}

// CLIDefaults - Returns the CLI's default config.
func CLIDefaults() map[string]string {
	m := make(map[string]string)

	return m
}

// MasterNodeDefaults - Returns the Masternode default config.
func MasterNodeDefaults() map[string]string {
	m := make(map[string]string)

	return m
}

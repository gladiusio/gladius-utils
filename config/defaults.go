package config

import (
	"os"
	"path/filepath"
	"runtime"
)

// Default values below

// NetworkDaemonDefaults - Returns the network daemon's default config.
func NetworkDaemonDefaults() map[string]string {
	m := make(map[string]string)

	// TODO: Fix windows location
	switch runtime.GOOS {
	case "windows":
		m["ContentDirectory"] = filepath.Join(os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"), ".gladius/content")
	case "linux":
		m["ContentDirectory"] = os.Getenv("HOME") + "/.config/gladius/content/"
	case "darwin":
		m["ContentDirectory"] = os.Getenv("HOME") + "/.config/gladius/content/"
	}

	return m
}

func ControlDaemonDefaults() map[string]string {
	m := make(map[string]string)

	// TODO: Fix windows location
	switch runtime.GOOS {
	case "windows":
		m["DirWallet"] = filepath.Join(os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"), ".gladius/wallet")
		m["DirKeys"] = filepath.Join(os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"), ".gladius/keys")
	case "linux":
		m["DirWallet"] = os.Getenv("HOME") + "/.config/gladius/wallet"
		m["DirKeys"] = os.Getenv("HOME") + "/.config/gladius/keys"
	case "darwin":
		m["DirWallet"] = os.Getenv("HOME") + "/.config/gladius/wallet"
		m["DirKeys"] = os.Getenv("HOME") + "/.config/gladius/keys"
	}

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

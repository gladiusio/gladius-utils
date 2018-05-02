package config

import (
	"os"
	"runtime"
)

// Default values below

// NetworkDaemonDefaults - Returns the network daemon's default config.
func NetworkDaemonDefaults() map[string]string {
	m := make(map[string]string)
	user := os.Getenv("USER")
	home := "/home/" + user
	if user == "root" {
		home = "/home/" + os.Getenv("SUDO_USER")
	}

	// TODO: Fix windows location
	switch runtime.GOOS {
	case "windows":
		m["ContentDirectory"] = "/.config/gladius/gladius-networkd"
	case "linux":
		m["ContentDirectory"] = home + "/.config/gladius/gladius-content/"
	case "darwin":
		m["ContentDirectory"] = home + "/.config/gladius/gladius-content/"
	}

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

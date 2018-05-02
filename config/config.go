package config

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// GetString - Wrapper around viper GetString
func GetString(key string) string {
	return viper.GetString(key)
}

// SetupConfig - Sets up, watches, and registers default config
func SetupConfig(configName string, defaults map[string]string) {
	viper.SetConfigName(configName)

	osPaths, err := getOSPaths()
	if err != nil {
		viper.AddConfigPath(".") // Search only for local config
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath(osPaths["config"]) // OS specifc
	}

	for key, value := range defaults {
		viper.SetDefault(key, value)
	}

	err = viper.ReadInConfig() // Find and read the config file
	// Should probably fix this...
	if err != nil {
		if strings.HasPrefix(err.Error(), "Config File") {
			fmt.Errorf("Cannot find config file: %s. Using defaults", err)
		} else { // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s", err))
		}
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}

// getOSPaths - Returns a mapping like "config" -> "/config/path"
func getOSPaths() (map[string]string, error) {
	m := make(map[string]string)
	var err error
	// TODO: Actually get correct filepath for macOS
	if os.Getenv("GLADIUSCONF") == "" {
		user := os.Getenv("USER")
		home := "/home/" + user
		if user == "root" {
			home = "/home/" + os.Getenv("SUDO_USER")
		}

		switch runtime.GOOS {
		case "windows":
			m["config"] = "%APPDATA%/gladius/"
		case "linux":
			m["config"] = home + "/.config/gladius/"
		case "darwin":
			m["config"] = home + "/.config/gladius/"
		default:
			m["config"] = ""
			err = errors.New("Unknown operating system")
		}
	} else {
		m["config"] = os.Getenv("GLADIUSCONF")
	}

	return m, err
}

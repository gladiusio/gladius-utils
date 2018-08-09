package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// coming soon
// var customBase = flag.String("baseDir", "", "custom gladius base directory")

// GetString - Wrapper around viper GetString
func GetString(key string) string {
	return viper.GetString(key)
}

// SetupConfig - Sets up, watches, and registers default config
func SetupConfig(configName string, defaults map[string]string) {
	viper.SetConfigName(configName)

	base, err := GetGladiusBase()
	if err != nil {
		viper.AddConfigPath(".") // Search only for local config
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath(base) // OS specifc
	}

	for key, value := range defaults {
		viper.SetDefault(key, value)
	}

	err = viper.ReadInConfig() // Find and read the config file
	// Should probably fix this...
	if err != nil {
		if strings.HasPrefix(err.Error(), "Config File") {
			log.Printf("Cannot find config file: %s. Using defaults", err)
		} else { // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s", err))
		}
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}

// GetGladiusBase - Returns the base directory
func GetGladiusBase() (string, error) {
	var m string
	var err error

	cmdArg := ""

	if len(os.Args) > 1 {
		if os.Args[1] != "start" && os.Args[1] != "stop" && os.Args[1] != "install" && os.Args[1] != "uninstall" {
			cmdArg = os.Args[1]
		}
	}

	if cmdArg != "" {
		m = cmdArg
	} else if os.Getenv("GLADIUSBASE") != "" {
		m = os.Getenv("GLADIUSBASE")
	} else {
		switch runtime.GOOS {
		case "windows":
			m = filepath.Join(os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"), ".gladius")
		case "linux":
			m = os.Getenv("HOME") + "/.config/gladius"
		case "darwin":
			m = os.Getenv("HOME") + "/.config/gladius"
		default:
			m = ""
			err = errors.New("unknown operating system, can't find gladius base directory. Set the GLADIUSBASE environment variable, or supply the directory as the first argument to add it manually")
		}
	}

	return m, err
}

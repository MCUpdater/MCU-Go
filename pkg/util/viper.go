package util

import (
	"fmt"
	"runtime"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func InitConfig() {
	viper.SetConfigName("config")

	if dir, err := homedir.Expand("~/.MCUpdater/"); err != nil {
		viper.AddConfigPath(dir)
	}
	switch runtime.GOOS {
	case "windows":
		viper.AddConfigPath("$APPDATA\\.MCUpdater\\")
	case "darwin":
		viper.AddConfigPath("$HOME/Library/Application Support/.MCUpdater/")
	default:
		viper.AddConfigPath("$HOME/.MCUpdater/")
	}

	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Unable to read config: %v\n", err)
	} else {
		fmt.Printf("Read config from %s.\n", viper.ConfigFileUsed())
		DumpConfig()
	}
}

func DumpConfig() {
	c := viper.AllSettings()
	if buf, err := yaml.Marshal(c); err != nil {
		fmt.Printf("Unable to dump config.\n")
	} else {
		fmt.Printf(string(buf))
	}
}

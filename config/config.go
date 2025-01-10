package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var GlobalConfig Config

type Config struct {
	Ai Ai `yaml:"ai"`
}

type Ai struct {
	Name    string `yaml:"name"`
	ApiKey  string `yaml:"apikey"`
	BaseUrl string `yaml:"baseurl"`
	Model   string `yaml:"model"`
}

func InitConfig() *Config {
	return loadingLweConfig("", "")
}

func loadingLweConfig(configPath, configName string) *Config {

	if len(configPath) == 0 {
		home, err := os.UserHomeDir()
		if err != nil {
			cobra.CheckErr(fmt.Sprintf("Can not get user home: %v", err))
		}

		configPath = filepath.Join(home, ".config", "lwe")

	}

	if len(configName) == 0 {
		configName = "config"
	}
	viper.SetConfigType("yaml")     // config file type
	viper.SetConfigName(configName) // config file name without extension
	viper.AddConfigPath(configPath) // config dir

	if err := viper.ReadInConfig(); err != nil {
		cobra.CheckErr(err)
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		cobra.CheckErr(fmt.Sprintf("Can not parse config file: %v", err))
	}

	return &GlobalConfig
}

package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yesAnd92/lwe/utils"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Ai Ai `yaml:"ai"`
}

type Ai struct {
	Lang    string `yaml:"lang"`
	Name    string `yaml:"name"`
	ApiKey  string `yaml:"apikey"`
	BaseUrl string `yaml:"baseurl"`
	Model   string `yaml:"model"`
}

func LoadingLweConfig(configPath, configName string) *Config {

	if len(configPath) == 0 {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Can not get user home: %v", err)
		}

		switch os := utils.OsEnv(); os {
		case utils.Mac:
			configPath = filepath.Join(home, ".config", "lwe")
		case utils.Linux:
			configPath = filepath.Join(home, ".config", "lwe")
		case utils.Win:
			// TODO: 2025/1/9 window config dir
		default:
			cobra.CheckErr("Not support this system currently")
		}
	}

	if len(configName) == 0 {
		configName = "config"
	}
	viper.SetConfigType("yaml")     // config file type
	viper.SetConfigName(configName) // config file name without extension
	viper.AddConfigPath(configPath) // config dir

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Can not read config file: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Can not parse config file: %v", err)
	}

	return &config
}

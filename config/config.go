package config

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var config *Config
var configOnce sync.Once

type Config struct {
	Server    ServerConfig `mapstructure:"server" validate:"required"`
	Firestore Firestore    `mapstructure:"firestore" validate:"required"`
	PubSub    PubSub       `mapstructure:"pubsub" validate:"required"`
	Secret    Secret       `mapstructure:"secret" validate:"required"`
}

type ServerConfig struct {
	Port    int `mapstructure:"port" validate:"required"`
	Timeout int `mapstructure:"timeout"`
}

type Firestore struct {
	ProjectID                string `mapstructure:"project-id" validate:"required"`
	Database                 string `mapstructure:"database" validate:"required"`
	TripSummaryCollection    string `mapstructure:"plan-summary-collection" validate:"required"`
	VideoSummaryCollection   string `mapstructure:"video-summary-collection" validate:"required"`
	VideoHighlightCollection string `mapstructure:"video-highlight-collection" validate:"required"`
	QueueHistoryCollection   string `mapstructure:"queue-history-collection" validate:"required"`
	CredentialFilePath       string `mapstructure:"credential-file-path"`
}

type PubSub struct {
	ProjectID          string `mapstructure:"project-id"`
	SubscriptionID     string `mapstructure:"subscription-id"`
	CredentialFilePath string `mapstructure:"credential-file-path"`
}

type Secret struct {
	AuthString string `mapstructure:"auth-string"`
}

func InitConfig() *Config {
	configOnce.Do(func() {
		configPath, ok := os.LookupEnv("API_CONFIG_PATH")
		if !ok {
			logger.Info("API_CONFIG_PATH not found, using default config")
			configPath = "./config"
		}

		configName, ok := os.LookupEnv("API_CONFIG_NAME")
		if !ok {
			logger.Info("API_CONFIG_NAME not found, using default config")
			configName = "config"
		}

		logger.Info("config path:" + configPath)
		logger.Info("config name:" + configName)
		viper.SetConfigName(configName)
		viper.SetConfigType("yaml")
		viper.AddConfigPath(configPath)

		GetSecretValue()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		if err := viper.ReadInConfig(); err != nil {
			logger.Info("config file not found. using default/env config: " + err.Error())
		}
		viper.AutomaticEnv()

		viper.WatchConfig()
		if err := viper.Unmarshal(&config); err != nil {
			panic(err)
		}
	})

	err := config.validate()
	if err != nil {
		panic(fmt.Sprintf("failed to get configs %s", err.Error()))
	}

	return config
}

func GetSecretValue() {
	for _, value := range os.Environ() {
		pair := strings.SplitN(value, "=", 2)
		if strings.Contains(pair[0], "SECRET_") {
			keys := strings.Replace(pair[0], "SECRET_", "secrets.", -1)
			keys = strings.Replace(keys, "_", ".", -1)
			newKey := strings.Trim(keys, " ")
			newValue := strings.Trim(pair[1], " ")
			viper.Set(newKey, newValue)
		}
	}
}

func (c *Config) validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return err
	}
	return nil
}

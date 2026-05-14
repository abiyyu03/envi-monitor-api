package config

import (
	"github.com/spf13/viper"
)

type MQTTConfig struct {
	Host     string `mapstructure:"MQTT_HOST"`
	Username string `mapstructure:"MQTT_USERNAME"`
	Password string `mapstructure:"MQTT_PASSWORD"`
	ClientID string `mapstructure:"MQTT_CLIENT_ID"`
	Port     string `mapstructure:"MQTT_PORT"`
}

func LoadMQTTConfig() MQTTConfig {
	InitEnv()

	var cfg = MQTTConfig{
		Host:     viper.GetString("MQTT_HOST"),
		Username: viper.GetString("MQTT_USERNAME"),
		Password: viper.GetString("MQTT_PASSWORD"),
		ClientID: viper.GetString("MQTT_CLIENT_ID"),
		Port:     viper.GetString("MQTT_PORT"),
	}

	return cfg
}

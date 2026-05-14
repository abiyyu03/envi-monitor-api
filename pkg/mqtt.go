package pkg

import (
	"go-projects/envi-monitor/config"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTT struct {
	mqtt.Client
	Config config.MQTTConfig
}

func InitMQTT(cfg config.MQTTConfig) (mqtt.Client, error) {
	mqttOpts := mqtt.NewClientOptions().
		AddBroker(cfg.Host).
		SetClientID(cfg.ClientID).
		SetUsername(cfg.Username).
		SetPassword(cfg.Password).
		SetCleanSession(true).
		SetConnectTimeout(30 * time.Second).
		SetAutoReconnect(true).
		SetConnectionLostHandler(func(c mqtt.Client, err error) {
			log.Printf("MQTT connection lost: %v", err)
		})

	client := mqtt.NewClient(mqttOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to MQTT broker: %v", token.Error())
	}
	return client, nil
}

func NewMQTT(cfg config.MQTTConfig) (*MQTT, error) {
	client, err := InitMQTT(cfg)
	if err != nil {
		return nil, err
	}

	return &MQTT{
		Client: client,
		Config: cfg,
	}, nil
}

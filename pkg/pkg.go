package pkg

import (
	"go-projects/envi-monitor/config"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Package struct {
	DB    *SQL
	Cache *Redis
	MQTT  mqtt.Client
}

func NewPackage() (Package, error) {
	config.InitEnv()

	var (
		postgresCfg = config.LoadPostgresConfig()
		redisCfg    = config.LoadRedisConfig()
		mqttCfg     = config.LoadMQTTConfig()
	)

	sql, err := NewSQL(postgresCfg)
	if err != nil {
		return Package{}, err
	}

	cache, err := NewRedis(redisCfg)
	if err != nil {
		return Package{}, err
	}

	mqttClient, err := InitMQTT(mqttCfg)
	if err != nil {
		return Package{}, err
	}

	return Package{
		DB:    sql,
		Cache: cache,
		MQTT:  mqttClient,
	}, nil
}

package sensor

import (
	"go-projects/envi-monitor/internal/service"
	"go-projects/envi-monitor/pkg"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Subscriber struct {
	client  mqtt.Client
	service service.Service
}

func New(pkg pkg.Package, svc service.Service) *Subscriber {
	return &Subscriber{
		client:  pkg.MQTT,
		service: svc,
	}
}

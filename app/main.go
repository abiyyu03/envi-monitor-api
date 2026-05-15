package main

import (
	"go-projects/envi-monitor/internal/adapter/inbound/mqtt"
	"go-projects/envi-monitor/internal/adapter/inbound/rest"
	"go-projects/envi-monitor/pkg"
	"go-projects/envi-monitor/pkg/di"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var app = fiber.New()

	container, err := di.Container()
	if err != nil {
		panic(err)
	}

	err = container.Invoke(func(pkg pkg.Package, restIb rest.Inbound, mqttIb mqtt.Inbound) error {
		restIb.ApiRoutes(app)

		// start subscribers
		mqttIb.Sensor.SubSensorData()

		return app.Listen(":8080")
	})
	if err != nil {
		panic(err)
	}
}

package main

import (
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

	err = container.Invoke(func(pkg pkg.Package, inbound rest.Inbound) error {
		inbound.ApiRoutes(app)

		app.Listen(":8080")

		return nil
	})
	if err != nil {
		panic(err)
	}

}

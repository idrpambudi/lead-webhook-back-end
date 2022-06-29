package main

import (
	"leadwebhook/pkg/dependency"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			echo.New,
			dependency.ProvideConfig,
			dependency.ProvideLogger,
			dependency.ProvideMySqlClient,
			dependency.ProvideLeadRepository,
			dependency.ProvideLeadService,
		),
		fx.Invoke(
			dependency.RegisterRoutes,
			dependency.RegisterHooks,
		),
	).Run()
}

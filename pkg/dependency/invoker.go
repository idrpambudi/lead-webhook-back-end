package dependency

import (
	"context"
	"fmt"

	"leadwebhook/cfg"
	"leadwebhook/pkg/handler"
	"leadwebhook/pkg/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func RegisterRoutes(e *echo.Echo, service *service.LeadService) {
	handler.RegisterLeadRoutes(e, service)
}

func RegisterHooks(
	lifecycle fx.Lifecycle, e *echo.Echo, logger *zap.SugaredLogger, config cfg.Configurations,
) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go e.Start(fmt.Sprintf(":%d", config.Port))
				return nil
			},
			OnStop: func(context.Context) error {
				return logger.Sync()
			},
		},
	)
}

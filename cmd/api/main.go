package main

import (
	"context"
	"fmt"

	"github.com/drdofx/talk-parmad/internal/api/controller"
	"github.com/drdofx/talk-parmad/internal/api/database"
	"github.com/drdofx/talk-parmad/internal/api/lib"
	"github.com/drdofx/talk-parmad/internal/api/repository"
	"github.com/drdofx/talk-parmad/internal/api/routes"
	"github.com/drdofx/talk-parmad/internal/api/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		lib.Module,
		database.Module,
		repository.Module,
		services.Module,
		controller.Module,
		routes.Module,
		fx.Invoke(
			startServer,
		),
	)

	app.Run()
}

func startServer(handler *lib.RequestHandler, routes routes.Routes, env *lib.Env, lifecycle fx.Lifecycle) {
	port := env.Port

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting server on port", port)

			go func() {
				handler.Gin.GET("/ping", func(c *gin.Context) {
					c.JSON(200, "pong")
				})

				routes.Setup()

				err := handler.Gin.Run(fmt.Sprintf(":%s", port))
				if err != nil {
					fmt.Println("Error starting server")
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("Stopping server")
			return nil
		},
	})

}

package main

import (
	"fmt"

	"github.com/drdofx/talk-parmad/internal/api/constants"
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
		fx.Provide(
			lib.NewEnv,
			lib.ValidatorInit,
			database.NewDatabase,
			repository.NewUserRepository,
			services.NewUserService,
			controller.NewUserController,
			routes.NewUserRoutes,
			gin.Default,
		),
		fx.Invoke(
			startServer,
		),
	)

	app.Run()
}

func startServer(router *gin.Engine, userRoutes routes.UserRoutes) {
	fmt.Println("Starting server...")

	v1 := router.Group(constants.API_PATH)
	userRoutes.SetupUserRoutes(v1)

	router.Run()
}

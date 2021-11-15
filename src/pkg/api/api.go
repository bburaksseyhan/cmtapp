package api

import (
	"github.com/bburaksseyhan/ctmapp/src/cmd/utils"
	"github.com/bburaksseyhan/ctmapp/src/pkg/handlers"
	"github.com/bburaksseyhan/ctmapp/src/pkg/repository"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	log "github.com/sirupsen/logrus"
)

// Initialize is create new Web Framework, Handlers and Repositories
func Initialize(dbSettings *utils.DbSettings) {

	log.Println("CTM App is starting...")

	//initialize echo
	e := echo.New()

	//register repository with handler
	customerRepo := repository.NewPostgresCustomerRepository(dbSettings)
	customerHandler := handlers.NewCustomerHandler(customerRepo, dbSettings)

	//initialize middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//routings
	e.GET("/health", customerHandler.Health)

	e.GET("/api/v1/customer", customerHandler.List)
	e.POST("/api/v1/customer", customerHandler.Add)
	e.DELETE("/api/v1/customer/:id", customerHandler.Delete)
	e.GET("api/v1/customer/:id", customerHandler.Get)

	e.Logger.Fatal(e.Start(":8091"))
}

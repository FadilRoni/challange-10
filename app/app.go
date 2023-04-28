package app

import (
	"challange-10/config"
	"challange-10/repository"
	"challange-10/route"
	"challange-10/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApp() {
	repo := repository.NewRepo(config.PSQL.DB)
	app := service.NewService(repo)
	// server := handler.NewHttpServer(app)

	route.RegisterApi(router, app)

	port := os.Getenv("APP_PORT")

	router.Run(fmt.Sprintf(":%s", port))
}

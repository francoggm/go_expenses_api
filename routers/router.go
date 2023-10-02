package routers

import (
	"expenses_api/configs"
	"expenses_api/internal/user"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func configureCors() {
	engine.Use(
		cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:8000"},
			AllowMethods:     []string{"GET", "POST"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           24 * time.Hour,
		}),
	)
}

func configureMiddlewares() {

}

func ConfigureRouters(userHandler *user.Handler) {
	engine = gin.Default()

	configureCors()
	configureMiddlewares()

	authGroup := engine.Group("/auth")
	authGroup.POST("/signup", userHandler.Register)
	authGroup.POST("/login", userHandler.Login)
}

func Start() error {
	cfg := configs.GetConfigs()

	return engine.Run(fmt.Sprintf("%s:%s", cfg.APIAddr, cfg.APIPort))
}
package server

import (
	"github.com/gin-gonic/gin"
	"github.com/indeedhat/juniper"
	"github.com/indeedhat/potato/internal/controllers"
	"github.com/indeedhat/potato/internal/env"
	"github.com/indeedhat/potato/internal/store"
)

const (
	ServerKey   = "serve"
	ServerUsage = "Start the http server that provides the webapp"
)

// Serve a web interface/api
func Serve(repo store.TheoryRepository) juniper.CliCommandFunc {
	return func([]string) error {
		controller := controllers.New(repo)

		router := gin.Default()
		router.LoadHTMLGlob("web/views/*")

		router.GET("", controller.Index)
		router.GET("/api/conspiracy", controller.GetConspiracy)

		port := env.Get(env.GinPort)
		if port == "" {
			port = ":8080"
		}
		return router.Run(port)
	}
}

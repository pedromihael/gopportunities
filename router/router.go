package router

import (
	g "github.com/gin-gonic/gin"
)

func Initialize() {
	router := g.Default() //router
	initializeRoutes(router)
  router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	
	v1.GET("/ping", func(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
	v1.GET("/openings", func (ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "List Openings",
		})
	})
	v1.PUT("/opening", func (ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Create Opening",
		})
	})
	v1.DELETE("/opening", func (ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Delete Opening",
		})
	})
	v1.PUT("/opening/edit", func (ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Edit Opening",
		})
	})
}
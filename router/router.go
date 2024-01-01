package router

import (
	"go-crawl/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/get", controller.GetData)
}

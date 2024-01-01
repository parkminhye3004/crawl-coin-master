package main

import (
	"go-crawl/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Routes(r)
	r.Run()
}

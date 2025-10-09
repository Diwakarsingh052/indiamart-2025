package main

import (
	"small-app/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.POST("/signup", handler.Signup)
	r.Run(":8080")
}

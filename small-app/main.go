package main

import (
	"small-app/handler"
	"small-app/internal/users"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	con := users.NewConn()
	h := handler.NewHandler(con)
	r.POST("/signup", h.Signup)
	r.Run(":8080")
}

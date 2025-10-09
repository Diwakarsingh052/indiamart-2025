package handler

import (
	"net/http"
	"small-app/internal/users"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	con *users.Conn
}

func NewHandler(con *users.Conn) *Handler {
	if con == nil {
		panic("users.Conn is nil")
	}
	return &Handler{
		con: con,
	}
}

func (h *Handler) Signup(c *gin.Context) {

	// Check if the size of body is more than 5KB
	if c.Request.ContentLength > 5*1024 {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": "Request body too large. Limit is 5KB"})
		return
	}
	var nu users.NewUser

	err := c.ShouldBindJSON(&nu)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": "Invalid request body"})
		return
	}

	user, err := h.con.CreateUser(nu)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, user)

}

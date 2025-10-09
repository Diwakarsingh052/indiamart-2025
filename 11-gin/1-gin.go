package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//Default returns an Engine instance with the Logger and Recovery middleware already attached.
	//r := gin.Default()
	r := gin.New()

	r.GET("/hello", hello)
	r.Use(gin.Logger())
	r.GET("/search", GetName)

	v1 := r.Group("/v1")

	{

		v1.GET("/user", User) // /v1/user

		// this would apply middleware to all the routes in this group
		// after the below line
		v1.Use(gin.Logger(), gin.Recovery())
		v1.GET("/posts", Posts) // /v1/posts
	}

	r.Run(":8080")

}

func hello(c *gin.Context) {

	//m := map[string]string{}
	//gin.H // it is a map of type string to any
	// setting the status code
	// setting the content type
	// converting the json
	//sending the response

	c.JSON(200, gin.H{
		"message": "Hello World"})

}

func GetName(c *gin.Context) {

	// getting the query params
	// e.g. localhost:8080/search?firstname=John&lastname=Doe // anything after ? is query params
	firstName := c.Query("firstname")
	lastName := c.DefaultQuery("lastname", "none")

	c.String(http.StatusOK, "Hello %s %s", firstName, lastName)
}

func User(c *gin.Context) {
	c.String(http.StatusOK, "user")
}

func Posts(c *gin.Context) {
	c.String(http.StatusOK, "posts")
}

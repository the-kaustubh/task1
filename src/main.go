package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/the-kaustubh/task1/handlers"
	"github.com/the-kaustubh/task1/middlewares"
	"github.com/the-kaustubh/task1/models"
)

// mongodb+srv://kaustubh:<password>@cluster0.annfa.mongodb.net/myFirstDatabase?retryWrites=true&w=majority

func main() {
	r := gin.Default()

	r.Use(middlewares.PrintURLMiddleware)
	r.LoadHTMLGlob("../templates/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/udata", func(c *gin.Context) {
		var newUser models.User
		if err := c.BindJSON(&newUser); err != nil {
			return
		}
		fmt.Println(newUser)

		c.IndentedJSON(http.StatusOK, newUser)
	})

	usersGroup := r.Group("/users")
	usersGroup.GET("/all", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, struct {
			Msg string
		}{
			Msg: "Here",
		})
	})

	usersGroup.POST("/login", handlers.HandleLogin)

	usersGroup.POST("/register", handlers.HandleRegister)

	usersGroup.GET("/accesspage", middlewares.VerifyToken, func(c *gin.Context) {

		c.JSON(http.StatusOK, struct {
			Msg string
		}{
			Msg: "ok",
		})
	})

	r.Static("/public", "../public")

	r.Run(":3000")
}

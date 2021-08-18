package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Mobile    string `json:"mob"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("../templates/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/udata", func(c *gin.Context) {
		var newUser User
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

	r.Static("/public", "../public")

	r.Run(":3000")

}

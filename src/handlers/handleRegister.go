package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/the-kaustubh/task1/database"
	"github.com/the-kaustubh/task1/models"
)

func HandleRegister(c *gin.Context) {
	var user models.User
	collection := database.GetCollection()
	c.BindJSON(&user)
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusCreated, struct {
		Msg string
	}{
		Msg: "User added",
	})
}

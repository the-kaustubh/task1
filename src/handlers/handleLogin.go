package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/the-kaustubh/task1/models"
	"github.com/the-kaustubh/task1/util"
)

func HandleLogin(c *gin.Context) {
	var incomingUser models.User
	// err := c.ShouldBindJSON(&incomingUser)
	err := c.BindJSON(&incomingUser)

	fmt.Println(incomingUser)
	user, err := util.GetUser(incomingUser.Username, incomingUser.Password)
	if err != nil {
		fmt.Println(err)
	}

	token, err := util.SignToken(user.Username)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"msg": "Could not sign token"})
	}

	c.JSON(http.StatusAccepted, gin.H{
		"username": user.Username,
		"token":    token,
	})
}

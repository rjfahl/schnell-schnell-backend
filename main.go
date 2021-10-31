package main

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
)

type credentials struct {
	Username	string	`json: "username"`
	Password	string	`json: "password"`
	Token		string	`json: "token"`
}

var validCredentials = credentials{Username: "c137@onecause.com", Password: "#th@nH@rm#y#r!$100%D0p#", Token: "0237"}

func loginUser(c *gin.Context){
	var newCredentials credentials

	if err := c.BindJSON(&newCredentials); err != nil { return }

	if isValidLoginCredentials(newCredentials) {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusUnauthorized)
	}
}

func isValidLoginCredentials(credentials credentials) bool {
	return credentials.Username == validCredentials.Username &&
		credentials.Password == validCredentials.Password &&
		credentials.Token == time.Now().Format("0304")
}

func main(){
	router := gin.Default()
	router.POST("/login", loginUser)

	router.Run("localhost:8080")
}

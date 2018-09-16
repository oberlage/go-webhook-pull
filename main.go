package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()

	initializeRoutes()
	router.Run() // listen and serve on 0.0.0.0:8080
}

func initializeRoutes() {
	router.GET("/pull", handlePull)
}

// correctToken validates if a givenToken is correct
// when no token is set, it is always correct
func correctToken(givenToken string) bool {
	sysToken := os.Getenv("TOKEN")

	if sysToken != "" && sysToken != givenToken {
		return false
	}
	return true
}

// handlePull handles the pull command and returns HTPP 200 on success with json {"status":"Done"}
// when the token is invalid it returns a 401 message with {"error":"message"}
// when the git pull command fails, it returns a 500 with {"error":"message"}
func handlePull(c *gin.Context) {

	paramToken := c.Query("token")

	if !correctToken(paramToken) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Token is wrong."})
		return
	}

	cmd := exec.Command("git", "pull")
	errCR := cmd.Run()

	if errCR != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", errCR)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Done"})
}

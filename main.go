package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"regexp"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var version = "1.0.0"

var validPullPath string

func init() {

	var pullPathValid bool
	pullPath := os.Getenv("PULLPATH")
	pullPathValid, validPullPath = isValidPullPath(pullPath)
	if !pullPathValid {
		fmt.Printf("Error: pullpath `%s` is not valid or does not exist.\n", pullPath)
		os.Exit(1)
	}

	doPull()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.New()

	initializeRoutes()
	router.Run() // listen and serve on 0.0.0.0:8080
}

func initializeRoutes() {
	router.GET("/pull", handlePull)
	router.POST("/pull", handlePull)
	router.GET("/info", handleInfo)
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

// directoryExists checks if the specified directory exists on the filesystem
func directoryExists(filePath string) (exists bool) {
	exists = true

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		exists = false
	}

	return
}

// isValidPullPath checks wether the specified pullpath is a valid string and if it exists
// when input variable pullPath is empty, is it assumed to be valid and set to ./ in the output variable.
func isValidPullPath(pullPath string) (isValid bool, validPath string) {
	re := regexp.MustCompile(`(?m)^[a-zA-Z0-9_\- ]+$`)

	if re.MatchString(pullPath) && directoryExists(pullPath) {
		isValid = true
		validPath = fmt.Sprintf("%s/", pullPath)
		return
	}

	if pullPath == "" {
		isValid = true
		validPath = "./"
		return
	}

	isValid = false
	return
}

// doPull performs the git pull command
func doPull() (errCR error) {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("cd %s; git pull;", validPullPath))
	errCR = cmd.Run()
	return
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

	errCR := doPull()

	if errCR != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", errCR)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Done"})
}

func handleInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"version": version})
}

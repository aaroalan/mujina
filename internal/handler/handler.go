package handler

import (
	"github.com/aaroalan/mujina/internal/config"
	"github.com/aaroalan/mujina/internal/help"
	"github.com/gin-gonic/gin"
)

// Handler : Single function that will take all the request and return the status and body set in the configuration.
func Handler(c *gin.Context) {
	// Get the global config.
	cfg, ok := c.MustGet("config").(*config.Config)
	help.PanicIfFalse(ok, "error reading config")
	// Match the endpoint with the current request.
	// FullPath is passed in order to match the original route configuration.
	endPoint := MatchEndPoints(c.FullPath(), c.Request.Method, cfg.Endpoints)
	// If request can't be match to a configured endpoint server will return a 400 status.
	if endPoint == nil {
		c.JSON(400, gin.H{"error": "error matching route with config"})
	} else {
		if endPoint.IsNoContent() {
			renderNoContent(c, endPoint.GetStatusCode())
		} else {
			renderContent(c, endPoint.GetStatusCode(), endPoint.BodyPath)
		}
	}
}

// renderNoContent : NoContent response required only the status code.
func renderNoContent(c *gin.Context, statusCode int) {
	c.AbortWithStatus(statusCode)
}

// renderContent : Parses the file in the config and set the value as the body response.
func renderContent(c *gin.Context, statusCode int, bodyPath string) {
	body, err := help.ReadFile(bodyPath)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(statusCode, body)
}

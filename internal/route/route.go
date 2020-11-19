package route

import (
	"fmt"
	"github.com/aaroalan/mujina/internal/config"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	getMethod    = "get"
	postMethod   = "post"
	putMethod    = "put"
	patchMethod  = "patch"
	deleteMethod = "delete"
)

// AddRoutes : Gets all the endpoints from config and creates routes.
func AddRoutes(r *gin.Engine, fn func(ctx *gin.Context), ep []config.Endpoint) {
	routes := uniqRoutes(ep)
	for route, method := range routes {
		printStatus(route, method)
		switch strings.ToLower(method) {
		case getMethod:
			r.GET(route, fn)
		case postMethod:
			r.POST(route, fn)
		case putMethod:
			r.PUT(route, fn)
		case patchMethod:
			r.PATCH(route, fn)
		case deleteMethod:
			r.DELETE(route, fn)
		default:
			r.GET(route, fn)
		}
	}
}

// printStatus : Output a message with the route that is being add.
func printStatus(route string, method string) {
	msg := "Adding route [" + method + "] " + route
	fmt.Println(msg)
}

// uniqRoutes : Routes need to be uniq.
func uniqRoutes(ep []config.Endpoint) map[string]string {
	uniqRoutes := make(map[string]string)
	for i := 0; i < len(ep); i++ {
		uniqRoutes[ep[i].Route] = ep[i].Method
	}
	return uniqRoutes
}

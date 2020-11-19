package handler

import (
	"github.com/aaroalan/mujina/internal/config"
	"strings"
)

// MatchEndPoints : Takes the current path, method and array of endpoints and returns the first one that matches.
func MatchEndPoints(path string, method string, ep []config.Endpoint) *config.Endpoint {
	for i := 0; i < len(ep); i++ {
		match := MatchEndPoint(path, method, ep[i])
		if match {
			return &ep[i]
		}
	}
	return nil
}

// MatchEndPoint : Match the path and method with a given endpoint.
func MatchEndPoint(path string, method string, ep config.Endpoint) bool {
	// Method needs to match e.g get, post.
	if strings.ToLower(method) != ep.GetMethod() {
		return false
	}

	if path != ep.Route {
		return false
	}

	return true
}

package handler

import (
	"backend-test-chenxianhao/user-management/domains"

	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("userHandlers")

// Handler struct holds required services for handler to function
type Handler struct {
	UserFunctions domains.UserFunctions
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R             *gin.Engine
	UserFunctions domains.UserFunctions
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	h := &Handler{
		UserFunctions: c.UserFunctions,
	} // currently has no properties

	// Create an account group
	g := c.R.Group("/v1/users/")
	FollowerRoute(g, h)
	UserEntityRoute(g, h)

}

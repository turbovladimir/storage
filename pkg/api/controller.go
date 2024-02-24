package api

import (
	"github.com/gin-gonic/gin"
	"github.com/turbovladimir/storage.git/pkg/api/crud"
)

type Router struct {
	*gin.Engine
}

func NewRouter() *Router {
	return &Router{gin.Default()}
}

func (c *Router) SetupRoutes() {
	g := c.Group("/api")
	g.POST("crud/:action/:model", crud.Handle)
	g.POST("checker", HandlerChecker)
}

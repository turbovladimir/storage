package routes

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	engin *gin.Engine
}

func NewController(engin *gin.Engine) *Controller {
	return &Controller{engin}
}

func (c Controller) SetupRoutes() {
	g := c.engin.Group("/api")
	g.POST("crud", func(context *gin.Context) {
		HandlerCRUD(context)
	})
}

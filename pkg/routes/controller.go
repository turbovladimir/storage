package routes

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	engin *gin.Engine
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

func NewController(engin *gin.Engine) *Controller {
	return &Controller{engin}
}

func (c Controller) SetupRoutes() {
	g := c.engin.Group("/api")
	g.POST("crud", HandlerCRUD)
	g.POST("checker", HandlerChecker)
}

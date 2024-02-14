package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseErrorInternal(c *gin.Context, err error) {
	c.Error(err)
	c.JSON(http.StatusInternalServerError, Response{
		Success: false,
		Error:   "Something went wrong.",
	})
}

func ResponseBadRequest(c *gin.Context, err error) {
	c.Error(err)
	c.JSON(http.StatusBadRequest, Response{
		Success: false,
		Error:   err.Error(),
		Data:    "",
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Error:   "",
		Data:    data,
	})
}

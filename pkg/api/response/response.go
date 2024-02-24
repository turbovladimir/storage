package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

func ErrorInternal(c *gin.Context, err error) {
	c.Error(err)
	c.JSON(http.StatusInternalServerError, Response{
		Success: false,
		Error:   "Something went wrong.",
	})
}

func BadRequest(c *gin.Context, err error) {
	c.Error(err)
	c.JSON(http.StatusBadRequest, Response{
		Success: false,
		Error:   err.Error(),
		Data:    "",
	})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Error:   "",
		Data:    data,
	})
}

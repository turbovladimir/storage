package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/turbovladimir/storage.git/pkg/repository"
)

type CheckerRequest struct {
	Phone string `json:"phone" binding:"required"`
}

func HandlerChecker(c *gin.Context) {
	defer c.Request.Body.Close()
	var req CheckerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseBadRequest(c, err)

		return
	}

	checker := repository.PhoneChecker{}
	result, err := checker.Check("phone")

	if err != nil {
		ResponseErrorInternal(c, err)

		return
	}

	ResponseSuccess(c, result)
}

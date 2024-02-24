package api

import (
	"github.com/gin-gonic/gin"
	"github.com/turbovladimir/storage.git/pkg/api/request"
	"github.com/turbovladimir/storage.git/pkg/api/response"
	"github.com/turbovladimir/storage.git/pkg/repository"
)

func HandlerChecker(c *gin.Context) {
	defer c.Request.Body.Close()
	var req request.CheckPhone

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err)

		return
	}

	checker := repository.PhoneChecker{}
	result, err := checker.Check("phone")

	if err != nil {
		response.ErrorInternal(c, err)

		return
	}

	response.Success(c, result)
}

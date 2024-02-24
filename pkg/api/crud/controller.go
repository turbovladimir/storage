package crud

import (
	"github.com/gin-gonic/gin"
	"github.com/turbovladimir/storage.git/pkg/api/request"
	"github.com/turbovladimir/storage.git/pkg/api/response"
	"github.com/turbovladimir/storage.git/pkg/db/models"
	"github.com/turbovladimir/storage.git/pkg/repository"
)

const (
	ActionCreate = "create"
	ActionFind   = "find"
	ActionUpdate = "update"
)

func find(c *gin.Context) {
	var req request.Find
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err)
		return
	}

	repo := repository.New()

	if req.Table == "registrations" {
		var m []models.Registration

		repo.Find(req.Filter, &m)

		response.Success(c, &m)
	}
}

func create(c *gin.Context) {
	var req request.Create

	switch c.Param("model") {
	case models.ModelRegistration:
		req.Rows = []*models.Registration{}
	default:
		panic("The model not set in url path.")
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err)
		return
	}
	repo := repository.New()
	p := &repository.CreateParams{
		Models: req.Rows,
	}
	repo.Create(p)

	if p.Error != nil {
		response.ErrorInternal(c, p.Error)
		return
	}

	response.Success(c, "created.")
}

func update(c *gin.Context) {
	var req request.Update
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err)
		return
	}

	//p := repository.UpdateParams{}
}

func Handle(c *gin.Context) {
	defer c.Request.Body.Close()
	switch c.Param("action") {
	case ActionFind:
		find(c)
	case ActionCreate:
		create(c)
	case ActionUpdate:
		update(c)
	default:
		panic("Undefined action")
	}
}

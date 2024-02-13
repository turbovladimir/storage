package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ActionCreate = "Create"
	ActionRead   = "Read"
)

type CRUDRequest struct {
	Action string      `json:"action"`
	Entity string      `json:"entity"`
	Data   interface{} `json:"data"`
}

func HandlerCRUD(c *gin.Context) {
	defer c.Request.Body.Close()

	var req CRUDRequest

	if err := c.BindJSON(&req); err != nil {
		return
	}

	//if err != nil {
	//	c.JSON(http.StatusBadRequest, Response{
	//		Success: false,
	//		Error:   fmt.Sprintf("Invalid request body (%s)", err.Error()),
	//	})
	//
	//	return
	//}
	//
	//if err = json.Unmarshal(body, &reg); err != nil {
	//	c.JSON(http.StatusBadRequest, Response{
	//		Success: false,
	//		Error:   fmt.Sprintf("Invalid request body (%s)", err.Error()),
	//	})
	//
	//	return
	//}
	//
	//err = r.Create(&reg)
	//
	//if err = json.Unmarshal(body, &reg); err != nil {
	//	c.JSON(http.StatusInternalServerError, Response{
	//		Success: false,
	//		Error:   fmt.Sprintf("Something went wrong (%s)", err.Error()),
	//	})
	//
	//	return
	//}

	c.JSON(http.StatusOK, &req.Data)
}

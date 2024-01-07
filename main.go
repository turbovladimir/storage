package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"github.com/turbovladimir/RestApi/pkg/api"
	"github.com/turbovladimir/storage.git/pkg/db"
	mylog "github.com/turbovladimir/storage.git/pkg/log"
	"golang.org/x/exp/slog"
	"net/http"
)

func main() {
	mylog.Init()
	s := db.New()

	defer func() { s.Close() }()

	routes := []api.Route{
		{
			Method: api.MethodGet,
			Path:   "repository/create",
			Handler: func(context *gin.Context) {
				s.CreateTable()
				s.FillTable()

				context.JSON(http.StatusOK, api.ResponseData{
					Data: "Table created and filled successfully",
				})
			},
		},
		{
			Method: api.MethodGet,
			Path:   "repository/list",
			Handler: func(context *gin.Context) {
				students := s.DisplayStudents()
				out, err := json.Marshal(students)

				if err != nil {
					panic("Cannot marshal students")
				}

				context.JSON(http.StatusOK, api.ResponseData{
					Data: string(out),
				})
			},
		},
	}

	midds := []gin.HandlerFunc{
		func(context *gin.Context) {
			defer func() {
				if err := recover(); err != nil {
					errStr := fmt.Sprintf("%s", err)
					slog.Error(`App get panic`, err)
					context.Abort()

					context.JSON(http.StatusInternalServerError, api.ResponseData{
						Error: errStr,
					})
				}
			}()

			context.Next()
		},
	}

	router := api.NewRouter(routes, midds)

	router.Run("8087")
}

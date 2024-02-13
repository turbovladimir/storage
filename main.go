package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	mylog "github.com/turbovladimir/storage.git/pkg/log"
	"github.com/turbovladimir/storage.git/pkg/routes"
	"log"
	"os"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

func main() {
	mylog.Init()
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := gin.Default()
	controller := routes.NewController(e)
	controller.SetupRoutes()

	if err = e.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))); err != nil {
		log.Fatal(err)
	}
}

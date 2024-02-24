package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/turbovladimir/storage.git/pkg/api"
	mylog "github.com/turbovladimir/storage.git/pkg/log"
	"log"
	"os"
)

func main() {
	mylog.Init()
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := api.NewRouter()
	router.SetupRoutes()

	if err = router.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))); err != nil {
		log.Fatal(err)
	}
}

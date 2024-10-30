package main

import (
	"backend/app/db/postgres"
	"backend/config"
	"backend/logger"
)

func main() {
	// init config
	err := config.Init()
	if err != nil {
		panic(err)
	}
	// init logger
	err = logger.Init()
	if err != nil {
		panic(err)
	}
	// init database
	err = postgres.Init("release")

}

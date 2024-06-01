package main

import (
	"3-design-core/api"
	"3-design-core/database"

	env "github.com/joho/godotenv"
)

func main() {
	env.Load()

	database.Init()
	api.InitServer()
}

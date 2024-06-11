package main

import (
	"3-design-core/api"
	"3-design-core/database"
	"3-design-core/migrations"

	env "github.com/joho/godotenv"
)

func main() {
	env.Load()

	database.Init()
	migrations.Migrate()

	api.InitServer()
}

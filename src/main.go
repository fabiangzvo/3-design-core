package main

import (
	"3-design-core/src/api"

	env "github.com/joho/godotenv"
)

func main() {
	env.Load()

	api.InitServer()
}

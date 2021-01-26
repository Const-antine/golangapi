package main

import (
	"goapi/mappings"

	_ "github.com/go-sql-driver/mysql"

	_ "goapi/docs"
)

// @title Golang API project
// @version 1.0
// @description Swagger API for Golang API first project.
// @termsOfService http://swagger.io/terms/

// @contact.name Constantine Bond
// @contact.email my.work.email@gmail.com

// @license.name MIT

// @BasePath /v1
func main() {

	mappings.CreateUrlMappings()
	// Listen and server on 0.0.0.0:8080
	mappings.Router.Run(":8080")

}

// link to docs - https://github.com/swaggo/swag#how-to-use-it-with-gin

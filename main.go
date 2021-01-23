package main

import (
	"goapi/mappings"

	_ "github.com/go-sql-driver/mysql"

	// _ "goapi/main/docs"
	_ "goapi/docs"
)

// @title Golang API project
// @version 1.0
// @description Swagger API for Golang API first project.
// @termsOfService http://swagger.io/terms/

// @contact.name Constantine Bond
// @contact.email kostia.bond4444@gmail.com

// @license.name MIT

// @BasePath /v1
func main() {

	mappings.CreateUrlMappings()
	// Listen and server on 0.0.0.0:8080
	mappings.Router.Run(":8080")

}


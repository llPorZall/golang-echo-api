package main

import (
	"api"
	"fmt"
)

// @title Echo Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @schemes http
func main() {
	fmt.Println("------------ MAIN START -------------")
	api.StartServer()
}

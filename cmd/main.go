package main

import (
	"github.com/esuEdu/reurb-backend/config"
	"github.com/esuEdu/reurb-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	config.InitDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run()
}

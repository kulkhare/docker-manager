package main

import (
	"example.com/dockermanager/configs"
	"example.com/dockermanager/routes"
	"github.com/docker/docker/client"
	"github.com/labstack/echo/v4"
)

var dockerClient *client.Client

func main() {
	e := echo.New()

	//setup routes
	routes.SetupRoutes(e)

	//start server
	e.Logger.Fatal(e.Start(":" + configs.GetEnvVar("APP_PORT")))
}

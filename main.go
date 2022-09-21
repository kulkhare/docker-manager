package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"example.com/dockermanager/configs"
	"example.com/dockermanager/routes"
	"github.com/docker/docker/client"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var dockerClient *client.Client

func main() {
	e := echo.New()

	// Instantiate a template registry and register all html files inside the view folder
	configs.RegisterTemplates(e)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))

	//setup routes
	routes.SetupRoutes(e)

	// Start server
	go func() {
		if err := e.Start(":" + configs.GetEnvVar("APP_PORT")); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 20 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	fmt.Println("Received terminate, graceful shutdown", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

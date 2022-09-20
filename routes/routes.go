package routes

import (
	"example.com/dockermanager/controllers/container"
	"example.com/dockermanager/controllers/front"
	"example.com/dockermanager/controllers/image"
	"example.com/dockermanager/controllers/volume"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func SetupRoutes(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Configure middleware for static files
	staticConfig := middleware.StaticConfig{
		Root:   "static",
		Browse: true,
	}
	e.Use(middleware.StaticWithConfig(staticConfig))

	// Front view pages routes
	e.GET("/containers", front.ContainerPage)
	e.GET("/volumes", front.VolumePage)
	e.GET("/images", front.ImagePage)

	apiRoutes := e.Group("/api/v1")

	// Container routes
	apiRoutes.POST("/container/create", container.CreateContainer)
	apiRoutes.POST("/container/update/:containerId", container.UpdateContainer)
	apiRoutes.POST("/container/save/:containerId", container.SaveContainer)
	apiRoutes.POST("/container/rename/:containerId", container.RenameContainer)
	apiRoutes.GET("/containers", container.ListContainers)
	apiRoutes.GET("/container/start/:containerId", container.StartContainer)
	apiRoutes.GET("/container/restart/:containerId", container.RestartContainer)
	apiRoutes.GET("/container/stop/:containerId", container.StopContainer)
	apiRoutes.GET("/container/pause/:containerId", container.PauseContainer)
	apiRoutes.GET("/container/resume/:containerId", container.ResumeContainer)
	apiRoutes.GET("/container/kill/:containerId", container.KillContainer)
	apiRoutes.GET("/container/logs/:containerId", container.ContainerLogs)

	// Volume routes
	apiRoutes.POST("/volume/create", volume.CreateVolume)
	apiRoutes.GET("/volumes", volume.ListVolumes)
	apiRoutes.GET("/volume/remove/:volumeId", volume.RemoveVolume)
	apiRoutes.GET("/volume/prune/:volumeId", volume.PruneVolumes)

	// Image routes
	apiRoutes.POST("/image/create", image.CreateImage)
	apiRoutes.GET("/images", image.ListImages)
	apiRoutes.GET("/image/search/:term", image.SearchImage)
	apiRoutes.GET("/image/remove/:imageId", image.RemoveImage)
	apiRoutes.GET("/image/build/:imageId", image.BuildImage)
	apiRoutes.GET("/image/save/:imageIds", image.SaveImage)
	apiRoutes.GET("/image/pull/:image", image.PullImage)
	apiRoutes.GET("/image/load/:image", image.LoadImage)
	apiRoutes.GET("/image/import/:image", image.ImportImage)
	apiRoutes.GET("/image/history/:imageId", image.ImageHistory)
	apiRoutes.GET("/image/prune/:imageId", image.ImageHistory)
}

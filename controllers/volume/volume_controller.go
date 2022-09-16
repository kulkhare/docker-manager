package volume

import (
	"context"
	"fmt"
	"net/http"

	"example.com/dockermanager/configs"
	"github.com/docker/docker/api/types/filters"
	"github.com/labstack/echo/v4"
)

var dockerClient = configs.DC

// TODO
func CreateVolume(c echo.Context) error {
	//func (cli *Client) VolumeCreate(ctx context.Context, options volumetypes.VolumeCreateBody) (types.Volume, error)
	return c.JSON(http.StatusOK, &echo.Map{"message": "CreateVolume not implemented yet"})
}

// return list of volumes
func ListVolumes(c echo.Context) error {
	volumes, err := dockerClient.VolumeList(context.Background(), filters.Args{})
	if err != nil {
		panic(err)
	}

	if len(volumes.Volumes) == 0 {
		fmt.Println("No running containers found")
	}

	/* for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	} */
	return c.JSON(http.StatusOK, volumes)
}

// TODO
func RemoveVolume(c echo.Context) error {
	//func (cli *Client) VolumeRemove(ctx context.Context, volumeID string, force bool) error
	return c.JSON(http.StatusOK, &echo.Map{"message": "RemoveVolume not implemented yet"})
}

// TODO
func PruneVolumes(c echo.Context) error {
	//func (cli *Client) VolumesPrune(ctx context.Context, pruneFilters filters.Args) (types.VolumesPruneReport, error)
	return c.JSON(http.StatusOK, &echo.Map{"message": "PruneVolumes not implemented yet"})
}

/*
TODO: Can also be implemented

func (cli *Client) VolumeInspect(ctx context.Context, volumeID string) (types.Volume, error)
func (cli *Client) VolumeInspectWithRaw(ctx context.Context, volumeID string) (types.Volume, []byte, error)
*/

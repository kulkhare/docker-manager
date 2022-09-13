package container

import (
	"context"
	"fmt"
	"net/http"

	"example.com/dockermanager/configs"
	"github.com/docker/docker/api/types"
	"github.com/labstack/echo/v4"
)

var dockerClient = configs.DC

// TODO
func CreateContainer(c echo.Context) error {
	//func (cli *Client) ContainerCreate(ctx context.Context, config *container.Config, ...) (container.ContainerCreateCreatedBody, error)
	return nil
}

// TODO
func UpdateContainer(e echo.Context) error {
	//func (cli *Client) ContainerUpdate(ctx context.Context, containerID string, updateConfig container.UpdateConfig) (container.ContainerUpdateOKBody, error)
	return nil
}

// TODO
func SaveContainer(e echo.Context) error {
	//func (cli *Client) ContainerUpdate(ctx context.Context, containerID string, updateConfig container.UpdateConfig) (container.ContainerUpdateOKBody, error)
	return nil
}

// TODO
func RenameContainer(c echo.Context) error {
	//func (cli *Client) ContainerRename(ctx context.Context, containerID, newContainerName string) error
	return nil
}

// List all or running containers
func ListContainers(c echo.Context) error {

	gA := c.QueryParam("getAll")
	getAll := false
	if gA == "true" {
		getAll = true
	}

	containers, err := dockerClient.ContainerList(context.Background(), types.ContainerListOptions{All: getAll})
	if err != nil {
		panic(err)
	}

	if len(containers) == 0 {
		fmt.Println("No running containers found")
		return c.JSON(http.StatusOK, &echo.Map{"message": "No running containers found"})
	} else {
		/* for _, container := range containers {
			fmt.Printf("%s %s\n", container.ID[:10], container.Image)
		} */
		return c.JSON(http.StatusOK, containers)
	}
}

// TODO
func StartContainer(c echo.Context) error {
	//func (cli *Client) ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error
	return nil
}

// TODO
func RestartContainer(c echo.Context) error {
	//func (cli *Client) ContainerRestart(ctx context.Context, containerID string, timeout *time.Duration) error
	return nil
}

// stop container
func StopContainer(c echo.Context) error {

	containerID := c.Param("containerID")

	err := dockerClient.ContainerStop(context.Background(), containerID, nil)

	if err != nil {
		fmt.Println(containerID + ": Could not stop")
		panic(err)
	}

	return c.JSON(http.StatusOK, &echo.Map{"message": containerID + ": Container stopped"})
}

// TODO
func PauseContainer(c echo.Context) error {
	//func (cli *Client) ContainerPause(ctx context.Context, containerID string) error
	return nil
}

// TODO
func ResumeContainer(c echo.Context) error {
	//func (cli *Client) ContainerUnpause(ctx context.Context, containerID string) error
	return nil
}

// TODO
func RemoveContainer(c echo.Context) error {

	containerID := c.Param("containerID")

	err := dockerClient.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{})

	if err != nil {
		fmt.Println(containerID + ": Could not remove")
		panic(err)
	}

	return c.JSON(http.StatusOK, &echo.Map{"message": containerID + ": Container removed"})
}

// TODO
func KillContainer(c echo.Context) error {
	//func (cli *Client) ContainerKill(ctx context.Context, containerID, signal string) error
	return nil
}

// TODO
func ContainerLogs(c echo.Context) error {
	//func (cli *Client) ContainerLogs(ctx context.Context, container string, options types.ContainerLogsOptions) (io.ReadCloser, error)
	return nil
}

/*
TODO: Can also be implemented

func (cli *Client) ContainerAttach(ctx context.Context, container string, options types.ContainerAttachOptions) (types.HijackedResponse, error)
func (cli *Client) ContainerCommit(ctx context.Context, container string, options types.ContainerCommitOptions) (types.IDResponse, error)
func (cli *Client) ContainerDiff(ctx context.Context, containerID string) ([]container.ContainerChangeResponseItem, error)
func (cli *Client) ContainerExecAttach(ctx context.Context, execID string, config types.ExecStartCheck) (types.HijackedResponse, error)
func (cli *Client) ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error)
func (cli *Client) ContainerExecInspect(ctx context.Context, execID string) (types.ContainerExecInspect, error)
func (cli *Client) ContainerExecResize(ctx context.Context, execID string, options types.ResizeOptions) error
func (cli *Client) ContainerExecStart(ctx context.Context, execID string, config types.ExecStartCheck) error
func (cli *Client) ContainerExport(ctx context.Context, containerID string) (io.ReadCloser, error)
func (cli *Client) ContainerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error)
func (cli *Client) ContainerInspectWithRaw(ctx context.Context, containerID string, getSize bool) (types.ContainerJSON, []byte, error)
func (cli *Client) ContainerResize(ctx context.Context, containerID string, options types.ResizeOptions) error
func (cli *Client) ContainerStatPath(ctx context.Context, containerID, path string) (types.ContainerPathStat, error)
func (cli *Client) ContainerStats(ctx context.Context, containerID string, stream bool) (types.ContainerStats, error)
func (cli *Client) ContainerStatsOneShot(ctx context.Context, containerID string) (types.ContainerStats, error)
func (cli *Client) ContainerTop(ctx context.Context, containerID string, arguments []string) (container.ContainerTopOKBody, error)
func (cli *Client) ContainerWait(ctx context.Context, containerID string, condition container.WaitCondition) (<-chan container.ContainerWaitOKBody, <-chan error)
func (cli *Client) ContainersPrune(ctx context.Context, pruneFilters filters.Args) (types.ContainersPruneReport, error)
*/

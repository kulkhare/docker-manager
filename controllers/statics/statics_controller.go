package statics

import (
	"context"
	"net/http"

	"example.com/dockermanager/configs"
	"github.com/labstack/echo/v4"
)

var dockerClient = configs.DC

func GetDiskUsage(c echo.Context) error {
	usage, err := dockerClient.DiskUsage(context.Background())
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, usage)
}

func GetClientVersion(c echo.Context) error {
	usage := dockerClient.ClientVersion()
	return c.JSON(http.StatusOK, usage)
}

//func (cli *Client) ServerVersion(ctx context.Context) (types.Version, error)
//func (cli *Client) ServiceList(ctx context.Context, options types.ServiceListOptions) ([]swarm.Service, error)
//func (cli *Client) ServiceLogs(ctx context.Context, serviceID string, options types.ContainerLogsOptions) (io.ReadCloser, error)

//func (cli *Client) Events(ctx context.Context, options types.EventsOptions) (<-chan events.Message, <-chan error)
//func (cli *Client) Info(ctx context.Context) (types.Info, error)
//func (cli *Client) Ping(ctx context.Context) (types.Ping, error)
//func (cli *Client) RegistryLogin(ctx context.Context, auth types.AuthConfig) (registry.AuthenticateOKBody, error)

package image

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
func CreateImage(c echo.Context) error {
	//func (cli *Client) ImageCreate(ctx context.Context, parentReference string, options types.ImageCreateOptions) (io.ReadCloser, error)
	return c.JSON(http.StatusOK, &echo.Map{"message": "CreateImage not implemented yet"})
}

// return list of images
func ListImages(c echo.Context) error {

	images, err := dockerClient.ImageList(context.Background(), types.ImageListOptions{})

	if err != nil {
		panic(err)
	}

	if len(images) == 0 {
		fmt.Println("No running images found")
		return c.JSON(http.StatusOK, &echo.Map{"message": "No running images found"})
	} else {
		return c.JSON(http.StatusOK, images)
	}
}

// TODO
func SearchImage(c echo.Context) error {
	//func (cli *Client) ImageSearch(ctx context.Context, term string, options types.ImageSearchOptions) ([]registry.SearchResult, error)
	return c.JSON(http.StatusOK, &echo.Map{"message": "SearchImage not implemented yet"})
}

// TODO
func BuildImage(c echo.Context) error {
	//func (cli *Client) ImageBuild(ctx context.Context, buildContext io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error)
	return c.JSON(http.StatusOK, &echo.Map{"message": "BuildImage not implemented yet"})
}

// TODO
func RemoveImage(c echo.Context) error {
	//func (cli *Client) ImageRemove(ctx context.Context, imageID string, options types.ImageRemoveOptions) ([]types.ImageDeleteResponseItem, error)
	return c.JSON(http.StatusOK, &echo.Map{"message": "RemoveImage not implemented yet"})
}

// TODO
func SaveImage(c echo.Context) error {
	//func (cli *Client) ImageSave(ctx context.Context, imageIDs []string) (io.ReadCloser, error)
	return c.JSON(http.StatusOK, &echo.Map{"message": "SaveImage not implemented yet"})
}

// TODO
func PullImage(c echo.Context) error {
	//func (cli *Client) ImagePull(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error)
	return c.JSON(http.StatusOK, &echo.Map{"message": "PullImage not implemented yet"})
}

// TODO
func LoadImage(c echo.Context) error {
	//func (cli *Client) ImageLoad(ctx context.Context, input io.Reader, quiet bool) (types.ImageLoadResponse, error)
	return c.JSON(http.StatusOK, &echo.Map{"message": "LoadImage not implemented yet"})
}

// TODO
func ImportImage(c echo.Context) error {
	//func (cli *Client) ImageImport(ctx context.Context, source types.ImageImportSource, ref string, ...) (io.ReadCloser, error)
	return c.JSON(http.StatusOK, &echo.Map{"message": "ImportImage not implemented yet"})
}

// TODO
func ImageHistory(c echo.Context) error {
	//func (cli *Client) ImageHistory(ctx context.Context, imageID string) ([]image.HistoryResponseItem, error)
	return c.JSON(http.StatusOK, &echo.Map{"message": "ImageHistory not implemented yet"})
}

// TODO
func PruneImages(c echo.Context) error {
	//func (cli *Client) ImagesPrune(ctx context.Context, pruneFilters filters.Args) (types.ImagesPruneReport, error)
	return c.JSON(http.StatusOK, &echo.Map{"message": "PruneImages not implemented yet"})
}

/*
TODO: Can also be implemented

func (cli *Client) ImageInspectWithRaw(ctx context.Context, imageID string) (types.ImageInspect, []byte, error)
func (cli *Client) ImagePush(ctx context.Context, image string, options types.ImagePushOptions) (io.ReadCloser, error)
func (cli *Client) ImageTag(ctx context.Context, source, target string) error
*/

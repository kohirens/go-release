package github

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	BaseUri        = "https://api.github.com/repos/%s/%s"
	epReleaseAsset = BaseUri + "/releases/%s/assets"
)

type HttpClient interface {
	Do(r *http.Request) (*http.Response, error)
}

// Client GitHub API Client
type Client struct {
	Http       HttpClient
	Org        string
	Repository string
	Token      string
}

var (
	ApiVersion = "2022-11-28"
)

// NewClient Return a GitHub API client.
func NewClient(h HttpClient, org, repository string) *Client {
	return &Client{
		Http:       h,
		Org:        org,
		Repository: repository,
	}
}

// UploadAsset
// see: https://docs.github.com/en/rest/releases/assets?apiVersion=2022-11-28#upload-a-release-asset
func (c *Client) UploadAsset(assetPath string, version string) error {
	if version == "" {
		return fmt.Errorf(stderr.VersionArgEmpty)
	}

	url := fmt.Sprintf(epReleaseAsset, c.Org, c.Repository, version)

	body, errBody := bodyFromFile(assetPath)
	if errBody != nil {
		return errBody
	}

	res, err2 := c.send("POST", url, body)
	if err2 != nil {
		return fmt.Errorf(stderr.CouldNotRequest, url, err2.Error())
	}

	if res.StatusCode != 200 {

	}

	return nil
}

func bodyFromFile(filepath string) (*bytes.Reader, error) {
	body, errBody := os.ReadFile(filepath)
	if errBody != nil {
		return nil, fmt.Errorf(stderr.CouldNotReadFile, filepath, errBody.Error())
	}

	return bytes.NewReader(body), nil
}

func (c *Client) send(method, url string, body io.Reader) (*http.Response, error) {
	req, err1 := http.NewRequest(method, url, body)
	if err1 != nil {
		return nil, fmt.Errorf(stderr.CouldNotBuildRequest, err1.Error())
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", ApiVersion)

	res, err2 := c.Http.Do(req)
	if err2 != nil {
		return nil, fmt.Errorf(stderr.CouldNotRequest, url, err2.Error())
	}

	return res, nil
}

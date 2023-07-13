package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kohirens/stdlib/log"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	BaseUri           = "https://api.github.com/repos/%s/%s"
	epReleaseAsset    = BaseUri + "/releases/%d/assets"
	epReleaseId       = BaseUri + "/releases/tags/%s"
	HeaderApiAccept   = "application/vnd.github+json"
	HeaderApiPostType = "application/octet-stream"
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
	HeaderApiVersion = "2022-11-28"
)

// GetReleaseIdByTag Get a published release with the specified tag.
// see: https://docs.github.com/en/rest/releases/releases?apiVersion=2022-11-28#get-a-release-by-tag-name
// sample: https://api.github.com/repos/OWNER/REPO/releases/tags/TAG
func (c *Client) GetReleaseIdByTag(version string) (*Release, error) {
	if version == "" {
		return nil, fmt.Errorf(stderr.VersionArgEmpty)
	}

	url := fmt.Sprintf(epReleaseId, c.Org, c.Repository, version)

	res, err1 := c.send("GET", url, nil)
	if err1 != nil {
		return nil, fmt.Errorf(stderr.CouldNotRequest, url, err1.Error())
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf(stderr.ReturnStatusCode, res.StatusCode)
	}

	bodyBits, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		return nil, fmt.Errorf(stderr.CouldNotReadResponseBody, err2.Error())
	}

	rel := &Release{}
	if e := json.Unmarshal(bodyBits, rel); e != nil {
		return nil, fmt.Errorf(stderr.CouldNotDecodeJson, e.Error())
	}

	return rel, nil
}

// NewClient Return a GitHub API client.
func NewClient(h HttpClient, org, repository, token string) *Client {
	return &Client{
		Http:       h,
		Org:        org,
		Repository: repository,
		Token:      token,
	}
}

// UploadAsset The endpoint you call to upload release assets is specific to
// your release. Use the upload_url
//
//	see: https://docs.github.com/en/rest/releases/assets?apiVersion=2022-11-28#upload-a-release-asset
func (c *Client) UploadAsset(assetPath string, release *Release) (*Asset, error) {

	basename := filepath.Base(assetPath)
	url := fmt.Sprintf(epReleaseAsset, c.Org, c.Repository, release.Id) + "?name=" + basename

	if release.UploadUrl != "" {
		log.Infof(stdout.UrlRequest, "POST", release.UploadUrl)
	}

	body, errBody := bodyFromFile(assetPath)
	if errBody != nil {
		return nil, errBody
	}

	res, err2 := c.send("POST", url, body)
	if err2 != nil {
		return nil, fmt.Errorf(stderr.CouldNotRequest, url, err2.Error())
	}

	if res.StatusCode != 201 {
		return nil, fmt.Errorf(stderr.ReturnStatusCode, res.StatusCode)
	}

	bodyBits, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		return nil, fmt.Errorf(stderr.CouldNotReadResponseBody, err2.Error())
	}

	ast := &Asset{}
	if e := json.Unmarshal(bodyBits, ast); e != nil {
		return nil, fmt.Errorf(stderr.CouldNotDecodeJson, e.Error())
	}

	return ast, nil
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
	req.Header.Set("Accept", HeaderApiAccept)
	req.Header.Set("X-GitHub-Api-Version", HeaderApiVersion)
	if method == "POST" {
		req.Header.Set("Content-Type", HeaderApiPostType)
	}

	log.Infof(stdout.UrlRequest, method, url)

	res, err2 := c.Http.Do(req)
	if err2 != nil {
		return nil, fmt.Errorf(stderr.CouldNotRequest, url, err2.Error())
	}

	return res, nil
}

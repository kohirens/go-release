package github

var stdout = struct {
	UrlRequest string
}{
	UrlRequest: "%s'ing to %s",
}

var stderr = struct {
	CouldNotBuildRequest     string
	CouldNotDecodeJson       string
	CouldNotReadFile         string
	CouldNotReadResponseBody string
	CouldNotRequest          string
	ReturnStatusCode         string
	VersionArgEmpty          string
}{
	CouldNotBuildRequest:     "could not build a request: %s",
	CouldNotDecodeJson:       "could not decode JSON: %s",
	CouldNotReadFile:         "could not read %s: %s",
	CouldNotReadResponseBody: "could not read response body: %s",
	CouldNotRequest:          "problem with request to %s: %s",
	ReturnStatusCode:         "unexpected return status code %d",
	VersionArgEmpty:          "version argument cannot be an empty string",
}

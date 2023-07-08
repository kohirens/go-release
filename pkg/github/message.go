package github

var stdout = struct {
}{}

var stderr = struct {
	CouldNotBuildRequest string
	CouldNotReadFile     string
	CouldNotRequest      string
	ReturnStatusCode     string
	VersionArgEmpty      string
}{
	CouldNotBuildRequest: "could not build a request: %s",
	CouldNotReadFile:     "could not read %s: %s",
	CouldNotRequest:      "problem with request to %s: %s",
	ReturnStatusCode:     "unexpected return status code %d",
	VersionArgEmpty:      "version argument cannot be an empty string",
}

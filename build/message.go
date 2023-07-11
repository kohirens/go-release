package build

var stdout = struct {
	Artifacts string
	Cs        string
	Built     string
	Wd        string
}{
	Artifacts: "artifacts to upload:",
	Cs:        "cs: %s",
	Built:     "built %s",
	Wd:        "work dir: %s",
}

var stderr = struct {
	CouldNotBuild       string
	CouldNotMakeArchive string
	ExecNameArgEmpty    string
	InvalidSrcDir       string
	MissingArgs         string
	PathNotExist        string
}{
	CouldNotBuild:       "could not build %s %s",
	CouldNotMakeArchive: "could not make archive %s",
	ExecNameArgEmpty:    "executable name argument cannot be empty",
	InvalidSrcDir:       "invalid source directory %s",
	MissingArgs:         "6 arguments are required, see -help and please try gain",
	PathNotExist:        "path does not exist: %s",
}

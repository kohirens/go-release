package build

var stdout = struct {
	Artifacts string
	Cs        string
	Built     string
}{
	Artifacts: "artifacts:",
	Cs:        "cs: %s",
	Built:     "built %s",
}

var stderr = struct {
	CouldNotBuild       string
	CouldNotMakeArchive string
	ExecNameArgEmpty    string
	MissingArgs         string
}{
	CouldNotBuild:       "could not build %s %s\n",
	CouldNotMakeArchive: "could not make archive %s\n",
	ExecNameArgEmpty:    "executable name argument cannot be empty\n",
	MissingArgs:         "4 arguments are required, see -help and please try gain\n",
}

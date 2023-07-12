package zip

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
	CouldNotMakeArchive string
	Generic             string
}{
	CouldNotMakeArchive: "could not make archive %s\n",
	Generic:             "problem %s",
}

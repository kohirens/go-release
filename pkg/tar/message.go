package tar

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
}{
	CouldNotMakeArchive: "could not make archive %s\n",
}

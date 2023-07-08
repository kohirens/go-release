package main

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
	MissingArgs string
}{
	MissingArgs: "4 arguments are required, see -help and please try gain\n",
}

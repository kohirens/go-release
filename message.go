package main

var stdout = struct {
	Version string
}{
	Version: "%s, %s",
}

var stderr = struct {
	MissingArgs string
	NoArgs      string
}{
	MissingArgs: "4 arguments are required, see -help and please try gain",
	NoArgs:      "there is nothing to do because no arguments were given",
}

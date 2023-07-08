package main

var stdout = struct {
	Version string
}{
	Version: "%s, %s",
}

var stderr = struct {
	InvalidCommand      string
	MissingArgs         string
	NoArgs              string
	ParseSubcommandArgs string
}{
	InvalidCommand:      "nothing to do for %s, please run: %s -help",
	MissingArgs:         "4 arguments are required, see -help and please try gain",
	NoArgs:              "there is nothing to do because no arguments were given",
	ParseSubcommandArgs: "parsing subcommand %s args: %s",
}

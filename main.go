package main

//go:generate git-tool-belt semver -save info.go -format go -packageName main -varName bi

import (
	"flag"
	"fmt"
	"github.com/kohirens/go-release/build"
	"github.com/kohirens/stdlib/cli"
	"github.com/kohirens/stdlib/log"
	"os"
)

const (
	appName = "go-release"
	scBuild = "build"
)

type appFlag struct {
	help       bool
	subcommand map[string]*flag.FlagSet
	version    bool
}

type buildInfo struct {
	CurrentVersion string
	CommitHash     string
}

var (
	af = &appFlag{
		subcommand: map[string]*flag.FlagSet{},
	}
	bi = &buildInfo{}
	um = map[string]string{
		"help":    "display program help information",
		scBuild:   "A command to build executables for multiple operating systems.\n\nexample: build <app_name> <build_tag> <directory> <github_org> <repo> <github-token>\n",
		"version": "display program version information",
	}
)

func init() {
	flag.BoolVar(&af.help, "help", false, um["help"])
	af.subcommand[scBuild] = flag.NewFlagSet(scBuild, flag.ExitOnError)
	build.Init(af.subcommand[scBuild])
	flag.BoolVar(&af.version, "version", false, um["version"])
}

func main() {
	var mainErr error

	defer func() {
		if mainErr != nil {
			log.Errf(mainErr.Error())
			os.Exit(1)
		}

		os.Exit(0)
	}()

	flag.Parse()

	if af.help {
		mainErr = cli.Usage(appName, um, af.subcommand)
		fmt.Print("\n\n")
	}

	if af.version {
		fmt.Printf(stdout.Version, bi.CurrentVersion, bi.CommitHash)
	}

	ca := os.Args[1:]
	if len(ca) == 0 {
		mainErr = fmt.Errorf(stderr.NoArgs)
		return
	}

	switch ca[0] {
	case scBuild:
		if e := af.subcommand[scBuild].Parse(ca[1:]); e != nil {
			mainErr = fmt.Errorf(stderr.ParseSubcommandArgs, scBuild, e.Error())
		}
		mainErr = build.Run(af.subcommand[scBuild].Args())
		return
	}

	mainErr = fmt.Errorf(stderr.InvalidCommand, ca[0], appName)
}

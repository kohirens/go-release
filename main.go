package main

import (
	"fmt"
	"github.com/kohirens/go-release/build"
	"github.com/kohirens/go-release/pkg/github"
	"github.com/kohirens/stdlib/log"
	"net/http"
	"os"
	"time"
)

func main() {
	var mainErr error

	defer func() {
		if mainErr != nil {
			log.Errf(mainErr.Error())
			os.Exit(1)
		}

		os.Exit(0)
	}()

	ca := os.Args[1:]
	if len(ca) < 5 {
		mainErr = fmt.Errorf(stderr.MissingArgs)
	}

	srcDir := ca[0]
	execName := ca[1]
	version := ca[2]
	org := ca[3]
	repo := ca[4]

	artifacts, err1 := build.Artifacts(srcDir, execName, build.Platforms)
	if err1 != nil {
		mainErr = err1
		return
	}

	gh := github.NewClient(&http.Client{Timeout: time.Second * 5}, org, repo)
	for _, artifact := range artifacts {
		if e := gh.UploadAsset(artifact, version); e != nil {
			mainErr = e
			break
		}
	}
}

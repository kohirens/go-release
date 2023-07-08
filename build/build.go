package build

import (
	"flag"
	"fmt"
	"github.com/kohirens/go-release/pkg/github"
	"github.com/kohirens/go-release/pkg/tar"
	"github.com/kohirens/go-release/pkg/zip"
	"github.com/kohirens/stdlib"
	"github.com/kohirens/stdlib/cli"
	"github.com/kohirens/stdlib/log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	dirMode = 0777
	ps      = string(os.PathSeparator) // building multiple systems here
)

type Executable struct {
	Dir  string
	Ext  string
	Name string
	Path string
}

type Platform struct {
	Os   string
	Arch string
}

var Platforms = []*Platform{
	{"darwin", "amd64"},
	{"darwin", "arm64"},
	{"linux", "386"},
	{"linux", "amd64"},
	{"windows", "386"},
	{"windows", "amd64"},
}

func Artifacts(srcDir, execName string, platforms []*Platform) ([]string, error) {
	wd, errWd := filepath.Abs(srcDir)
	if errWd != nil {
		return nil, fmt.Errorf(stderr.InvalidSrcDir, srcDir)
	}

	if execName == "" {
		return nil, fmt.Errorf(stderr.ExecNameArgEmpty)
	}

	artifacts := []string{}

	for _, pf := range platforms {
		goOs := pf.Os
		goArch := pf.Arch
		prefix := fmt.Sprintf("%s-%s-%s", execName, goOs, goArch)
		executable, err1 := buildExecutable(wd, execName, prefix, pf)
		if err1 != nil {
			return nil, err1
		}

		// Archiving for Windows
		if goOs == "windows" {
			archivePath, err2 := zip.ArchiveFile(executable.Dir, prefix, execName+executable.Ext)
			if err2 != nil {
				return nil, err2
			}

			artifacts = append(artifacts, archivePath)
			continue
		}

		// Archiving for all other OSes
		archivePath, err3 := tar.ArchiveFile(executable.Dir, prefix, executable.Name)
		if err3 != nil {
			return nil, err3
		}

		artifacts = append(artifacts, archivePath)
	}

	log.Logf(stdout.Artifacts)
	for _, artifact := range artifacts {
		log.Logf("  %v\n", artifact)
	}

	return artifacts, nil
}

func Init(flagSet *flag.FlagSet) {
	// Implement flags for the subcommand when needed here
}

func Run(ca []string) error {
	srcDir := ca[0]
	execName := ca[1]
	version := ca[2]
	org := ca[3]
	repo := ca[4]
	token := ca[6]

	artifacts, err1 := Artifacts(srcDir, execName, Platforms)
	if err1 != nil {
		return err1
	}

	gh := github.NewClient(&http.Client{Timeout: time.Second * 5}, org, repo, token)

	var err2 error
	for _, artifact := range artifacts {
		if e := gh.UploadAsset(artifact, version); e != nil {
			err2 = e
			break
		}
	}

	if err2 != nil {
		return err2
	}

	return nil
}

// buildExecutable produces an executable in the source directory with an
// optional prefix.
//
//	<src>/<prefix>/<execName><ext>
//	Example: go-get-latest/go-get-latest-windows-amd64/ggl.exe
func buildExecutable(srcDir string, execName string, prefix string, pf *Platform) (*Executable, error) {
	ext := ""

	if pf.Os == "windows" {
		ext = ".exe"
	}

	executable := &Executable{
		Ext:  ext,
		Name: execName + ext,
		Dir:  srcDir + ps + prefix,
	}

	if !stdlib.PathExist(executable.Dir) {
		if e := os.Mkdir(executable.Dir, dirMode); e != nil {
			return nil, e
		}
	}

	executable.Path = executable.Dir + ps + execName + ext
	// build an executable for a platform
	so, se, _, cs := cli.RunCommandWithInputAndEnv(
		srcDir,
		"go",
		[]string{"build", "-o", executable.Path, srcDir},
		nil,
		map[string]string{"GOOS": pf.Os, "GOARCH": pf.Arch},
	)

	log.Infof(stdout.Cs, cs)

	if se != nil {
		return nil, fmt.Errorf("%s: %s\n", so, se.Error())
	}

	if !stdlib.PathExist(executable.Path) {
		return nil, fmt.Errorf(stderr.CouldNotBuild, executable.Path)
	}

	log.Logf(stdout.Built, executable.Path)

	return executable, nil
}

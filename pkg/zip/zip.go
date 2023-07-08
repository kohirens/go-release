package zip

import (
	"fmt"
	"github.com/kohirens/stdlib"
	"github.com/kohirens/stdlib/cli"
	"os"
)

const ps = string(os.PathSeparator)

// ArchiveFile Archive a single file and return the path.
func ArchiveFile(workDir, name, filepath string) (string, error) {
	archiveName := name + ".zip"
	so, se, _, _ := cli.RunCommand(
		workDir,
		"zip",
		[]string{archiveName, filepath},
	)
	if se != nil {
		return "", fmt.Errorf("%s: %s\n", so, se.Error())
	}

	archivePath := workDir + ps + archiveName
	if !stdlib.PathExist(archivePath) {
		return "", fmt.Errorf(stderr.CouldNotMakeArchive, archivePath)
	}

	return archivePath, nil
}

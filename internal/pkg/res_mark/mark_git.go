package res_mark

import (
	_ "embed"
	"fmt"
	"github.com/sinlov-go/go-git-tools/git_info"
	"path/filepath"
)

const (
	resMarkFileName = ".git_rev_parse"
	gitHashLen      = 6
)

//go:embed .git_rev_parse
var markGitHeadShort string

func MainProgramRes() string {
	return markGitHeadShort
}

var (
	ErrorNotGitRepo = fmt.Errorf("res mark find out, not find git repo root path")
)

func generateMarkGitHeadShort() error {

	currentFolderPath, err := getCurrentFolderPath()
	if err != nil {
		return err
	}

	// find out git root path
	gitRootPath := findGitRootPath(currentFolderPath)

	if !git_info.IsPathUnderGitManagement(gitRootPath) {
		return ErrorNotGitRepo
	}

	shortName := ""
	var errRes error
	header, err := git_info.RepositoryHeadByPath(gitRootPath)
	if err != nil {
		errRes = err
	} else {
		hash := header.Hash()
		if hash.IsZero() {
			errRes = fmt.Errorf("can not find hash now is zero")
		} else {
			shortName = hash.String()[:gitHashLen]
		}
	}

	targetFile := filepath.Join(currentFolderPath, resMarkFileName)
	err = writeFileByString(targetFile, shortName, true)
	if err != nil {
		return err
	}

	return errRes
}

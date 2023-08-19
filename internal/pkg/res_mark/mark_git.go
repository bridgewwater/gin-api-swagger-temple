package res_mark

import (
	"embed"
	"fmt"
	"github.com/sinlov-go/go-git-tools/git_info"
	"path/filepath"
)

const (
	resMarkFolder   = "res_mark_record"
	resMarkFileName = ".git_rev_parse"
	gitHashLen      = 6
)

var (
	//go:embed res_mark_record
	embedResMark     embed.FS
	markGitHeadShort = ""
)

func MainProgramRes() string {
	if markGitHeadShort == "" {
		resMarkFileContent, err := embedResMark.ReadFile(filepath.Join(resMarkFolder, resMarkFileName))
		if err == nil {
			markGitHeadShort = string(resMarkFileContent)
		} else {
			markGitHeadShort = "000000"
		}
	}

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

	markGitHeadShort = shortName
	targetFile := filepath.Join(currentFolderPath, resMarkFolder, resMarkFileName)
	err = writeFileByString(targetFile, shortName, true)
	if err != nil {
		return err
	}

	return errRes
}

package res_mark

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
)

func findGitRootPath(path string) string {
	return filepath.Dir(filepath.Dir(filepath.Dir(path)))
}

// getCurrentFolderPath
//
//	can get run path this golang dir
func getCurrentFolderPath() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("can not get current file info")
	}
	return filepath.Dir(file), nil
}

func writeFileByString(path string, content string, coverage bool) error {
	return writeFileByByte(path, []byte(content), os.FileMode(0766), coverage)
}

// writeFileByByte
//
//	write bytes to file
//	path most use Abs Path
//	data []byte
//	fileMod os.FileMode(0766) os.FileMode(0666) os.FileMode(0644)
//	coverage true will coverage old
func writeFileByByte(path string, data []byte, fileMod fs.FileMode, coverage bool) error {
	if !coverage {
		exists, err := pathExists(path)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("not coverage, which path exist %v", path)
		}
	}
	parentPath := filepath.Dir(path)
	if !pathExistsFast(parentPath) {
		err := os.MkdirAll(parentPath, fileMod)
		if err != nil {
			return fmt.Errorf("can not writeFileByByte at new dir at mode: %v , at parent path: %v", fileMod, parentPath)
		}
	}
	err := os.WriteFile(path, data, fileMod)
	if err != nil {
		return fmt.Errorf("write data at path: %v, err: %v", path, err)
	}
	return nil
}

// pathExistsFast
//
//	path exists fast
func pathExistsFast(path string) bool {
	exists, _ := pathExists(path)
	return exists
}

// pathExists
//
//	path exists
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

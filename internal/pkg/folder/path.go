package folder

import (
	"fmt"
	"os"
)

// PathExists path exists
//
//nolint:golint,unused
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// PathExistsFast path exists fast
//
//nolint:golint,unused
func PathExistsFast(path string) bool {
	exists, _ := PathExists(path)
	return exists
}

// PathIsDir path is dir
//
//nolint:golint,unused
func PathIsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// RmDir remove dir by path
//
//nolint:golint,unused
func RmDir(path string, force bool) error {
	if force {
		return os.RemoveAll(path)
	}
	exists, err := PathExists(path)
	if err != nil {
		return err
	}
	if exists {
		return os.RemoveAll(path)
	}
	return fmt.Errorf("remove dir not exist path: %s , use force can cover this err", path)
}

// Mkdir will use FileMode 0766
//
//nolint:golint,unused
func Mkdir(path string) error {
	err := os.MkdirAll(path, os.FileMode(0766))
	if err != nil {
		return fmt.Errorf("fail MkdirAll at path: %s , err: %v", path, err)
	}
	return nil
}

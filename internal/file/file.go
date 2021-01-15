package file

import (
	"errors"
	"os"
	"path/filepath"
)

func CreateDirByPath(filePath string) error {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			dirName, fileName := filepath.Split(filePath)
			if dirName == "" || fileName == "" {
				return errors.New("dir or file is empty")
			}

			err = os.MkdirAll(dirName, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func CreateDir(fileDir string) error {
	_, err := os.Stat(fileDir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(fileDir, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

package path

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListFiles(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if info.IsDir() {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path) // Do something with the file
			return nil
		})
		if err != nil {
			return err
		}
	} else {
		// If it's a file, do something with it
		fmt.Println(path)
	}

	return nil
}

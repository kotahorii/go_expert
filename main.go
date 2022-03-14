package main

import (
	"os"
	"path/filepath"
)

func main() {

}

func Walk(dir string, f func(b []byte, err error) error) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		path := filepath.Join(dir, file.Name())
		if !file.IsDir() {
			b, err := os.ReadFile(path)
			if err := f(b, err); err != nil {
				return err
			}
			continue
		}
		if err := Walk(path, f); err != nil {
			return err
		}
	}
	return nil
}

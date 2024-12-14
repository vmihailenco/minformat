package main

import (
	"go/format"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-toolsmith/minformat"
)

func main() {
	if err := filepath.WalkDir(".", func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == "." {
			return nil
		}
		if info.IsDir() {
			if strings.HasPrefix(info.Name(), ".") {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		src, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		src, err = minformat.Source(src)
		if err != nil {
			return err
		}

		src, err = format.Source(src)
		if err != nil {
			return err
		}

		if err := os.WriteFile(path, src, 0644); err != nil {
			return err
		}

		return nil
	}); err != nil {
		panic(err)
	}
}

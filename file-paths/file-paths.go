package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// `filepath.Join` constructs paths in a portable way
	pl := fmt.Println
	p := filepath.Join("dir1", "dir2", "filename")
	pl("p:", p)

	// `filepath.Join` normalizes paths
	pl(filepath.Join("dir1//", "filename"))
	pl(filepath.Join("dir1/../dir1", "filename"))

	pl("Dir(p):", filepath.Dir(p))
	pl("Base(p):", filepath.Base(p))

	pl(filepath.IsAbs("dir/file"))
	pl(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename)
	pl(ext)

	pl(strings.TrimSuffix(filename, ext))

	// `filepath.Rel` finds a relative path between a `base` and the `target`
	// It return an error if the target canno be made relative to base.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	pl(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	pl(rel)
}

package files

import (
	"binker/md-blog/structs"
	"fmt"
	"os"
	"path/filepath"
)

func fileGetter(basePath string) []string {
	var paths []string

	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	return paths
}

func createPostStruct(path string) structs.Post {

	filename := filepath.Base(path)

	return structs.Post{
		Title:    filename,
		FilePath: path,
	}
}

func createBlogStruct(basePath string) structs.Blog {

}

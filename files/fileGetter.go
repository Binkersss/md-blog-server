package files

import (
	"binker/md-blog/structs"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CreateBlogStruct(basePath string) structs.Blog {
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

	var blog structs.Blog
	for _, path := range paths {
		post := createPostStruct(path)
		blog.Posts = append(blog.Posts, post)
	}

	return blog
}

func createPostStruct(path string) structs.Post {

	filename := filepath.Base(path)

	if strings.HasSuffix(filename, ".md") {
		filename = filename[:len(filename)-3]
	}
	return structs.Post{
		Title:    filename,
		FilePath: path,
	}
}

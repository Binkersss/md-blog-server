package main

import (
	"binker/md-blog/files"
	"binker/md-blog/structs"
	"fmt"
)

func main() {

	var basePath string
	basePath = "/home/binker/dev/go/md-blog/content"

	p1 := structs.Post{
		Title:    "First Blog Post",
		FilePath: basePath + "/blog_1.md",
	}

	fmt.Println("Hello from main")

	fmt.Println(p1.Title)
}

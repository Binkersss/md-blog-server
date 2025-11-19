package main

import (
	"binker/md-blog/files"
	"binker/md-blog/structs"
	"fmt"
)

func main() {

	var basePath string
	basePath = "/home/binker/dev/go/md-blog/content"

	var blog structs.Blog = files.CreateBlogStruct(basePath)
	fmt.Println("Hello from main")

	fmt.Println(blog)
}

package main

import (
	"binker/md-blog/files"
	"binker/md-blog/parser"
	"binker/md-blog/server"
	"binker/md-blog/structs"
	"fmt"
)

func main() {

	var basePath string
	basePath = "/home/binker/dev/go/md-blog/content"

	var blog structs.Blog = files.CreateBlogStruct(basePath)
	fmt.Println("Hello from main")

	htmlBytes := parser.ParseMdFile(blog.Posts[0].FilePath)

	html := structs.PostData{
		Title: blog.Posts[0].Title,
		HTML:  htmlBytes,
	}

	fmt.Println(html.Title)

	server.Server()

	server.PostHtml(html)

	fmt.Println(html)
}

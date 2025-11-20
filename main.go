package main

import (
	"binker/md-blog/files"
	"binker/md-blog/parser"
	"binker/md-blog/server"
	"binker/md-blog/structs"
	"fmt"
	"os"
	"time"
)

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Too many args; just basepath of blog files")
		os.Exit(1)
	}

	var basePath string = args[0]
	var blog structs.Blog = files.CreateBlogStruct(basePath)
	fmt.Println("Hello from main")

	htmlBytes := parser.ParseMdFile(blog.Posts[0].FilePath)

	html := structs.PostData{
		Title: blog.Posts[0].Title,
		HTML:  htmlBytes,
	}

	fmt.Println(html.Title)

	go server.Server()

	time.Sleep(200 * time.Millisecond)

	err := server.PostHtml(html)
	if err != nil {
		fmt.Println("POST error:", err)
	}

	fmt.Println("Post uploaded. Visit in broswer:")
	fmt.Printf("http://<hostname>:8080/posts/title:%s\n", html.Title)

	select {}
}

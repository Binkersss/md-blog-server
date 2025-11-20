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

	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Wrong args just basepath of blog files, basepath of static files")
		os.Exit(1)
	}

	var basePath string = args[0]
	var staticPath string = args[1]
	var blog structs.Blog = files.CreateBlogStruct(basePath)
	fmt.Println("Hello from main")

	htmlBytes := parser.ParseMdFile(blog.Posts[0].FilePath, "./templates/styles.html")

	html := structs.PostData{
		Title: blog.Posts[0].Title,
		HTML:  htmlBytes,
	}

	fmt.Println(html.Title)

	go server.Server(staticPath)

	time.Sleep(200 * time.Millisecond)

	fmt.Println("Server running at port 8080")

	err := server.PostHtml(html)
	if err != nil {
		fmt.Println("POST error:", err)
	}

	select {}
}

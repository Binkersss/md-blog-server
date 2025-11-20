package server

import (
	"binker/md-blog/structs"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

var posts = make(map[string][]byte)
var mu sync.RWMutex

func postHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Posts: %s", posts)

	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "missing title", http.StatusBadRequest)
		return
	}

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusInternalServerError)
		return
	}

	mu.Lock()
	posts[title] = buf.Bytes()
	mu.Unlock()

	fmt.Printf("Stored '%s' (%d bytes)\n", title, len(buf.Bytes()))

	fmt.Fprintf(w, "Post received: %s", title)
}

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	const prefix = "/posts/title:"
	if !strings.HasPrefix(r.URL.Path, prefix) {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	title := r.URL.Path[len(prefix):]

	mu.RLock()
	html, exists := posts[title]
	mu.RUnlock()

	if !exists {
		http.Error(w, "post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(html)
}

func PostHtml(data structs.PostData) error {
	url := fmt.Sprintf("http://localhost:8080/posts?title=%s", data.Title)
	resp, err := http.Post(url, "text/html", bytes.NewBuffer(data.HTML))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned status %d", resp.StatusCode)
	}
	return nil
}

func Server() {
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			postHandler(w, r)
		} else if r.Method == http.MethodGet {
			if strings.HasPrefix(r.URL.Path, "/posts/title:") {
				getPostHandler(w, r)
			} else {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

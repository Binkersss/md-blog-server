package server

import (
	"binker/md-blog/structs"
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
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

	fmt.Printf("Received post with title: %s\nHTML:\n%s", title, buf.String())

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Post received: %s", title)
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
	http.HandleFunc("/posts", postHandler)
	http.HandleFunc("/functs", PostHtml)
	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

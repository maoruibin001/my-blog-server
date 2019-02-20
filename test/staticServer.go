package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("/Users/ruibin/blog/my-blog/dist")))
	http.ListenAndServe(":8080", nil)
}

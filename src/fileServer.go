package main

import (
	"net/http"
)

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		f, err := os.Open("public" + r.URL.Path)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			log.Println(err)
// 		}
// 		defer f.Close()
// 		var contentType string

// 		file := r.URL.Path

// 		switch {
// 		case strings.HasSuffix(file, "css"):
// 			contentType = "text/css"
// 		case strings.HasSuffix(file, "html"):
// 			contentType = "text/html"
// 		case strings.HasSuffix(file, "png"):
// 			contentType = "image/png"
// 		default:
// 			contentType = "text/plain"
// 		}
// 		w.Header().Add("Content-Type", contentType)
// 		io.Copy(w, f)
// 	})
// 	http.ListenAndServe(":3000", nil)
// }

// another way to do this 👆 is to make use of the built-in file server

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public"+r.URL.Path)
	})
	http.ListenAndServe(":3000", nil)
}

// or server up static files alone with just one line of code
// http.LIstenAndServe(":3000", http.FileServer(http.Dir("public")))

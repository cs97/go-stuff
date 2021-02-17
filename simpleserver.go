package main

import (
 	"fmt"
	"os"
    "path/filepath"
	"net/http"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
            fmt.Println(err)
    }

	http.Handle("/", http.FileServer(http.Dir(dir)))
	http.ListenAndServe(":8080", nil)
}

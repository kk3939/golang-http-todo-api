package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hellow World!")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "/ access is success!")
		fmt.Println("logged.")
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}

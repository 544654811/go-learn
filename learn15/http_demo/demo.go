package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "%v \n", r.Form)
	fmt.Fprintf(w, "url: %s \n", r.URL)
	fmt.Fprintf(w, "scheme: %s \n", r.URL.Scheme)
	fmt.Fprintf(w, "hello world ... \n")
}

func main() {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http listen failed... err: ", err)
	}
}

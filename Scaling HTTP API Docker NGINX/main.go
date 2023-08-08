package main

import (
	"fmt"
	"io"
	"net/http"
)

func rootPath(w http.ResponseWriter, r *http.Request){
	fmt.Println("Request is in  the Server")
	io.WriteString(w, "You're here\n")
}

func handleRequests() {
    http.HandleFunc("/", rootPath)
    err := http.ListenAndServe("simpleapi:23480", nil)
	if err !=nil{
		fmt.Print(err)
	}

}

func main() {
    handleRequests()
}

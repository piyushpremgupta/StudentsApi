package main

import (
	"App/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/getStudents/", handler.GetStudents)

	fmt.Println("Starting Server")
	http.ListenAndServe(":5001", nil)
}

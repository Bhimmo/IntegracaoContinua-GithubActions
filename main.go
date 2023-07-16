package main

import (
    "log"
    "fmt"
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Home Page")
}

func HandleRequest() {
    http.HandleFunc("/", Home)
    log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	HandleRequest()
}
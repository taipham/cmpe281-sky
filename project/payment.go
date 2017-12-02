package main

import ("fmt" 
        "net/http")

func indexHandler(w http.ResponseWriter, r * http.Request) {
    fmt.Fprintf(w, "<h1> Payment page </h1>")
}

func main() {
   fmt.Println("Payment")
    http.HandleFunc("/", indexHandler)
    http.ListenAdnServe(":8000", nil)
}

package main

import (

   hq "cmpe281-sky/project/main.go"

)

func storePageHandler(response http.ResponseWriter, request *http.Request) {

}

var route = mux.NewRouter()

func main() {

   route.HandleFunc("/store", storePageHandler)

}

package main

import (



)

func storePageHandler(response http.ResponseWriter, request *http.Request) {

}

var route = mux.NewRouter()

func main() {

   route.HandleFunc("/store", storePageHandler)

}
package main

import (
        "net/html"
        "fmt"
	    "log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
        "github.com/gorilla/mux"
)

const landingPage = `
<h1>Login</h1>
<form method="post" action="/login">
   <label for="name"> Username</label>
   <input type="text" id="name" name="name">
   <label for="password">Password</label>
   <input type="password" id="password" name="password">
   <button type="submit">Login</button>
</form>   
`
func landingPageHandler(response httmp.ResponseWriter, request *http.Request) {
   fmt.Fprintf(response, landingPage)
}

var route = mux.NewRouter()

func main() {

   route.HandleFunc("/", landingPageHandler)
   http.Handle("/", route)
   http.ListenAndServe(":8080", nil)

}
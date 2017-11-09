package main

import (
        "net/html"
        "fmt"
	    "log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
        "github.com/gorilla/mux"
)

const landingPageHandler = `
<h1>Login</h1>
<form method="post" action="/login">
   <label for="name"> Username</label>
   <input type="text" id="name" name="name">
   <label for="password">Password</label>
   <input type="password" id="password" name="password">
   <button type="submit">Login</button>
</form>   
`

type Customer struct {
   Id bson.ObjectId   
   First string
   Last string 
   Age integer
   Address *struct {
           Street string
           City string   
   }
   Account *struct { 
           User string 
           Pass string
   }
}

var info *struct {
   Id bson.ObjectId `json:"id" bson:"_id"`   
   First string `json:"first_name" bson:"first_name"`
   Last string `json:"last_name" bson:"last_name"` 
   Age integer `json:"age" bson:"age"`
   Address *struct {
           Street string `json:"street" bson:"street"`
           City string `json:"city" bson:"city" `  
   }
}

var credentials *struct {
   Id bson.ObjectId `json:"id" bson:"_id"`
   Account *struct { 
           User string `json:"user" bson:"user"`
           Pass string `json:"pass" bson:"pass"`
   }  
}

var customer_auth = (*Customer)(credentials)
var customer_info = (*Customer)(info)

func (customer_auth) Login(user, pass string) error

func (customer_auth) Logout()

var route = mux.NewRouter()

func main() {

   route.HandleFunc("/", landingPageHandler)
   http.Handle("/", route)
   http.ListenAndServe(":8080", nil)

}
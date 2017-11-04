package main

import (
        "net/html"
        "fmt"
	"log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

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

func main() {



}

package main

import (
        "fmt"
	"log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

type Credential struct {

   Username admin
   Password admin

}

type Song struct {

   ID bson.ObjectId `json:"id" bson:"_id"`
   Title string
   Album string
   Artist string

}

func (db *Database) Login(user, pass string) error

func (db *Database) Logout()

func main() {



}

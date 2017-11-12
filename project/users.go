// Handle all user information
package main

import (
        "fmt"
	    "log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

type user struct {
	Id 			bson.ObjectId 	`json:"id" bson:"_id"`
	First 		string        	`json:"first_name" bson:"first_name"`
	Last   		string        	`json:"last_name" bson:"last_name"`
	Username	string			`json:"user_name" bson:"user_name"`			
	Age   		int     		`json:"age" bson:"age"`
}

func handleRead(w http.ResponseWriter, r *http.Request) {
  db := context.Get(r, “database”).(*mgo.Session)
  // load the comments
  var comments []*comment
  if err := db.DB(“commentsapp”).C(“comments”).
    Find(nil).Sort(“-when”).Limit(100).All(&comments); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  // write it out
  if err := json.NewEncoder(w).Encode(comments); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func main() {
	// connect to the database
	db, err := mgo.Dial("localhost")
	if err != nil {
	log.Fatal("cannot dial mongo", err)
	}
	defer db.Close() // clean up when we’re done
	// Adapt our handle function using withDB
	h := Adapt(http.HandlerFunc(handle), withDB(db))
	// add the handler
	http.Handle("/comments", context.ClearHandler(h))
	// start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
	log.Fatal(err)
	}
}
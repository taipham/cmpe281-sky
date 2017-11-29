// Handle all user information
package main

import (
	    "log"
        "fmt"
	    "encoding/json"
	    "net/http"

	    "goji.io"
    	"goji.io/pat"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

type User struct {
	id 			string 	        `json:"id"` 
	First 		string        	`json:"first_name"` 
	Last   		string        	`json:"last_name"` 
	Username	string			`json:"user_name"` 			
	Age   		int     		`json:"age"` 
}

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {  
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(code)
    fmt.Fprintf(w, "{message: %q}", message)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {  
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(code)
    w.Write(json)
}
/*
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
}*/

func ensureIndex(s *mgo.Session) {  
    session := s.Copy()
    defer session.Close()

    c := session.DB("cmpe281").C("users")

    index := mgo.Index{
        Key:        []string{"ID"},
        Unique:     true,
        DropDups:   true,
        Background: true,
        Sparse:     true,
    }
    err := c.EnsureIndex(index)
    if err != nil {
        panic(err)
    }
}

func allUsers(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {  
    return func(w http.ResponseWriter, r *http.Request) {
        session := s.Copy()
        defer session.Close()

        c := session.DB("cmpe281").C("users")

        var users []User
        err := c.Find(bson.M{}).All(&users)
        if err != nil {
            ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
            log.Println("Failed get all users: ", err)
            return
        }

        respBody, err := json.MarshalIndent(users, "", "  ")
        if err != nil {
            log.Fatal(err)
        }

        ResponseWithJSON(w, respBody, http.StatusOK)
    }
}

func addUser(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {  
    return func(w http.ResponseWriter, r *http.Request) {
        session := s.Copy()
        defer session.Close()

        var user User
        decoder := json.NewDecoder(r.Body)
        err := decoder.Decode(&user)
        if err != nil {
            ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
            return
        }

        c := session.DB("cmpe281").C("users")

        err = c.Insert(user)
        if err != nil {
            if mgo.IsDup(err) {
                ErrorWithJSON(w, "Book with this id already exists", http.StatusBadRequest)
                return
            }

            ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
            log.Println("Failed insert user: ", err)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Location", r.URL.Path+"/"+ user.id)
        w.WriteHeader(http.StatusCreated)
    }
}

func userById(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {  
    return func(w http.ResponseWriter, r *http.Request) {
        session := s.Copy()
        defer session.Close()

        id := pat.Param(r, "id")

        c := session.DB("cmpe281").C("users")

        var user User
        err := c.Find(bson.M{"Id": id}).One(&user)
        if err != nil {
            ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
            log.Println("Failed find user: ", err)
            return
        }

        if user.id == "" {
            ErrorWithJSON(w, "User not found", http.StatusNotFound)
            return
        }

        respBody, err := json.MarshalIndent(user, "", "  ")
        if err != nil {
            log.Fatal(err)
        }

        ResponseWithJSON(w, respBody, http.StatusOK)
    }
}

func updateUser(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {  
    return func(w http.ResponseWriter, r *http.Request) {
        session := s.Copy()
        defer session.Close()

        id := pat.Param(r, "id")

        var user User
        decoder := json.NewDecoder(r.Body)
        err := decoder.Decode(&user)
        if err != nil {
            ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
            return
        }

        c := session.DB("cmpe281").C("users")

        err = c.Update(bson.M{"id": id}, &user)
        if err != nil {
            switch err {
            default:
                ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
                log.Println("Failed update user: ", err)
                return
            case mgo.ErrNotFound:
                ErrorWithJSON(w, "User not found", http.StatusNotFound)
                return
            }
        }

        w.WriteHeader(http.StatusNoContent)
    }
}

func deleteUser(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {  
    return func(w http.ResponseWriter, r *http.Request) {
        session := s.Copy()
        defer session.Close()

        id := pat.Param(r, "id")

        c := session.DB("cmpe281").C("users")

        err := c.Remove(bson.M{"id": id})
        if err != nil {
            switch err {
            default:
                ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
                log.Println("Failed delete user: ", err)
                return
            case mgo.ErrNotFound:
                ErrorWithJSON(w, "User not found", http.StatusNotFound)
                return
            }
        }

        w.WriteHeader(http.StatusNoContent)
    }
}


func main() {
	session, err := mgo.Dial("db")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    session.SetMode(mgo.Monotonic, true)
    ensureIndex(session)

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/users"), allUsers(session))
	mux.HandleFunc(pat.Post("/users"), addUser(session))
    mux.HandleFunc(pat.Get("/users/id"), userById(session))
    mux.HandleFunc(pat.Put("/users/id"), updateUser(session))
    mux.HandleFunc(pat.Delete("/users/id"), deleteUser(session))
	//route.HandleFunc("/", landingPageHandler)
	//route.HandleFunc("/listAll", listAllUserHandler)
   

	// add the handler
	//http.Handle("/comments", context.ClearHandler(h))
	// start the server
	http.ListenAndServe("0.0.0.0:8080", mux)
	
}

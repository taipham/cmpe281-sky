// Product handling
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
	SongsTitle 		string        	`json:"songs_title"` 
	Singer   		string        	`json:"singer"` 
	Price	string			`json:"price"` 			
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

func ensureIndex(s *mgo.Session) {  
    session := s.Copy()
    defer session.Close()

    c := session.DB("cmpe281").C("products")

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

        c := session.DB("cmpe281").C("products")

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

        c := session.DB("cmpe281").C("products")

        err = c.Insert(user)
        if err != nil {
            if mgo.IsDup(err) {
                ErrorWithJSON(w, "This User ID already exists", http.StatusBadRequest)
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

        c := session.DB("cmpe281").C("products")

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

        c := session.DB("cmpe281").C("products")

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

        c := session.DB("cmpe281").C("products")

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
	mux.HandleFunc(pat.Get("/products"), allUsers(session))
	mux.HandleFunc(pat.Post("/products"), addUser(session))
    mux.HandleFunc(pat.Get("/products/id"), userById(session))
    mux.HandleFunc(pat.Put("/products/id"), updateUser(session))
    mux.HandleFunc(pat.Delete("/products/id"), deleteUser(session))
	
    // start the server
	http.ListenAndServe("0.0.0.0:8080", mux)
	
}

package main

import (
        "net/html"
        "fmt"
	    "log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
        "github.com/gorilla/mux"
	// import mongodb package with alias mongo
	mongo "/app/data/db"
        // to utilize users.go with alias user
	user "cmpe281-sky/project/users.go"
	//to utilize music.go with alias music
        music "cmpe281-sky/project/music.go"
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

func landingPageHandler(response http.ResponseWriter, request *http.Request) {

   fmt.Fprintf(response, landingPage)

}

const homePage = `
<h1>Home</h1>
<hr>
<small>User: %s</small>
<form method="post" action="/store">
   <button type="submit">Store</button>
</form>
<form method="post" action="/logout">
   <button type="submit">Logout</button>
</form>   
`

func homePageHandler(response http.ResponseWriter, request *http.Request) {
   userName := getUserName(request)
   if userName != "" {
      fmt.Fprintf(response, homePage, userName)
   }
   else {
      http.Redirect(response, request, "/", 302)
   }
}

func loginHandler(response http.ResponseWriter, request *http.Request) {

   name := request.FormValue("name")
   pass := request.FormValue("password")
   redirectTarget := "/"
   if name != "" && pass != "" {
      setSession(name, response)
      redirectTarget = "/home"
   }
   http.Redirect(response, request, redirectTarget, 302)   
}

func logoutHandler(response http.ResponseWriter, request *http.Request) {
   clearSession(response)
   http.Redirect(response, request, "/", 302)
}

func setSession(userName string, response http.ResponseWriter) {
   value := map[string] string{
      "name": userName,
   }
   if encoded, err := cookieHandler.Encode("session", value); err == nil {
      cookie := &http.Cookie{
         Name:   "session",
         Value:  encoded,   
         Path:   "/",
      }
      http.SetCookie(response, cookie)
   }
}

func getUserName(request *http.Request) (userName string) {
   if cookie, err := request.Cookie("session"); err == nil {
      cookieValue := make(map[string]string)
      if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
         userName = cookieValue["name"]
      }
   }
   
   return userName   
}

func clearSession(response http.ResponseWriter) {
   cookie := &http.Cookie{
      Name:   "session",
      Value:  "",
      Path:   "/",
      MaxAge: -1,
   }
   http.SetCookie(response, cookie)
}

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
   route.HandleFunc("/home", homePageHandler)
   
   route.HandleFunc("/login", loginHandler).Methods("POST")
   route.HandleFunc("/logout", logoutHandler).Methods("POST")
   
   http.Handle("/", route)
   http.ListenAndServe(":8080", nil)

}

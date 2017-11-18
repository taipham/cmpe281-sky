package main

import( "log"
	"net/http")

func main() {
	router :=productdb.NewRouter()

	m_origins := handlers.AllowedOrigins([]string{"*"}) 
	m_methods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

 	log.Fatal(http.ListenAndServe(":8000",
  	handlers.CORS(m_origins, m_methods)(router)))
}

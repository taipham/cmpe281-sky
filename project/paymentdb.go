package main

import( "log"
		"net/http"
		"time"
		"fmt"
		"gopkg.in/mgo.v2"
		"gopkg.in/mgo.v2/bson"
	)

type Song struct {
	songID		uint8
	songName	string
	singerName	string
	musicGroup	string
	albumName	string
	price		float32

}

type Order struct {
	orderID			uint8
	createDate		string
	customerName	string
	customerID		uint8
	songID			uint8
	songName		string
	totalCost		float32
	statusOrder		string
}
func main() {

	Host := []string{
			"127.0.0.1:8000" }
			const (
				Username = "cmpe281"
				Password = "fall"
				Database = "musicPayment_DB"
				Collection = "payment_process"
			)
			session, err := mgo.DialWithInfo(&mgo.DialWithInfo{
				Addrs: Host,

			})
			if err != nil {
				panic(err)
			}

			defer session.Close()

			
	router :=productdb.NewRouter()

	m_origins := handlers.AllowedOrigins([]string{"*"}) 
	m_methods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

 	log.Fatal(http.ListenAndServe(":8000",
  	handlers.CORS(m_origins, m_methods)(router)))
}

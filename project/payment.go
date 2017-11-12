// Source : https://pythonprogramming.net/go/introduction-go-language-programming-tutorial/

package main

import ("fmt" 
        "net/http"
	"html/template")

type NewsAggPage struct {
	Title string
	News string
}
func indexHandler(w http.ResponseWriter, r * http.Request) {
    fmt.Fprintf(w, "<h1> Payment page </h1>")
}

func PaymentHandler(w http.ResponseWriter, r * http.Request) {
    p := NewsAggPage{Title:"Payment Process", News :"Shopping cart"}   
    temp, _ := template.ParseFiles("paymentui.html")
    temp.Execute(w, p)
}

func main() {
   fmt.Println("Payment")
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/payment", PaymentHandler)
    http.ListenAndServe(":8000", nil)
}

// Reference : https://pythonprogramming.net/go/introduction-go-language-programming-tutorial/

package main

import ("fmt" 
        "net/http"
	"io/ioutil"
	"html/templete"
	"encoding/xml")

type Sitemapindex struct{
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Location []string `xml:"url>loc"` 
}

type NewsMap struct {
	Keyword string
	Location string
}

type NewsPaymentPage struct {
	Title string
	News string
}

func indexHandler(w http.ResponseWriter, r * http.Request) {
    fmt.Fprintf(w, "<h1> Payment page </h1>")
}

func PaymentHandler(w http.ResponseWriter, r * http.Request) {
    var s Sitemapindex
    var n News
    resp _ := http.Get("https://github.com/nguyensjsu/cmpe281-sky/tree/master")
    bytes, _ := ioutil.ReadAll(resp.Body)
    xml.Unmarshal(bytes, &s)
    news_map := make(map[string]NewsMap)


    p := NewsPaymentPage{Title:"Payment Process", News :"Shopping cart"}   
    temp, _ := template.ParseFiles("paymentui.html")
    temp.Execute(w, p)
}

func main() {
   fmt.Println("Payment")
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/payment", PaymentHandler)
    http.ListenAndServe(":8000", nil)
}

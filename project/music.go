package main

import (

   hq "github.com/nguyensjsu/cmpe281-sky/blob/project/main.go"
   pay "github.com/nguyensjsu/cmpe281-sky/blob/project/payment.go"

)

const storePage = `
<h1>Store</h1>
<hr>
<form method="post" action="/payment">
   <button type="submit">Checkout</button>
</form>   
`

func storePageHandler(response http.ResponseWriter, request *http.Request) {
   fmt.Fprintf(response, storePage)
}

var route = mux.NewRouter()

func main() {

   route.HandleFunc("/store", storePageHandler)
   route.HandleFunc("/payment", pay)

}

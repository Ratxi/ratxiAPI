package main

import (
	"bytes"
	"github.com/gorilla/mux"
	"io/ioutil"
	//	"encoding/json"
	"log"
	"net/http"
)

// Post Paypal webhook
func PaypalHook(w http.ResponseWriter, r *http.Request) {
	buf, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		log.Print("bodyErr ", bodyErr.Error())
		http.Error(w, bodyErr.Error(), http.StatusInternalServerError)

	}

	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	log.Printf("BODY: %q", rdr1)

	log.Println(r.URL)
}

//  Main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/payratxi/webhook/paypal", PaypalHook).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
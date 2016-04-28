package guesthandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kcraybould/guestdata"
)

func GuestGetSearchLnameHandler(w http.ResponseWriter, r *http.Request) {
	lastName := r.URL.Query().Get("lastName")
	if guests, ok := guestdata.ReturnGuestsSearch(lastName); ok {
		outJson, error := json.Marshal(guests)
		if error != nil {
			log.Println(error.Error())
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, string(outJson))
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Guests not found")
	}

}

func GuestGetSearchNameHandler(w http.ResponseWriter, r *http.Request) {
	lastName := r.URL.Query().Get("lastName")
	firstName := r.URL.Query().Get("firstName")
	if guests, ok := guestdata.ReturnGuestsSearch(lastName, firstName); ok {
		outJson, error := json.Marshal(guests)
		if error != nil {
			log.Println(error.Error())
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, string(outJson))
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Guests not found")
	}

}

func GuestGetSearchEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("emailAddress")
	if guests, ok := guestdata.ReturnGuestEmailSearch(email); ok {
		outJson, error := json.Marshal(guests)
		if error != nil {
			log.Println(error.Error())
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, string(outJson))
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Guests not found")
	}

}

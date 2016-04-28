package guesthandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kcraybould/guestdata"
)

func GuestGetListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if guests, ok := guestdata.ReturnGuestsView(); ok {
		outJson, error := json.Marshal(guests)
		if error != nil {
			log.Println(error.Error())
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, string(outJson))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

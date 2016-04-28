package guesthandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kcraybould/guestdata"
)

func GuestGetByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//this gets the input form the web service, spefically thre URI
	vars := mux.Vars(r)
	guestId := vars["guestId"]

	log.Println("Request for:", guestId)

	if guest, ok := guestdata.ReturnGuestsById(guestId); ok {
		outJson, error := json.Marshal(guest)
		if error != nil {
			log.Println(error.Error())
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, string(outJson))
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Guest not found")
	}
}

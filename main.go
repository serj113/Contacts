package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/serj113/contacts/config"
	. "github.com/serj113/contacts/dao"
	. "github.com/serj113/contacts/db"
	. "github.com/serj113/contacts/model"
)

var config = Config{}
var dao = ContactsDao{}

func AllContactsEndPoint(w http.ResponseWriter, r *http.Request) {
	var contact Contact
	var arr_contact []Contact
	var response Response

	db := Connect()
	defer db.Close()

	rows, err := db.Query("Select * from contact")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone, &contact.Email); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_contact = append(arr_contact, contact)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_contact

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/contacts", AllContactsEndPoint).Methods("GET")
	r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")
	r.HandleFunc("/", AllContactsEndPoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"encoding/json"
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

func CreateNewContact(w http.ResponseWriter, r *http.Request) {
	var response Response
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
	}
	if r.Form.Get("name") != "" && r.Form.Get("phone") != "" {
		name := r.Form.Get("name")
		phone := r.Form.Get("phone")
		email := ""
		if r.Form.Get("email") != "" {
			email = r.Form.Get("email")
		}
		db := Connect()
		defer db.Close()

		_, err := db.Query("insert into contact (name, phone, email) value (?, ?, ?)", name, phone, email)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		} else {
			response.Status = 1
			response.Message = "Success"
		}
	} else {
		response.Status = 1
		response.Message = "invalid params"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var contact Contact
	var arr_contact []Contact
	var response Response

	db := Connect()
	defer db.Close()

	rows, err := db.Query("Select * from contact where id=?", params["id"])
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

func UpdateContact(w http.ResponseWriter, r *http.Request) {

}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Contact ID")
		return
	}
	respondWithJson(w, http.StatusOK, movie)
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
	r.HandleFunc("/contacts", CreateNewContact).Methods("POST")
	r.HandleFunc("/contacts/{id}", GetContact).Methods("GET")
	r.HandleFunc("/contacts/{id}", UpdateContact).Methods("PUT")
	r.HandleFunc("/contacts/{id}", DeleteContact).Methods("DELETE")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

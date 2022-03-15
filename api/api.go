package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ErrorHandler(err error, msg string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(msg, err)))
}

func GetAllValues(w http.ResponseWriter, req *http.Request) {
	response, err := GetAll()
	if err != nil {
		ErrorHandler(err, "Failed to get items: %v", w)
		return
	}

	enc, err := json.Marshal(response)
	if err != nil {
		ErrorHandler(err, "Failed to get marshal data: %v", w)
		return
	}
	w.Write(enc)
}

func GetOne(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	response, err := Get(id)
	if err != nil {
		ErrorHandler(err, "Failed to get data: %v", w)
		return
	}

	enc, err := json.Marshal(response)
	if err != nil {
		ErrorHandler(err, "Failed to get data: %v", w)
		return
	}
	w.Write(enc)
}

func CreateValue(w http.ResponseWriter, req *http.Request) {
	Name := req.FormValue("name")
	Surname := req.FormValue("surname")
	Age := req.FormValue("age")
	val, err := strconv.Atoi(Age)
	if err != nil {
		ErrorHandler(err, "Failed to parse data: %v", w)
	}
	item := Value{Name: Name, Surname: Surname, Age: val}

	if err := Create(item); err != nil {
		ErrorHandler(err, "Failed to save data: %v", w)
		return
	}
	w.Write([]byte("Success create"))
}

func DeleteById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	if err := Delete(id); err != nil {
		ErrorHandler(err, "Failed to delete data: %v", w)
		return
	}

	w.Write([]byte("Success delete"))
}

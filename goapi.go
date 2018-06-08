package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type Person struct {
	ID string `json:"id,omitempty"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Address *Address `json:"address"`
}

type Address struct {
	City string `json:"city"`
	State string `json:"state"`
}

var people []Person

func GetPerson(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/people/get/"):]
	w.Header().Set("Content-Type", "application/json")

	for _, item := range people {
		if item.ID == id {
			fmt.Println(item.ID)
			fmt.Println(id)
			data, _ := json.Marshal(item)
			w.Write(data)
			return
		}
	}

	data := make(map[string]interface{})
	data["msg"] = "Not found"
	jsonOut, _ := json.Marshal(data)
	w.Write(jsonOut)
}


func GetPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}


func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	fmt.Println(person)
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}


func DeletePerson(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/people/delete/"):]
	w.Header().Set("Content-Type", "application/json")

	for i, item := range people {
		if item.ID == id {
			people = append(people[:i], people[i+1:]...)
			json.NewEncoder(w).Encode(people)
			return
		}
	}

	data := make(map[string]interface{})
	data["msg"] = "Not found"
	jsonOut, _ := json.Marshal(data)
	w.Write(jsonOut)
}

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/people", GetPeople)
	r.HandleFunc("/people/get/", GetPerson)
	r.HandleFunc("/people/create", CreatePerson)
	r.HandleFunc("/people/delete/", DeletePerson)

	http.ListenAndServe(":8075", r)
}

func init() {
	people = append(people, Person{ID:"1", Firstname:"Cyan", Lastname:"Tarek", Address:&Address{City:"Gazipur", State:"Dhaka"}})
	people = append(people, Person{ID:"2", Firstname:"Amilin", Lastname:"Diazo"})
}
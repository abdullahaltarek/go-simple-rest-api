package main

//important packages
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//data structure to hold person data
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Address   *Address `json:"address"`
}

//data structure to hold address data
type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

var people []Person

func JsonResponse(m string) []byte {
	data := make(map[string]interface{})
	data["msg"] = m
	jsonOut, _ := json.Marshal(data)
	return jsonOut
}

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(time.Now(), r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		w.Write(JsonResponse("method not allowed"))
		return
	}

	id := r.URL.Path[len("/people/get/"):]

	for _, item := range people {
		if item.ID == id {
			data, _ := json.Marshal(item)
			w.Write(data)
			return
		}
	}
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.Write(JsonResponse("method not allowed"))
		return
	}

	json.NewEncoder(w).Encode(people)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		w.Write(JsonResponse("method not allowed"))
		return
	}

	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	fmt.Println(person)
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "PUT" {
		w.Write(JsonResponse("method not allowed"))
		return
	}

	id := r.URL.Path[len("/people/update/"):]

	var personData Person

	json.NewDecoder(r.Body).Decode(&personData)

	for i, item := range people {
		if item.ID == id {
			if len(personData.Firstname) > 0 {
				people[i].Firstname = personData.Firstname
			}

			if len(personData.Lastname) > 0 {
				people[i].Lastname = personData.Lastname
			}

			json.NewEncoder(w).Encode(people[i])
			return
		}
	}
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "DELETE" {
		w.Write(JsonResponse("method not allowed"))
		return
	}

	id := r.URL.Path[len("/people/delete/"):]

	for i, item := range people {
		if item.ID == id {
			people = append(people[:i], people[i+1:]...)
			json.NewEncoder(w).Encode(people)
			return
		}
	}

	w.Write(JsonResponse("not found"))
}

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/people", GetPeople)
	r.HandleFunc("/people/get/", GetPerson)
	r.HandleFunc("/people/create", CreatePerson)
	r.HandleFunc("/people/delete/", DeletePerson)
	r.HandleFunc("/people/update/", UpdatePerson)

	http.ListenAndServe(":8075", Logger(r))
}

func init() {
	people = append(people, Person{ID: "1", Firstname: "Cyan", Lastname: "Tarek", Address: &Address{City: "Gazipur", State: "Dhaka"}})
	people = append(people, Person{ID: "2", Firstname: "Amilin", Lastname: "Diazo"})
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

var People = []Person{
	{Name: "abc", Id: "2"},
	{Name: "def", Id: "1"},
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page") // sending text response
	fmt.Println("Endpoint Hit: homePage")
}

func sendAllPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(People) // sending Json response
}

func sendPeopleById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	for _, article := range People {
		if article.Id == id {
			fmt.Println(id)
			json.NewEncoder(w).Encode(article) // sending Json response
		}
	}
}

func addNewPerson(w http.ResponseWriter, req *http.Request) {
	req_body, _ := ioutil.ReadAll(req.Body)
	var new_person Person
	json.Unmarshal(req_body, &new_person) // converting Json data to struct
	People = append(People, new_person)
	json.NewEncoder(w).Encode(People) // sending json response back to client
}

func handleRequests() {

	Router := mux.NewRouter().StrictSlash(true)

	Router.HandleFunc("/", homePage)
	Router.HandleFunc("/people", sendAllPeople)
	Router.HandleFunc("/people/{id}", sendPeopleById)
	Router.HandleFunc("/person", addNewPerson).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", Router)) // app is running on port 3000
}

func main() {
	handleRequests()
}

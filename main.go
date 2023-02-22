
package main

import(
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
);

type Person struct {
	ID string `json:id,omitempty"`
	FirstName string `json:firstname,omitempty"`
	LastName string `json:lastname,omitempty"`
	Address *Address `json:address,omitempty"`
};

type Address struct {
	City string `json:city,omitempty"`
	State string `json:state,omitempty"`
}


var people []Person;

//Get a list of people
func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request){
	json.NewEncoder(w).Encode(people)
}
//Get One Person from list
func GetPersonEndPoint(w http.ResponseWriter, req *http.Request){

	//Obtain the value for get in the request
	params :=mux.Vars(req)
		for _, item :=range people{
			if item.ID == params["id"] {
				//return find person by ID
				json.NewEncoder(w).Encode(item)
				return
			}
		}
		//Not find return nothing
		json.NewEncoder(w).Encode(&Person{})
}

func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request){

		//Obtain the values for post from the request to params
		params :=mux.Vars(req)
		//Defined the new person
		var person Person
		//Decoding the json
		_=json.NewDecoder(req.Body).Decode(&person)
		//Adding ID
		person.ID = params["id"]
		//Adding to var person 
		people=append(people, person)
		//Send the request
		json.NewEncoder(w).Encode(people)


}
func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request){

	params:= mux.Vars(req)
	for index, item := range people{
		if item.ID == params["id"] {
			people =append(people[:index], people[index + 1:]... )
		//For outgoing for
			break
		}
	}
	//Return all if find or not
	json.NewEncoder(w).Encode(people)
}


func main(){

	//Routes
	router :=mux.NewRouter();

	//insert People to Arr people to demonstrate
	people = append(people, Person{ID:"1", FirstName:"Bruce", LastName:"Wayne", Address:&Address{City: "Gotham", State: "New Jersey"}})
	people = append(people, Person{ID:"2", FirstName:"Dead", LastName:"Pool"})
	
	//EndPoints

	router.HandleFunc("/people", GetPeopleEndPoint) .Methods("GET")
	//	URL to test GET ALL  	//http://localhost:3000/people
	
	router.HandleFunc("/people/{id}", GetPersonEndPoint) .Methods("GET")
	//	URL to test GET ONE Person 	//http://localhost:3000/people/ID here ID=1 or ID=2
	
	router.HandleFunc("/people/{id}", CreatePersonEndPoint) .Methods("POST")
	//	URL to test POST Person 	//http://localhost:3000/people/3

/*
1.- Add in the Header Content-Type : application/json
2.- The numer of ID the url (http://localhost:3000/people/3)
3.- The raw body type json:

{
	"firstname":"Nombre",
	"lastname":"Apellido"
}

*/

	router.HandleFunc("/people/{id}", DeletePersonEndPoint) .Methods("DELETE")
	//	URL to test DELETE Person 	//http://localhost:3000/people/3


	//Server LIsterner Methods

	//Error Handler
	log.Fatal(http.ListenAndServe(":3000", router))

}
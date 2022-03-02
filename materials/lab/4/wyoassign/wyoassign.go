package wyoassign

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Response struct {
	Assignments []Assignment `json:"assignments"`
}

type Assignment struct {
	Id          string `json:"id"`
	Title       string `json:"title`
	Description string `json:"desc"`
	Points      int    `json:"points"`
}

func unmarshalAssignment(r *http.Request) (Assignment, error) {
	var assignment Assignment

	if r.FormValue("id") == "" {
		return assignment, fmt.Errorf("'id' must be provided as a non-empty field")
	}
	assignment.Id = r.FormValue("id")

	if r.FormValue("title") == "" {
		return assignment, fmt.Errorf("'title' must be provided as a non-empty field")
	}
	assignment.Title = r.FormValue("title")

	if r.FormValue("desc") == "" {
		return assignment, fmt.Errorf("'desc' must be provided as a non-empty field")
	}
	assignment.Description = r.FormValue("desc")

	if r.FormValue("points") == "" {
		return assignment, fmt.Errorf("'points' must be provided as a non-empty field")
	}
	assignment.Points, _ = strconv.Atoi(r.FormValue("points"))

	return assignment, nil
}

var Assignments []Assignment

const Valkey string = "FooKey"

func InitAssignments() {
	var assignmnet Assignment
	assignmnet.Id = "Mike1A"
	assignmnet.Title = "Lab 4 "
	assignmnet.Description = "Some lab this guy made yesteday?"
	assignmnet.Points = 20
	Assignments = append(Assignments, assignmnet)
}

func APISTATUS(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func GetAssignments(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response Response

	response.Assignments = Assignments

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	//TODO
	w.Write(jsonResponse)
}

func GetAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	for _, assignment := range Assignments {
		if assignment.Id == params["id"] {
			json.NewEncoder(w).Encode(assignment)
			break
		}
	}

	w.WriteHeader(http.StatusBadRequest)
}

func DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, assignment := range Assignments {
		if assignment.Id == params["id"] {
			Assignments = append(Assignments[:index], Assignments[index+1:]...)
			response["status"] = "Success"
			break
		}
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")

	var response Response
	response.Assignments = Assignments

	r.ParseForm()
	id := r.FormValue("id")
	if id != "" {
		for i, assignment := range Assignments {
			if assignment.Id == id {
				if r.FormValue("title") != "" {
					Assignments[i].Title = r.FormValue("title")
				}
				if r.FormValue("desc") != "" {
					Assignments[i].Description = r.FormValue("desc")
				}
				if r.FormValue("points") != "" {
					Assignments[i].Points, _ = strconv.Atoi(r.FormValue("points"))
				}
				json.NewEncoder(w).Encode(Assignments[i])
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound)

}

func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()

	assignmnet, err := unmarshalAssignment(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotAcceptable)
	}

	Assignments = append(Assignments, assignmnet)
	w.WriteHeader(http.StatusCreated)
}

// Package LunchMenu serves up static HTML pages and responds to REST requests
// to enable lunch menus to be stored, updated, queried, and to enable homeroom
// teachers to specify how many of each item are requested by their students.
package LunchMenu

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

// Menu represents the menu items (as a JSON list) for a specific date.
type Menu struct {
	Date  string
	Items string
}

// MenuDto represents the menu items (as a list) for the specific date.
type MenuDto struct {
	Date  string
	Items []string
}

// MenuRequest represents the count of each menu item (as a JSON map) requested by a homeroom for a specific date.
type MenuRequest struct {
	Date     string
	Homeroom string
	Choices  string
}

// MenuRequestDto represents the count of each menu item (as a map) requested by a homeroom for a specific date.
type MenuRequestDto struct {
	Date     string
	Homeroom string
	Choices  map[string]int
}

func init() {
	r := mux.NewRouter()

	// Menu CRUD operations
	r.HandleFunc("/menus/{date}", getMenu).Methods("GET")
	r.HandleFunc("/menus/{date}", createMenu).Methods("POST")
	r.HandleFunc("/menus/{date}", deleteMenu).Methods("DELETE")

	// MenuRequest CRUD operations
	r.HandleFunc("/requests/{date}", getMenuRequests).Methods("GET")
	r.HandleFunc("/requests/{date}/{homeroom}", createMenuRequest).Methods("POST")
	r.HandleFunc("/requests/{date}/{homeroom}", deleteMenuRequest).Methods("DELETE")

	// Static content
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static/"))))

	http.Handle("/", r)
}

func getMenu(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// Get date from URL
	params := mux.Vars(r)
	date := params["date"]

	// Query
	menu := new(Menu)
	key := datastore.NewKey(c, "Menu", date, 0, nil)
	err := datastore.Get(c, key, menu)
	if err == datastore.ErrNoSuchEntity {
		log.Infof(c, "No menu found for [%s]", date)
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert to DTO
	dto := new(MenuDto)
	dto.Date = date
	dto.Items = make([]string, 0)
	err = json.Unmarshal([]byte(menu.Items), &dto.Items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto)
}

func createMenu(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// Get date from URL
	params := mux.Vars(r)
	date := params["date"]

	// Get items from Body
	var items []string
	err := json.NewDecoder(r.Body).Decode(&items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save Menu
	menu := new(Menu)
	menu.Date = date
	bytes, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	menu.Items = string(bytes)
	key := datastore.NewKey(c, "Menu", date, 0, nil)
	key, err = datastore.Put(c, key, menu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteMenu(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// Get date from URL
	params := mux.Vars(r)
	date := params["date"]

	log.Infof(c, "Deleting menu for [%s]", date)

	// Delete Menu
	key := datastore.NewKey(c, "Menu", date, 0, nil)
	err := datastore.Delete(c, key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getMenuRequests(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// Get date from URL
	params := mux.Vars(r)
	date := params["date"]

	// Query MenuRequests for this date
	q := datastore.NewQuery("MenuRequest").Filter("Date =", date)
	var requests []MenuRequest
	_, err := q.GetAll(c, &requests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert to DTOs
	dtos := make([]MenuRequestDto, len(requests))
	for i, request := range requests {
		dto := new(MenuRequestDto)
		dto.Date = request.Date
		dto.Homeroom = request.Homeroom
		dto.Choices = make(map[string]int)
		err := json.Unmarshal([]byte(request.Choices), &dto.Choices)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		dtos[i] = *dto
	}

	// Return as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dtos)
}

func createMenuRequest(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// Get keys from URL
	params := mux.Vars(r)
	date := params["date"]
	homeroom := params["homeroom"]

	// Get choices from Body
	var choices map[string]int
	err := json.NewDecoder(r.Body).Decode(&choices)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(choices)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create MenuRequest
	request := new(MenuRequest)
	request.Date = date
	request.Homeroom = homeroom
	request.Choices = string(bytes)

	log.Infof(c, "Saving MenuRequest for [%s] and [%s] : %s", date, homeroom, request.Choices)

	// Store MenuRequest
	id := date + "|" + homeroom
	key := datastore.NewKey(c, "MenuRequest", id, 0, nil)
	key, err = datastore.Put(c, key, request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteMenuRequest(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// Get keys from URL
	params := mux.Vars(r)
	date := params["date"]
	homeroom := params["homeroom"]

	log.Infof(c, "Deleting MenuRequest for [%s] and [%s]", date, homeroom)

	// Delete MenuRequest
	id := date + "|" + homeroom
	key := datastore.NewKey(c, "MenuRequest", id, 0, nil)
	err := datastore.Delete(c, key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

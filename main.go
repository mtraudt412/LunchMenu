package hello

import (
  "encoding/json"
  "github.com/gorilla/mux"
	"google.golang.org/appengine"
  "google.golang.org/appengine/datastore"
  "net/http"
)

type Menu struct {
  Date string // Must be formatted as YYYYMMDD (e.g. 20170901)
  Items string // Comma-separated list of menu items (e.g. Pizza, Grilled cheese)
}

func init() {
  r := mux.NewRouter()
  r.HandleFunc("/menus", getMenus).Methods("GET")
  r.HandleFunc("/menus/{date}", getMenu).Methods("GET")
  r.HandleFunc("/menus/{date}", addMenu).Methods("POST")
  r.PathPrefix("/scripts/").Handler(http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts"))))
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
  r.Path("/").Handler(http.FileServer(http.Dir("static")))
  http.Handle("/", r)
}

func getMenus(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  q := datastore.NewQuery("Menu").Order("-Date")
  var menus []Menu
  _, err := q.GetAll(c, &menus)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  json.NewEncoder(w).Encode(menus)
}

func getMenu(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  menu := new(Menu)
  params := mux.Vars(r)
  date := params["date"]
  key := datastore.NewKey(c, "Menu", date, 0, nil)
  err := datastore.Get(c, key, menu)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  err = json.NewEncoder(w).Encode(menu)
  if err != nil {
    panic(err)
  }
}

func addMenu(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  menu := new(Menu)
  err := json.NewDecoder(r.Body).Decode(menu)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  params := mux.Vars(r)
  date := params["date"]
  if date != menu.Date {
    http.Error(w, "Invalid date", http.StatusUnprocessableEntity)
  }
  key := datastore.NewKey(c, "Menu", menu.Date, 0, nil)
  key, err = datastore.Put(c, key, menu)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}


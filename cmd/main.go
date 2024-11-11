package main

import (
	"encoding/json"
	"github.com/bourd0n/go-interview-task/pkg/models"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Register the routes and handlers
	mux.Handle("/", &HomeHandlerDefaultImplementation{})
	mux.Handle("/create-country", &HomeHandlerDefaultImplementation{})
	mux.Handle("/get-country-by-id", &HomeHandlerDefaultImplementation{})
	mux.Handle("/set-town-population", &HomeHandlerDefaultImplementation{})
	mux.Handle("/calculate", &HomeHandlerDefaultImplementation{})

	// Run the server
	http.ListenAndServe(":8080", mux)
}

type HomeHandlerDefaultImplementation struct{}

func (h *HomeHandlerDefaultImplementation) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case ((r.Method == http.MethodPost) && (r.URL.Path == "create-country")):
		CreateCountry(w, r)
		return
	case (r.Method == http.MethodPost && (r.URL.Path == "get-country-by-id")):
		GetCountryByName(w, r)
		return
	case (r.Method == http.MethodPost && (r.URL.Path == "set-town-population")):
		SetTownPopulation(w, r)
		return
	default:
		w.Write([]byte("This is my home page"))
		return
	}
}

func CreateCountry(w http.ResponseWriter, r *http.Request) {
	var country models.CountryModel

	json.NewDecoder(r.Body).Decode(&country)

	models.CreateCountry(&country)

	w.WriteHeader(http.StatusOK)
}

func GetCountryByName(w http.ResponseWriter, r *http.Request) {
	var request models.CountryByIdRequest

	json.NewDecoder(r.Body).Decode(&request)

	models.GetCountry(request.Name)

	w.WriteHeader(http.StatusOK)
}

func SetTownPopulation(w http.ResponseWriter, r *http.Request) {
	var request models.Town

	json.NewDecoder(r.Body).Decode(&request)

	models.SetTownPopulation(&request)

	w.WriteHeader(http.StatusOK)
}

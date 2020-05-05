package handlers

import (
	"golang-fifa-world-cup-web-service/data"
	"net/http"
)

// RootHandler returns an empty body status code
func RootHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNoContent)
}

// ListWinners returns winners from the list
func ListWinners(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var winners []byte
	var err error

	year := req.URL.Query().Get("year")
	if year == "" {
		winners, err = data.ListAllJSON()
	} else {
		winners, err = data.ListAllByYear(year)
	}

	if err != nil {
		if year != "" {
			res.WriteHeader(http.StatusBadRequest)
		} else {
			res.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	res.Write(winners)
}

// AddNewWinner adds new winner to the list
func AddNewWinner(res http.ResponseWriter, req *http.Request) {

}

// WinnersHandler is the dispatcher for all /winners URL
func WinnersHandler(res http.ResponseWriter, req *http.Request) {

}

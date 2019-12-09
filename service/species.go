package service

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/XinHeLianSheng/server/database"
)

func getSpeciesById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	data := database.GetValue([]byte("species"), []byte(vars["id"]))
	if len(data) <= 10 {
		w.Write([]byte("404 Not Found"))
	} else {
		w.Write([]byte(data))

	}
}

func speciesPagesHandler(w http.ResponseWriter, req *http.Request) {
	data := database.GetBucketCount([]byte("species"))
	w.Write([]byte(strconv.Itoa(data)))
}

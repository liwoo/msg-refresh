package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"msgRefresh3/refresh"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/entity-refreshes", handleEntityRefreshes).Methods("POST")
	log.Fatal(http.ListenAndServe(":7007", router))
}

func handleEntityRefreshes(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var entity refresh.RefreshEntityDto

	marshalError := json.Unmarshal(body, &entity)
	if err != nil {
		log.Fatalf(marshalError.Error())
	}
	refresher := refresh.EntityRefresher{&entity}
	refreshErr := refresher.Refresh()
	if refreshErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(refreshErr.Error()))
	}
	w.WriteHeader(http.StatusOK)
}

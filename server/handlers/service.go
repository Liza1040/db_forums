package handlers

import (
	. "db-forums/server/db_requests"
	"encoding/json"
	"net/http"
)

func Cleardb(response http.ResponseWriter, request *http.Request) {
	str, err := Cleardb_all()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(toMessage(str + ". Error: " + err.Error()))
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write(toMessage(str))
}

func Statusdb(response http.ResponseWriter, request *http.Request) {
	stats, err := GetStats()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(toMessage("Invalid db request"))
		return
	}

	body, err := json.Marshal(stats)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(toMessage("Can't marshal JSON file"))
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write(body)
}

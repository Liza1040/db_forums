package handlers

import (
	. "DB_forums/server/DB-requests"
	"encoding/json"
	"net/http"
)

func ClearDB(response http.ResponseWriter, request *http.Request) {
	str, err := ClearDB_all()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(toMessage(str + ". Error: " + err.Error()))
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write(toMessage(str))
}

func StatusDB(response http.ResponseWriter, request *http.Request) {
	stats, err := GetStats()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(toMessage("Invalid DB request"))
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

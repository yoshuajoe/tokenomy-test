package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Data struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

var dummyData = []Data{
	{Id: 1, Name: "A"},
	{Id: 2, Name: "B"},
	{Id: 3, Name: "C"},
}

func main() {
	http.HandleFunc("/", getData)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func getData(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		sendResponse(w, http.StatusOK, "OK", dummyData)
		return
	}

	ids := strings.Split(id, ",")
	var data []Data

	for _, idStr := range ids {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			sendResponse(w, http.StatusBadRequest, fmt.Sprintf("invalid or empty ID: \"%s\"", strings.Join(ids, "")), nil)
			return
		}
		found := false
		for _, d := range dummyData {
			if d.Id == id {
				data = append(data, d)
				found = true
				break
			}
		}
		if !found && len(ids) == 1 {
			sendResponse(w, http.StatusNotFound, fmt.Sprintf("resource with ID %s not exist", strings.Join(ids, "")), nil)
			return
		}
	}
	sendResponse(w, http.StatusOK, "OK", data)
}

func sendResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	response := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	responseJson, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"code":500,"message":"internal server error"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(responseJson)
}

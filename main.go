package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Activity struct {
	Email string `json:"email"`
	Title string `json:"title"`
}

type response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Kosong struct {
}

var kosong Kosong

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/activity-groups", ActivityRest)
	http.ListenAndServe(":8000", nil)
}

func ActivityRest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var t Activity
		err := decoder.Decode(&t)

		if t.Title == "" {
			var resp response
			resp.Status = "Bad Request"
			resp.Message = "title cannot be null"
			resp.Data = kosong
			w.WriteHeader(http.StatusBadRequest)

			jData, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}
			w.Write(jData)
			return
		}

		if err != nil {
			fmt.Fprint(w, "Internal Server Error")
		} else {
			fmt.Fprint(w, "Internal Server Error")
		}
	case "GET":
		w.Write([]byte("get"))
	default:
		http.Error(w, "", http.StatusBadRequest)
	}

}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	jData, err := json.Marshal(names)
	if err != nil {
		fmt.Fprint(w, "Internal Server Error")
	} else {
		w.Write(jData)
	}

}

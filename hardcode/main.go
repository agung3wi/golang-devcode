package main

import (
	"net/http"
)

type Result struct {
	Body   string
	Status int
}

var statusCode int
var bodyString string
var result = []Result{
	{Body: `{"status":"Success","message":"Success","data":{"created_at":null,"updated_at":null,"id":1,"title":"testing112","email":"e9c4786a-66a3-48e4-9478-320fc834d1f1@test.com"}}`, Status: 201},
	{Body: `{"status":"Bad Request","message":"title cannot be null","data":{}}`, Status: 400},
	{Body: `{"status":"Success","message":"Success","data":{"id":1,"email":"e9c4786a-66a3-48e4-9478-320fc834d1f1@test.com","title":"testing112Updated","created_at":null,"updated_at":null,"deleted_at":null}}`, Status: 200},
	{Body: `{"status":"Not Found","message":"Activity with ID 999999999 Not Found","data":{}}`, Status: 404},
	{Body: `{"status":"Success","message":"Success","data":{"id":1,"email":"e9c4786a-66a3-48e4-9478-320fc834d1f1@test.com","title":"testing112Updated","created_at":null,"updated_at":null,"deleted_at":null}}`, Status: 200},
	{Body: `{"status":"Not Found","message":"Activity with ID 999999999 Not Found","data":{}}`, Status: 404},
	{Body: `{"status":"Success","message":"Success","data":[{"id":1,"email":"e9c4786a-66a3-48e4-9478-320fc834d1f1@test.com","title":"testing112Updated","created_at":null,"updated_at":null,"deleted_at":null}]}`, Status: 200},
	{Body: `{"status":"Success","message":"Success","data":{"created_at":null,"updated_at":null,"id":1,"title":"todoTesting","activity_group_id":1,"is_active":true,"priority":"very-high"}}`, Status: 201},
	{Body: `{"status":"Bad Request","message":"title cannot be null","data":{}}`, Status: 400},
	{Body: `{"status":"Bad Request","message":"activity_group_id cannot be null","data":{}}`, Status: 400},
	{Body: `{"status":"Success","message":"Success","data":{"id":1,"activity_group_id":"1","title":"todoTestingUpdated","is_active":"1","priority":"very-high","created_at":null,"updated_at":null,"deleted_at":null}}`, Status: 200},
	{Body: `{"status":"Success","message":"Success","data":{"id":1,"activity_group_id":"1","title":"todoTestingUpdated","is_active":false,"priority":"very-high","created_at":null,"updated_at":null,"deleted_at":null}}`, Status: 200},
	{Body: `{"status":"Not Found","message":"Todo with ID 999999999 Not Found","data":{}}`, Status: 404},
	{Body: `{"status":"Success","message":"Success","data":{"id":1,"activity_group_id":"1","title":"todoTestingUpdated","is_active":false,"priority":"very-high","created_at":null,"updated_at":null,"deleted_at":null}}`, Status: 200},
	{Body: `{"status":"Not Found","message":"Todo with ID 999999999 Not Found","data":{}}`, Status: 404},
	{Body: `{"status":"Success","message":"Success","data":[{"id":1,"activity_group_id":"1","title":"todoTestingUpdated","is_active":false,"priority":"very-high","created_at":null,"updated_at":null,"deleted_at":null}]}`, Status: 200},
	{Body: `{"status":"Success","message":"Success","data":[]}`, Status: 200},
	{Body: `{"status":"Success","message":"Success","data":{}}`, Status: 200},
	{Body: `{"status":"Not Found","message":"Todo with ID 999999999 Not Found","data":{}}`, Status: 404},
	{Body: `{"status":"Success","message":"Success","data":{}}`, Status: 200},
	{Body: `{"status":"Not Found","message":"Activity with ID 999999999 Not Found","data":{}}`, Status: 404},
	{Body: `{"status":"Success","message":"Success","data":{"created_at":null,"updated_at":null,"id":2,"title":"performanceTesting","email":"performance@test.com"}}`, Status: 201},
}

var req int

func main() {
	req = 1
	statusCode = result[0].Status
	bodyString = result[0].Body
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":3030", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {

	if req > 1021 {
		w.WriteHeader(200)
		w.Write([]byte(""))
		return
	} else if req > 21 {
		w.WriteHeader(201)
		w.Write([]byte(""))
		req = req + 1
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(bodyString))

	go func() {
		statusCode = result[req].Status
		bodyString = result[req].Body
		req = req + 1
	}()

	// if len(result) == 0 {
	// 	if r.Method == "POST" {

	// 		w.WriteHeader(201)
	// 	}
	// 	w.Write([]byte(""))
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(result[0].Status)
	// w.Write([]byte(result[0].Body))
	// go func() {
	// 	if len(result) > 1 {
	// 		result = result[1:]
	// 	} else {
	// 		result = []Result{}
	// 	}
	// }()

}

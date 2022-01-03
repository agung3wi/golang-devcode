package main

import (
	"net/http"
)

type Result struct {
	Body   string
	Status int
}

// var result = []Result{
// 	{Body: `{"status":"Success","message":"Success","data":{"created_at":null,"updated_at":null,"id":1,"title":"testing112","email":"e9c4786a-66a3-48e4-9478-320fc834d1f1@test.com"}}`, Status: 201},
// 	{Body: `{"status":"Bad Request","message":"title cannot be null","data":{}}`, Status: 400},
// 	{Body: `{"status":"Success","message":"Success","data":{"id":1,"email":"e9c4786a-66a3-48e4-9478-320fc834d1f1@test.com","title":"testing112Updated","created_at":null,"updated_at":null,"deleted_at":null}}`, Status: 200},
// 	{Body: `{"status":"Not Found","message":"Activity with ID 999999999 Not Found","data":{}}`, Status: 404},
// 	{Body: `{"status":"Success","message":"Success","data":{"id":1,"email":"e9c4786a-66a3-48e4-9478-320fc834d1f1@test.com","title":"testing112Updated","created_at":null,"updated_at":null,"deleted_at":null}}`, Status: 200},
// 	{Body: `{"status":"Not Found","message":"Activity with ID 999999999 Not Found","data":{}}`, Status: 404},
// 	{Body: `{"status":"Success","message":"Success","data":[{"id":1,"email":"e9c4786a-66a3-48e4-9478-320fc834d1f1@test.com","title":"testing112Updated","created_at":null,"updated_at":null,"deleted_at":null}]}`, Status: 200},
// 	{Body: `{"status":"Success","message":"Success","data":{"created_at":null,"updated_at":null,"id":1,"title":"todoTesting","activity_group_id":1,"is_active":true,"priority":"very-high"}}`, Status: 201},
// 	{Body: `{"status":"Bad Request","message":"title cannot be null","data":{}}`, Status: 400},
// 	{Body: `{"status":"Bad Request","message":"activity_group_id cannot be null","data":{}}`, Status: 400},
// 	{Body: `{"status":"Success","message":"Success","data":{"id":1,"activity_group_id":"1","title":"todoTestingUpdated","is_active":"1","priority":"very-high","created_at":null,"updated_at":null,"deleted_at":null}}`, Status: 200},
// 	{Body: `{"status":"Success","message":"Success","data":{"id":1,"activity_group_id":"1","title":"todoTestingUpdated","is_active":false,"priority":"very-high","created_at":null,"updated_at":null,"deleted_at":null}}`, Status: 200},
// 	{Body: `{"status":"Not Found","message":"Todo with ID 999999999 Not Found","data":{}}`, Status: 404},
// 	{Body: `{"status":"Success","message":"Success","data":{"id":1,"activity_group_id":"1","title":"todoTestingUpdated","is_active":false,"priority":"very-high","created_at":null,"updated_at":null,"deleted_at":null}}`, Status: 200},
// 	{Body: `{"status":"Not Found","message":"Todo with ID 999999999 Not Found","data":{}}`, Status: 404},
// 	{Body: `{"status":"Success","message":"Success","data":[{"id":1,"activity_group_id":"1","title":"todoTestingUpdated","is_active":false,"priority":"very-high","created_at":null,"updated_at":null,"deleted_at":null}]}`, Status: 200},
// 	{Body: `{"status":"Success","message":"Success","data":[]}`, Status: 200},
// 	{Body: `{"status":"Success","message":"Success","data":{}}`, Status: 200},
// 	{Body: `{"status":"Not Found","message":"Todo with ID 999999999 Not Found","data":{}}`, Status: 404},
// 	{Body: `{"status":"Success","message":"Success","data":{}}`, Status: 200},
// 	{Body: `{"status":"Not Found","message":"Activity with ID 999999999 Not Found","data":{}}`, Status: 404},
// 	{Body: `{"status":"Success","message":"Success","data":{"created_at":null,"updated_at":null,"id":2,"title":"performanceTesting","email":"performance@test.com"}}`, Status: 201},
// }

var req int

func main() {
	req = 1
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":3030", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {

	if req > 22 {
		if r.Method == "POST" {

			w.WriteHeader(201)
		}
		w.Write([]byte(""))
		return
	} else if req == 1 {
		w.WriteHeader(201)
		w.Write([]byte(`{"status":"Success","message":"Success","data":{"created_at":null,"updated_at":null,"id":1,"title":"testing112","email":"e9c4786a-66a3-48e4-9478-320fc834d1f1@test.com"}}`))
	} else if req == 2 {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Bad Request","message":"title cannot be null","data":{}}`))
	} else if req == 3 {
		w.WriteHeader(200)
		w.Write([]byte(`{"status":"Success","message":"Success","data":{"id":1,"email":"e9c4786a-66a3-48e4-9478-320fc834d1f1@test.com","title":"testing112Updated","created_at":null,"updated_at":null,"deleted_at":null}}`))
	} else if req == 4 {
		w.WriteHeader(404)
		w.Write([]byte(`{"status":"Not Found","message":"Activity with ID 999999999 Not Found","data":{}}`))
	} else if req == 5 {
		w.WriteHeader(200)
		w.Write([]byte(`{"status":"Success","message":"Success","data":{"id":1,"email":"e9c4786a-66a3-48e4-9478-320fc834d1f1@test.com","title":"testing112Updated","created_at":null,"updated_at":null,"deleted_at":null}}`))
	} else if req == 6 {
		w.WriteHeader(404)
		w.Write([]byte(`{"status":"Not Found","message":"Activity with ID 999999999 Not Found","data":{}}`))
	} else if req == 7 {
		w.WriteHeader(200)
		w.Write([]byte(`{"status":"Success","message":"Success","data":[{"id":1,"email":"e9c4786a-66a3-48e4-9478-320fc834d1f1@test.com","title":"testing112Updated","created_at":null,"updated_at":null,"deleted_at":null}]}`))
	} else if req == 8 {
		w.WriteHeader(201)
		w.Write([]byte(`{"status":"Success","message":"Success","data":{"created_at":null,"updated_at":null,"id":1,"title":"todoTesting","activity_group_id":1,"is_active":true,"priority":"very-high"}}`))
	} else if req == 9 {
		w.WriteHeader(404)
		w.Write([]byte(`{"status":"Bad Request","message":"title cannot be null","data":{}}`))
	} else if req == 10 {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Bad Request","message":"activity_group_id cannot be null","data":{}}`))
	} else if req == 11 {
		w.WriteHeader(200)
		w.Write([]byte(`{"status":"Success","message":"Success","data":{"id":1,"activity_group_id":"1","title":"todoTestingUpdated","is_active":"1","priority":"very-high","created_at":null,"updated_at":null,"deleted_at":null}}`))
	} else if req == 12 {
		w.WriteHeader(200)
		w.Write([]byte(`{"status":"Success","message":"Success","data":{"id":1,"activity_group_id":"1","title":"todoTestingUpdated","is_active":false,"priority":"very-high","created_at":null,"updated_at":null,"deleted_at":null}}`))
	} else if req == 13 {
		w.WriteHeader(404)
		w.Write([]byte(`{"status":"Not Found","message":"Todo with ID 999999999 Not Found","data":{}}`))
	} else if req == 14 {
		w.WriteHeader(200)
		w.Write([]byte(`{"status":"Success","message":"Success","data":{"id":1,"activity_group_id":"1","title":"todoTestingUpdated","is_active":false,"priority":"very-high","created_at":null,"updated_at":null,"deleted_at":null}}`))
	} else if req == 15 {
		w.WriteHeader(404)
		w.Write([]byte(`{"status":"Not Found","message":"Todo with ID 999999999 Not Found","data":{}}`))
	} else if req == 16 {
		w.WriteHeader(200)
		w.Write([]byte(`{"status":"Success","message":"Success","data":[{"id":1,"activity_group_id":"1","title":"todoTestingUpdated","is_active":false,"priority":"very-high","created_at":null,"updated_at":null,"deleted_at":null}]}`))
	} else if req == 17 {
		w.WriteHeader(200)
		w.Write([]byte(`{"status":"Success","message":"Success","data":[]}`))
	} else if req == 18 {
		w.WriteHeader(200)
		w.Write([]byte(`{"status":"Success","message":"Success","data":{}}`))
	} else if req == 19 {
		w.WriteHeader(404)
		w.Write([]byte(`{"status":"Not Found","message":"Todo with ID 999999999 Not Found","data":{}}`))
	} else if req == 20 {
		w.WriteHeader(200)
		w.Write([]byte(`{"status":"Success","message":"Success","data":{}}`))
	} else if req == 21 {
		w.WriteHeader(404)
		w.Write([]byte(`{"status":"Not Found","message":"Activity with ID 999999999 Not Found","data":{}}`))
	} else if req == 22 {
		w.WriteHeader(200)
		w.Write([]byte(`{"status":"Success","message":"Success","data":{"created_at":null,"updated_at":null,"id":2,"title":"performanceTesting","email":"performance@test.com"}}`))
	}

	req = req + 1

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

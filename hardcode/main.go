package main

import (
	"net/http"
)

type Result struct {
	Body   string
	Status int
}

var result = []Result{
	{Body: `{ "status": "Success", "message": "Success", "data": { "created_at": "2021-12-01T09:23:05.825Z", "updated_at": "2021-12-01T09:23:05.826Z", "id": 1, "title": "testing112", "email": "e9c4786a-66a3-48e4-9478-320fc834d1f1@test.com" } }`, Status: 201},
	{Body: `{ "status": "Bad Request", "message": "title cannot be null", "data": {} }`, Status: 400},
	{Body: `{ "status": "Success", "message": "Success", "data": { "id": 1, "email": "e9c4786a-66a3-48e4-9478-320fc834d1f1@test.com", "title": "testing112Updated", "created_at": "2021-11-30T05:29:24.000Z", "updated_at": "2021-11-30T05:29:24.000Z", "deleted_at": null } }`, Status: 200},
	{Body: `{ "status": "Not Found", "message": "Activity with ID 999999999 Not Found", "data": {} }`, Status: 404},
	{Body: `{ "status": "Success", "message": "Success", "data": { "id": 1, "email": "e9c4786a-66a3-48e4-9478-320fc834d1f1@test.com", "title": "testing112Updated", "created_at": "2021-11-30T05:29:24.000Z", "updated_at": "2021-11-30T05:29:24.000Z", "deleted_at": null } }`, Status: 200},
	{Body: `{ "status": "Not Found", "message": "Activity with ID 999999999 Not Found", "data": {} }`, Status: 404},
	{Body: `{ "status": "Success", "message": "Success", "data": [{ "id": 1, "email": "e9c4786a-66a3-48e4-9478-320fc834d1f1@test.com", "title": "testing112Updated", "created_at": "2021-11-30T05:29:24.000Z", "updated_at": "2021-11-30T05:29:24.000Z", "deleted_at": null }] }`, Status: 200},
	{Body: `{ "status": "Success", "message": "Success", "data": { "created_at": "2021-12-09T02:00:00.067Z", "updated_at": "2021-12-09T02:00:00.068Z", "id": 1, "title": "todoTesting", "activity_group_id": 1, "is_active": true, "priority": "very-high" } }`, Status: 201},
	{Body: `{ "status": "Bad Request", "message": "title cannot be null", "data": {} }`, Status: 400},
	{Body: `{ "status": "Bad Request", "message": "activity_group_id cannot be null", "data": {} }`, Status: 400},
	{Body: `{ "status": "Success", "message": "Success", "data": { "id": 1, "activity_group_id": "1", "title": "todoTestingUpdated", "is_active": "1", "priority": "very-high", "created_at": "2021-11-30T05:29:24.000Z", "updated_at": "2021-12-09T02:17:15.758Z", "deleted_at": null } }`, Status: 200},
	{Body: `{ "status": "Success", "message": "Success", "data": { "id": 1, "activity_group_id": "1", "title": "todoTestingUpdated", "is_active": false, "priority": "very-high", "created_at": "2021-11-30T05:29:24.000Z", "updated_at": "2021-12-09T02:17:15.758Z", "deleted_at": null } }`, Status: 200},
	{Body: `{ "status": "Not Found", "message": "Todo with ID 999999999 Not Found", "data": {} }`, Status: 404},
	{Body: `{ "status": "Success", "message": "Success", "data": { "id": 1, "activity_group_id": "1", "title": "todoTestingUpdated", "is_active": false, "priority": "very-high", "created_at": "2021-11-30T05:29:24.000Z", "updated_at": "2021-12-09T02:17:15.758Z", "deleted_at": null } }`, Status: 200},
	{Body: `{ "status": "Not Found", "message": "Todo with ID 999999999 Not Found", "data": {} }`, Status: 404},
	{Body: `{ "status": "Success", "message": "Success", "data": [{ "id": 1, "activity_group_id": "1", "title": "todoTestingUpdated", "is_active": false, "priority": "very-high", "created_at": "2021-11-30T05:29:24.000Z", "updated_at": "2021-12-09T02:17:15.758Z", "deleted_at": null }] }`, Status: 200},
	{Body: `{ "status": "Success", "message": "Success", "data": [] }`, Status: 200},
}

var req int

func main() {
	req = 0
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":3030", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	text := result[req].Body
	w.WriteHeader(result[req].Status)
	w.Write([]byte(text))
	req = req + 1
}

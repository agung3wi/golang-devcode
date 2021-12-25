package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"unicode"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Activity struct {
	ID        int         `gorm:"primaryKey" json:"id"`
	Email     string      `gorm:"size:255" json:"email"`
	Title     string      `gorm:"size:255" json:"title"`
	CreatedAt string      `gorm:"size:255" json:"created_at"`
	UpdatedAt string      `gorm:"size:255" json:"updated_at"`
	DeletedAt interface{} `gorm:"type:varchar(255)" json:"deleted_at"`
}

type Todo struct {
	ID              int         `gorm:"primaryKey" json:"id"`
	ActivityGroupId interface{} `gorm:"type:varchar(255)" json:"activity_group_id"`
	Title           interface{} `gorm:"type:varchar(255)" json:"title"`
	IsActive        interface{} `gorm:"type:varchar(255)" json:"is_active"`
	Priority        interface{} `gorm:"type:varchar(255)" json:"priority"`
	CreatedAt       string      `gorm:"size:255" json:"created_at"`
	UpdatedAt       string      `gorm:"size:255" json:"updated_at"`
	DeletedAt       interface{} `gorm:"type:varchar(255)" json:"deleted_at"`
}

type response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Kosong struct {
}

var kosong Kosong
var activities = []Activity{}
var todos = []Todo{}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_HOST") + ":3306)/" + os.Getenv("MYSQL_DBNAME")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Database Not Connected")
	} else {
		db.AutoMigrate(&Activity{})
		db.AutoMigrate(&Todo{})
	}
	// r := mux.NewRouter()
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/activity-groups", ActivityRest)
	http.HandleFunc("/todo-items", TodoRest)
	http.ListenAndServe(":3030", nil)
}

func ActivityRest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var t Activity
		err := decoder.Decode(&t)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}

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

		t.CreatedAt = "2021-12-01T09:23:05.825Z"
		t.UpdatedAt = "2021-12-01T09:23:05.825Z"
		t.DeletedAt = nil
		t.ID = len(activities) + 1
		activities = append(activities, t)

		var resp response
		w.WriteHeader(http.StatusCreated)
		resp.Status = "Success"
		resp.Message = "Success"
		resp.Data = t
		jData, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}
		w.Write(jData)

	case "GET":
		var resp response
		resp.Status = "Success"
		resp.Message = "Success"
		data := []Activity{}
		for i := range activities {
			if activities[i].DeletedAt == nil {
				data = append(data, activities[i])
			}
		}
		resp.Data = data
		jData, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}
		w.Write(jData)

	default:
		http.Error(w, "", http.StatusBadRequest)
	}

}

func TodoRest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var t Todo
		err := decoder.Decode(&t)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}

		if t.ActivityGroupId == "" || t.ActivityGroupId == nil {
			var resp response
			resp.Status = "Bad Request"
			resp.Message = "activity_group_id cannot be null"
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

		if t.Title == "" || t.Title == nil {
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

		if t.Priority == "" || t.Priority == nil {
			t.Priority = "very-high"
		}

		t.IsActive = "1"
		t.CreatedAt = "2021-12-01T09:23:05.825Z"
		t.UpdatedAt = "2021-12-01T09:23:05.825Z"
		t.DeletedAt = nil
		t.ID = len(todos) + 1
		todos = append(todos, t)

		t.IsActive = true

		var resp response
		w.WriteHeader(http.StatusCreated)
		resp.Status = "Success"
		resp.Message = "Success"
		resp.Data = t
		jData, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}
		w.Write(jData)

	case "GET":
		var resp response
		resp.Status = "Success"
		resp.Message = "Success"

		param1 := r.URL.Query().Get("activity_group_id")

		if param1 != "" {

			data := []Todo{}
			for i := range todos {
				if todos[i].ActivityGroupId == param1 {
					data = append(data, todos[i])
				}
			}

			resp.Data = data
		} else {
			resp.Data = todos
		}
		jData, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}
		w.Write(jData)

	default:
		http.Error(w, "", http.StatusBadRequest)
	}

}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func HandleParamActivity(w http.ResponseWriter, r *http.Request, ids string) {
	var resp response
	id, err := strconv.Atoi(ids)
	if err != nil {
		return
	}

	if id > len(activities) {
		resp.Status = "Not Found"
		resp.Message = "Activity with ID " + ids + " Not Found"
		resp.Data = kosong
		jData, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write(jData)

	} else {

		if r.Method == "GET" {
			resp.Status = "Success"
			resp.Message = "Success"
			resp.Data = activities[id-1]
			jData, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}
			w.Write(jData)
			return
		} else if r.Method == "PATCH" {
			decoder := json.NewDecoder(r.Body)
			var t Activity
			err := decoder.Decode(&t)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}

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

			activities[id-1].Title = t.Title
			if t.Email != "" {
				activities[id-1].Email = t.Email
			}

			resp.Status = "Success"
			resp.Message = "Success"
			resp.Data = activities[id-1]
			jData, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}
			w.Write(jData)
			return

		} else if r.Method == "DELETE" {
			resp.Status = "Success"
			resp.Message = "Success"
			activities[id-1].DeletedAt = "2021-12-01T09:23:05.825Z"
			resp.Data = kosong
			jData, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}
			w.Write(jData)
			return

		}
	}
	return
}

func HandleParamTodo(w http.ResponseWriter, r *http.Request, ids string) {
	var resp response
	id, err := strconv.Atoi(ids)
	if err != nil {
		return
	}

	if id > len(todos) {
		resp.Status = "Not Found"
		resp.Message = "Todo with ID " + ids + " Not Found"
		resp.Data = kosong
		jData, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write(jData)

	} else {

		if r.Method == "GET" {
			resp.Status = "Success"
			resp.Message = "Success"
			resp.Data = todos[id-1]
			jData, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}
			w.Write(jData)
			return
		} else if r.Method == "PATCH" {
			decoder := json.NewDecoder(r.Body)
			var t Todo
			err := decoder.Decode(&t)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}

			todos[id-1].IsActive = t.IsActive

			if t.Title != "" && t.Title != nil {
				todos[id-1].Title = t.Title
			}

			resp.Status = "Success"
			resp.Message = "Success"
			resp.Data = todos[id-1]
			jData, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}
			w.Write(jData)
			return

		} else if r.Method == "DELETE" {
			resp.Status = "Success"
			resp.Message = "Success"
			activities[id-1].DeletedAt = "2021-12-01T09:23:05.825Z"
			resp.Data = kosong
			jData, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}
			w.Write(jData)
			return

		}
	}
	return
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var path = r.URL.Path
	lenPath := len(path)
	if lenPath > 17 {
		if path[0:17] == "/activity-groups/" && isInt(path[17:lenPath]) {
			ids := path[17:lenPath]
			HandleParamActivity(w, r, ids)
			return
		}
	}

	if lenPath > 12 {
		if path[0:12] == "/todo-items/" && isInt(path[12:lenPath]) {
			ids := path[12:lenPath]
			HandleParamTodo(w, r, ids)
			return
		}
	}

	names := "Oke"
	jData, err := json.Marshal(names)
	if err != nil {
		fmt.Fprint(w, "Internal Server Error")
	} else {
		w.Write(jData)
	}

}

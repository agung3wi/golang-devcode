package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
)

type Activity struct {
	ID        int         `json:"id"`
	Email     string      `json:"email"`
	Title     string      `json:"title"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	DeletedAt interface{} `json:"deleted_at"`
}

type Todo struct {
	ID              int         `json:"id"`
	ActivityGroupId interface{} `json:"activity_group_id"`
	Title           interface{} `json:"title"`
	IsActive        interface{} `json:"is_active"`
	Priority        interface{} `json:"priority"`
	CreatedAt       string      `json:"created_at"`
	UpdatedAt       string      `json:"updated_at"`
	DeletedAt       interface{} `json:"deleted_at"`
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
var resp response

// var currentActivity int
// var currentTodo int

// var db *gorm.DB
var err error
var db *sql.DB

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// dsn := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_HOST") + ":3306)/" + os.Getenv("MYSQL_DBNAME")
	// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	fmt.Println("Database Not Connected")
	// } else {
	// db.AutoMigrate(&Activity{})
	// db.AutoMigrate(&Todo{})
	// }

	db, err = sql.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp("+os.Getenv("MYSQL_HOST")+":3306)/"+os.Getenv("MYSQL_DBNAME"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.Query(`CREATE TABLE IF NOT EXISTS activities (
		id bigint(20) NOT NULL,
		email varchar(255) DEFAULT NULL,
		title varchar(255) DEFAULT NULL,
		created_at varchar(255) DEFAULT NULL,
		updated_at varchar(255) DEFAULT NULL,
		deleted_at varchar(255) DEFAULT NULL
	  ) ENGINE=InnoDB DEFAULT CHARSET=latin1;`)

	db.Query(`CREATE TABLE IF NOT EXISTS todos (
		id bigint(20) NOT NULL,
		activity_group_id varchar(255) DEFAULT NULL,
		title varchar(255) DEFAULT NULL,
		is_active varchar(255) DEFAULT NULL,
		priority varchar(255) DEFAULT NULL,
		created_at varchar(255) DEFAULT NULL,
		updated_at varchar(255) DEFAULT NULL,
		deleted_at varchar(255) DEFAULT NULL
	  ) ENGINE=InnoDB DEFAULT CHARSET=latin1;`)

	// currentActivity = 1
	// currentTodo = 1
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":3030", nil)
}

func ActivityRest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var activity Activity
		err := decoder.Decode(&activity)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}

		if activity.Title == "" {

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

		activity.CreatedAt = "2021-12-01T09:23:05.825Z"
		activity.UpdatedAt = "2021-12-01T09:23:05.825Z"
		activity.DeletedAt = nil
		activity.ID = len(activities) + 1

		w.WriteHeader(http.StatusCreated)
		resp.Status = "Success"
		resp.Message = "Success"
		resp.Data = activity
		jData, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}
		// go func() {
		// currentActivity = currentActivity + 1
		activities = append(activities, activity)
		db.Exec("INSERT INTO activities(id,title,email) VALUES(?,?,?)", activity.ID, activity.Title, activity.Email)

		// db.Create(&activity)
		// }()
		w.Write(jData)

	case "GET":

		resp.Status = "Success"
		resp.Message = "Success"
		// data := []Activity{}
		// for i := range activities {
		// 	if activities[i].DeletedAt == nil {
		// 		data = append(data, activities[i])
		// 	}
		// }
		resp.Data = activities
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
		t.IsActive = true

		w.WriteHeader(http.StatusCreated)
		resp.Status = "Success"
		resp.Message = "Success"
		resp.Data = t
		jData, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}
		// go func() {
		// db.Create(&t)
		todos = append(todos, t)
		db.Exec("INSERT INTO todos(id,title,activity_group_id,is_active,priority,created_at,updated_at) VALUES(?,?,?,?,?,?,?)", t.ID, t.Title, t.ActivityGroupId, t.IsActive, t.Priority, t.CreatedAt, t.UpdatedAt)
		// currentTodo = currentTodo + 1

		// }()
		w.Write(jData)

	case "GET":

		resp.Status = "Success"
		resp.Message = "Success"

		param1 := r.URL.Query().Get("activity_group_id")

		if param1 != "" {
			resp.Data = []Todo{}
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
			var activity Activity
			decoder := json.NewDecoder(r.Body)

			err := decoder.Decode(&activity)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}

			if activity.Title == "" {

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

			activities[id-1].Title = activity.Title
			if activity.Email != "" {
				activities[id-1].Email = activity.Email
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

	if path == "/todo-items" {
		TodoRest(w, r)
		return
	}

	if path == "/activity-groups" {
		ActivityRest(w, r)
		return
	}

	if lenPath > 12 {
		if path[0:12] == "/todo-items/" && isInt(path[12:lenPath]) {
			ids := path[12:lenPath]
			HandleParamTodo(w, r, ids)
			return
		}
	}

	if lenPath > 17 {
		if path[0:17] == "/activity-groups/" && isInt(path[17:lenPath]) {
			ids := path[17:lenPath]
			HandleParamActivity(w, r, ids)
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

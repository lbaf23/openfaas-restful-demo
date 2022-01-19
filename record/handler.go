package function

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error

	secretfile, _ := ioutil.ReadFile("/var/openfaas/secrets/openfaas-restful-demo-db-password")
	password := strings.Split(string(secretfile), "\n")[0]

	user := os.Getenv("postgres_user")
	host := os.Getenv("postgres_host")
	dbName := os.Getenv("postgres_db")
	port := os.Getenv("postgres_port")
	sslmode := os.Getenv("postgres_sslmode")

	connStr := "postgres://" + user + ":" + password + "@" +
		host + ":" + port + "/" + dbName + "?sslmode=" + sslmode

	fmt.Println(connStr)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}
}

type Response struct {
	Data    []Record `json:"data"`
	Message string   `json:"message"`
	Code    int      `json:"code"`
}

type Record struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"createAt"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		Get(w, r)
	} else if r.Method == http.MethodPost {
		Post(w, r)
	} else if r.Method == http.MethodPut {
		Put(w, r)
	} else if r.Method == http.MethodDelete {
		Delete(w, r)
	} else {
		w.WriteHeader(http.StatusBadGateway)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	var response Response

	rows, err := db.Query(`select * from record`)

	if err != nil {
		log.Printf("Record: get error: %s", err.Error())
		response = Response{
			Message: "error",
			Code:    500,
		}

		res, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}

	records := []Record{}
	defer rows.Close()
	for rows.Next() {
		result := Record{}
		scanErr := rows.Scan(&result.Id, &result.Title, &result.Content, &result.CreateAt)
		if scanErr != nil {
			log.Println("Record: scan err:", scanErr.Error())
		}
		records = append(records, result)
	}
	response = Response{
		Data:    records,
		Message: "ok",
		Code:    200,
	}
	res, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Put(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "ok",
		Code:    200,
	}

	res, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Post(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "ok",
		Code:    200,
	}
	res, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "ok",
		Code:    200,
	}
	res, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

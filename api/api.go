package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"syndio/db"

	"github.com/gorilla/mux"
)

func setJsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func fromJson[T any](body io.Reader, target *T) error {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf.Bytes(), target)
}

func returnJson[T any](w http.ResponseWriter, withData func() (T, error)) {
	setJsonHeader(w)
	data, serverError := withData()

	if serverError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		serverErrorJson, err := json.Marshal(&serverError)
		if err != nil {
			log.Print(err)
			return
		}
		w.Write(serverErrorJson)
		return
	}

	dataJson, err := json.Marshal(&data)
	if err != nil {
		log.Print("Failed to marshal data:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(dataJson)
}

func returnErr(w http.ResponseWriter, err error, code int) {
	returnJson(w, func() (interface{}, error) {
		w.WriteHeader(code)
		return map[string]string{"error": err.Error()}, nil
	})
}

type EmployeeJob struct {
	Department string `json:"department"`
	JobTitle   string `json:"job_title"`
}

func CreateEmployeeJob(database db.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		employeeIdStr := vars["employeeId"]
		employeeId, err := strconv.Atoi(employeeIdStr)
		if err != nil {
			returnErr(w, errors.New("invalid employee ID"), http.StatusBadRequest)
			return
		}

		var entry EmployeeJob
		if err := fromJson(r.Body, &entry); err != nil {
			returnErr(w, errors.New("invalid JSON payload"), http.StatusBadRequest)
			return
		}

		addedId, err := database.AddEmployeeJob(employeeId, entry.Department, entry.JobTitle)
		if err != nil {
			returnErr(w, errors.New("failed to add employee job"), http.StatusInternalServerError)
			return
		}

		returnJson(w, func() (interface{}, error) {
			return map[string]interface{}{
				"message": "employee job added successfully",
				"id":      addedId,
			}, nil
		})
	})
}

func UpdateEmployeeJob(database db.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		jobIdStr := vars["jobId"]
		jobId, err := strconv.Atoi(jobIdStr)
		if err != nil {
			returnErr(w, errors.New("invalid job ID"), http.StatusBadRequest)
			return
		}

		var entry EmployeeJob
		if err := fromJson(r.Body, &entry); err != nil {
			returnErr(w, errors.New("invalid JSON payload"), http.StatusBadRequest)
			return
		}

		returnJson(w, func() (interface{}, error) {
			return database.UpdateEmployeeJob(jobId, entry.Department, entry.JobTitle)
		})
	})
}

func GetEmployeeJobsByEmployeeId(database db.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		employeeIdStr := vars["employeeId"]
		employeeId, err := strconv.Atoi(employeeIdStr)
		if err != nil {
			returnErr(w, errors.New("invalid employee ID"), http.StatusBadRequest)
			return
		}

		returnJson(w, func() (interface{}, error) {
			return database.GetEmployeeJobByEmployeeId(employeeId)
		})
	})
}

func GetEmployeeJobs(database db.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		returnJson(w, func() (interface{}, error) {
			return database.GetEmployeeJobs()
		})
	})
}

func GetEmployees(database db.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		returnJson(w, func() (interface{}, error) {
			return database.GetEmployees()
		})
	})
}

func Serve(bind string, database db.DB) {
	router := mux.NewRouter()
	router.Handle("/employees", GetEmployees(database)).Methods("GET")
	router.Handle("/employees/{employeeId}/jobs", CreateEmployeeJob(database)).Methods("POST")
	router.Handle("/employees/{employeeId}/jobs", GetEmployeeJobsByEmployeeId(database)).Methods("GET")
	router.Handle("/employees/jobs", GetEmployeeJobs(database)).Methods("GET")
	router.Handle("/employees/jobs/{jobId}", UpdateEmployeeJob(database)).Methods("PATCH")

	log.Printf("Server listening on %s\n", bind)
	log.Fatal(http.ListenAndServe(bind, router))
}

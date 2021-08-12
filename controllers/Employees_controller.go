package controllers

import (
	"encoding/json"
	"fmt"
	Emp_model "golang_restful_api/models"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type ResponseEmployees struct {
	Employees []Emp_model.Employees `json:"employees"`
}
type ResponseError struct {
	Status   int           `json:"status"`
	Error    int           `json:"error"`
	Messages MessagesError `json:"messages"`
}
type MessagesError struct {
	Error string `json:"error"`
}
type ResponseSuccess struct {
	Status   int             `json:"status"`
	Error    string          `json:"error"`
	Messages MessagesSuccess `json:"messages"`
}
type MessagesSuccess struct {
	Success string `json:"success"`
}

func Index(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
	if r.Method == "GET" {
		employees_data, err := Emp_model.GetDataEmployeeAll()
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
		}

		var rt = ResponseEmployees{
			Employees: employees_data,
		}
		rs, err := json.Marshal(rt)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
		}

		rw.WriteHeader(200)
		rw.Write(rs)

	}

}
func Show(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		vars := mux.Vars(r)
		emp_data, err := Emp_model.GetDataEmployeeByID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
		}
		if len(emp_data) > 0 { // found employee
			rs, err := json.Marshal(emp_data)
			if err != nil {
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write([]byte(err.Error()))
			}
			rw.WriteHeader(200)
			rw.Write(rs)
		} else { //No found employee
			// var res ResponseError
			// res.Status =

			messagesError := ResponseError{
				Status: http.StatusNotFound,
				Error:  http.StatusNotFound,
				Messages: MessagesError{
					Error: "No employee found",
				},
			}
			rs, err := json.Marshal(messagesError)
			if err != nil {
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write([]byte(err.Error()))
			}
			rw.WriteHeader(http.StatusNotFound)
			rw.Write(rs)
		}

	}
}
func Create(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}
		var e Emp_model.EmployeesPost
		err = json.Unmarshal(b, &e)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}
		err = Emp_model.CreateDataEmployee(e)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}
		var rs = ResponseSuccess{
			Status: http.StatusCreated,
			Error:  "null",
			Messages: MessagesSuccess{
				Success: "Employee created successfully",
			},
		}
		rt, _ := json.Marshal(rs)
		rw.WriteHeader(200)
		rw.Write(rt)
	}
}

func Update(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		vars := mux.Vars(r)
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}
		var e Emp_model.EmployeesPost
		err = json.Unmarshal(b, &e)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}
		err = Emp_model.UpdateDataEmployee(vars["id"], e)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}
		var rs = ResponseSuccess{
			Status: http.StatusOK,
			Error:  "null",
			Messages: MessagesSuccess{
				Success: "Employee updated successfully",
			},
		}
		rt, _ := json.Marshal(rs)
		rw.WriteHeader(200)
		rw.Write(rt)
	}
}

func Delate(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		vars := mux.Vars(r)
		err := Emp_model.DeleteDataEmployee(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}
		var rs = ResponseSuccess{
			Status: http.StatusOK,
			Error:  "null",
			Messages: MessagesSuccess{
				Success: "Employee successfully deleted",
			},
		}
		rt, _ := json.Marshal(rs)
		rw.WriteHeader(200)
		rw.Write(rt)
	}
}

package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Employees struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type EmployeesPost struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ci_rest_api")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	return db, nil
}
func GetDataEmployeeAll() ([]Employees, error) {
	db, err := ConnectDB()
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	defer db.Close()
	result, err := db.Query("SELECT * FROM employees ORDER BY id DESC")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	var e Employees
	var arrEmployees = []Employees{}
	for result.Next() {
		err := result.Scan(
			&e.Id,
			&e.Name,
			&e.Email,
		)
		if err != nil {
			return nil, err
		}
		arrEmployees = append(arrEmployees, e)
	}

	return arrEmployees, err
}

func GetDataEmployeeByID(id string) ([]Employees, error) {
	db, err := ConnectDB()
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	defer db.Close()
	result, err := db.Query("SELECT * FROM employees WHERE id=" + id + " ORDER BY id DESC")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	var e Employees
	var arrEmployees = []Employees{}
	for result.Next() {
		err := result.Scan(
			&e.Id,
			&e.Name,
			&e.Email,
		)
		if err != nil {
			return nil, err
		}
		arrEmployees = append(arrEmployees, e)
	}

	return arrEmployees, err
}

func CreateDataEmployee(e EmployeesPost) error {
	db, err := ConnectDB()
	if err != nil {
		fmt.Print(err)
		return err
	}
	defer db.Close()
	_, err = db.Query("INSERT INTO employees (name,email) VALUES ('" + e.Name + "','" + e.Email + "')")
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}
func UpdateDataEmployee(id string, e EmployeesPost) error {
	db, err := ConnectDB()
	if err != nil {
		fmt.Print(err)
		return err
	}
	defer db.Close()
	_, err = db.Query("UPDATE employees SET name='" + e.Name + "', email='" + e.Email + "' WHERE id=" + id)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

func DeleteDataEmployee(id string) error {
	db, err := ConnectDB()
	if err != nil {
		fmt.Print(err)
		return err
	}
	defer db.Close()
	_, err = db.Query("DELETE FROM employees WHERE id=" + id)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

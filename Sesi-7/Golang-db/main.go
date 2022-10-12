package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID       int
	FullName string
	Email    string
	Age      int
	Division string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345678"
	dbname   = "db-go-sql"
)

var (
	db  *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	// CreateEmployee()
	// GetEmployees()
	// UpdateEmployee()
	DeleteEmployee()
}

func CreateEmployee() {
	var employee = Employee{}

	sqlStatement :=
		`INSERT INTO employess (full_name, email, age,division)
	VALUES ($1, $2, $3, $4)
	Returning *`

	err = db.QueryRow(sqlStatement, "GhifariAhmadL", "hahaha@gmail.com", 22, "IT").
		Scan(&employee.ID, &employee.FullName, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data : %+v\n", employee)
}

func GetEmployees() {
	var result = []Employee{}

	sqlStatement := `SELECT * FROM employess`

	row, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer row.Close()

	for row.Next() {
		var employee = Employee{}

		err = row.Scan(&employee.ID, &employee.FullName, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}

		result = append(result, employee)
	}

	fmt.Println("Employee data:", result)
}

func UpdateEmployee() {
	sqlStatement := `
	UPDATE employess
	SET full_name =$2, email = $3, division = $4, age =$5
	WHERE id = $1`
	res, err := db.Exec(sqlStatement, 1, "AhmadGhifari", "ghifariahmad@upnvj.ac.id", "Manager", 23)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Update Data amount:", count)
}

func DeleteEmployee() {
	sqlStatement := `
	DELETE from employess
	WHERE id = $1`

	res, err := db.Exec(sqlStatement, 1)

	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Delete amount Data:", count)
}

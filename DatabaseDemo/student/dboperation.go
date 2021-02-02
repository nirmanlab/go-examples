// Package classification Petstore API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /v2
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: John Doe<john.doe@example.com> http://john.doe.com
//
//     Consumes:
//     - application/json
//     - application/xml
//
//     Produces:
//     - application/json
//     - application/xml
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: KEY
//          in: header
//     oauth2:
//         type: oauth2
//         authorizationUrl: /oauth2/auth
//         tokenUrl: /oauth2/token
//         in: header
//         scopes:
//           bar: foo
//         flow: accessCode
//
//     Extensions:
//     x-meta-value: value
//     x-meta-array:
//       - value1
//       - value2
//     x-meta-array-obj:
//       - name: obj
//         value: field
//
// swagger:meta
package student

import (
	"databasedemo/graph/model"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connection() *sqlx.DB {
	db, err := sqlx.Open("postgres", "host=localhost user=school_admin dbname=school password=school_admin sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// InsertStudent inserts the student into database
func InsertStudent(db *sqlx.DB, name string) (int, error) {
	id := 0
	err := db.QueryRow("INSERT INTO students (name) VALUES($1)RETURNING id", name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err
}

// FetchStudent return students
func FetchStudent(db *sqlx.DB) *sqlx.Rows {
	raw, err := db.Queryx("SELECT * FROM students")
	if err != nil {
		log.Fatal(err)
	}
	return raw
}

// FetchStudentbyid return student
func FetchStudentbyid(id int, db *sqlx.DB) model.Student {
	var s model.Student
	err := db.QueryRowx("SELECT * FROM students Where id=$1", id).StructScan(&s)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

//UpdateStudent update student from database
func UpdateStudent(db *sqlx.DB, name string, id int) {
	_, err := db.Exec("UPDATE students SET name=$2 WHERE id=$1", id, name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Name updated succesfully")
}

// DeleteStudent delete student from database
func DeleteStudent(db *sqlx.DB, id int64) {
	_, err := db.Exec("DELETE FROM students where id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Name Deleted Succesfully")
}

//CallFunctionFromDatabase
func CallFunctionFromDatabase(db *sqlx.DB, name string) {
	_, err := db.Exec("CALL insert_data($1)", name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Function called")
}

// Total student
func TotalStudent(db *sqlx.DB) {
	var count int64 = 0
	err := db.Get(&count, "SELECT totalStudent()")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}

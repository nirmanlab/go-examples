package Students

import (
	"log"

	"github.com/jmoiron/sqlx"
)

// InsertStudent : Insert student value in to database
func InsertStudent(student *StudentsModel, db *sqlx.DB) (int, error) {
	id := 0
	err := db.QueryRow("INSERT INTO students (name,name_fr) VALUES ($1,$2)RETURNING id", student.Name, student.NameFr).Scan(&id)
	if err != nil {
		log.Println(err)
	}
	return id, nil
}

// FetchStudent :Fetching Student value from database
func FetchStudent(db *sqlx.DB) (*sqlx.Rows, error) {
	raw, err := db.Queryx("SELECT * FROM students")
	if err != nil {
		log.Fatal(err)
	}
	return raw, nil
}

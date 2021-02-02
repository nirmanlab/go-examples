package main

import (
	"databasedemo/student"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db := connection()

	defer db.Close()

	// s := student.CreateStrudent(0, "Sagar")
	// id, err := student.InsertStudent(db, s.StudentName)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(id)
	// raw := student.FetchStudent(db)
	// for raw.Next() {
	// 	var s student.Student
	// 	err := raw.StructScan(&s)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("ID:-%d, Name :- %s \n", s.StudentId, s.StudentName)
	// }
	// a := student.FetchStudentbyid(1, *db)
	// fmt.Printf("ID:-%d, Name :- %s \n", a.StudentId, a.StudentName)
	//student.UpdateStudent(db, "John Deo", 6)
	student.CallFunctionFromDatabase(db, "Brijesh")
	student.TotalStudent(db)

}

//Establish Connection to Database
func connection() *sqlx.DB {
	db, err := sqlx.Open("postgres", "host=localhost user=school_admin dbname=school password=school_admin sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

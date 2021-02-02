package main

import (
	"GolangTranslatedemo/Students"
	"fmt"

	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {

	db, err := sqlx.Open("postgres", "host=localhost user=school_admin dbname=school password=school_admin sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer func() { db.Close() }()
	raws, _ := Students.FetchStudent(db)
	for raws.Next() {
		s := Students.StudentsModel{}
		err := raws.StructScan(&s)
		if err != nil {
			log.Fatal(err)
		}
		message.SetString(language.French, "Student's Name %v", "Nom de l'étudiant %v")
		p := message.NewPrinter(language.Make("fr"))
		p.Printf("Student's Name %v", s.NameFr)
		fmt.Println()
		p = message.NewPrinter(language.Make("en"))
		p.Printf("Student's Name %v", s.Name)
		fmt.Println()
	}

}

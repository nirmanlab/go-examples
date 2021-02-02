package Students

//StudentsModel	 : Model of Student
type StudentsModel struct {
	Id     int    `db:"id"`
	Name   string `db:"name"`
	NameFr string `db:"name_fr"`
}

//CreateStudents : Method for Createing Student
func CreateStudents(name, nameFr string) *StudentsModel {
	return &StudentsModel{
		Name:   name,
		NameFr: nameFr,
	}
}

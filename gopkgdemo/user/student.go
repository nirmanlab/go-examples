package user

// Student model
type Student struct {
	U         User
	StudentID int
}

// CreateStudent creating student
func CreateStudent(u User, id int) *Student {
	return &Student{
		U:         u,
		StudentID: id,
	}
}

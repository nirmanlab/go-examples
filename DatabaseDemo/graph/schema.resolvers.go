package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"databasedemo/graph/generated"
	"databasedemo/graph/model"
	"databasedemo/student"
	"fmt"
	"log"
)

func (r *mutationResolver) CreateStudent(ctx context.Context, input model.NewStudent) (*model.Student, error) {
	createUser, err := student.InsertStudent(r.Db, input.Name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created User ID:- ", createUser)
	return &model.Student{ID: createUser, Name: input.Name}, nil
}

func (r *mutationResolver) UpdateStudent(ctx context.Context, input model.UpdateStudent) (*model.Student, error) {
	student.UpdateStudent(r.Db, input.NewName, input.ID)

	raw := student.FetchStudentbyid(input.ID, r.Db)

	return &raw, nil
}

func (r *queryResolver) Student(ctx context.Context) ([]*model.Student, error) {
	raws := student.FetchStudent(r.Db)

	students := make([]*model.Student, 0, 0)
	for raws.Next() {
		s := &model.Student{}
		err := raws.StructScan(&s)
		if err != nil {
			log.Fatalln(err)
		}
		students = append(students, s)
	}

	return students, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

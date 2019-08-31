package person

import "github.com/graph-gophers/graphql-go"

type person struct {
	ID        graphql.ID
	FirstName string
	LastName  string
}

type personResolver struct {
	p *person
}

func (r *personResolver) ID() graphql.ID {
	return r.p.ID
}

func (r *personResolver) FirstName() string {
	return r.p.FirstName
}

func (r *personResolver) LastName() *string {
	return &r.p.LastName
}

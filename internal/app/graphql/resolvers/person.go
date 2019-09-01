package resolvers

import "github.com/graph-gophers/graphql-go"

// Person : Resolver function for the "Person" query
func (r *QueryResolver) Person(args struct{ ID graphql.ID }) *PersonResolver {
	if p := peopleData[args.ID]; p != nil {
		return &PersonResolver{p}
	}
	return nil
}

// Person person
type Person struct {
	ID        graphql.ID
	FirstName string
	LastName  string
}

// PersonResolver Person resolver
type PersonResolver struct {
	p *Person
}

// ID id for person
func (r *PersonResolver) ID() graphql.ID {
	return r.p.ID
}

// FirstName first name for person
func (r *PersonResolver) FirstName() string {
	return r.p.FirstName
}

// LastName last name for
func (r *PersonResolver) LastName() *string {
	return &r.p.LastName
}

var people = []*Person{
	{
		ID:        "1000",
		FirstName: "Pedro",
		LastName:  "Marquez",
	},

	{
		ID:        "1001",
		FirstName: "John",
		LastName:  "Doe",
	},
}

var peopleData = make(map[graphql.ID]*Person)

package graphql

import "github.com/graph-gophers/graphql-go"

// GraphQL graphql
type GraphQL struct {
	Schema   *graphql.Schema
	Resolver *Resolver
}

// StitchSchema stitches together schema fragments
func (g *GraphQL) StitchSchema(fragments *[]SchemaStitcher, _ error) {

}

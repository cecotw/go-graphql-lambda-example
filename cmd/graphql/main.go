package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cecotw/go-graphql-lambda-example/internal/app/graphql/resolvers"
	"github.com/cecotw/go-graphql-lambda-example/internal/pkg/handler"
)

// Schema schema
var Schema = `
	schema {
		query: Query
	}
	type Person{
		id: ID!
		firstName: String!
		lastName: String
	}
	type Query{
		person(id: ID!): Person
	}
`

var graphql = new(handler.GraphQl)

func init() {
	graphql.BuildSchema(Schema, &resolvers.QueryResolver{})
}

func main() {
	lambda.Start(graphql.Lambda)
}

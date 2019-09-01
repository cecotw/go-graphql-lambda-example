package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cecotw/go-graphql-lambda-example/internal/app/graphql/resolvers"
	"github.com/cecotw/go-graphql-lambda-example/internal/pkg/handler"
	"github.com/cecotw/go-graphql-lambda-example/internal/pkg/loader"
)

var graphql = new(handler.GraphQl)

func init() {
	var schemaLoader = &loader.Schema{
		SchemaPaths: []string{
			"./internal/app/graphql/schemas/schema.graphql",
			"./internal/app/graphql/schemas/types/person.graphql",
		},
	}
	var schema = schemaLoader.MergeSchema()
	graphql.BuildSchema(schema, &resolvers.QueryResolver{})
}

func main() {
	lambda.Start(graphql.Lambda)
}

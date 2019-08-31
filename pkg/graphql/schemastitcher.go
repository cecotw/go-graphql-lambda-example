package graphql

// SchemaStitcher loads schemas and resolvers
type SchemaStitcher interface {
	GetSchema() string
	GetResolver() interface{}
}

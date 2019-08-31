package graphql

// Resolver : Struct with all the resolver functions
type Resolver struct{}

// ResolverBuilder builds resolver
type ResolverBuilder interface {
	AddQuery(*Resolver)
}

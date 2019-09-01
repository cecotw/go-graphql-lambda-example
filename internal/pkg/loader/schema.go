package loader

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// Schema concatenates the .graphql schema files
type Schema struct {
	SchemaPaths []string
}

// MergeSchema merges the schema files to one string
func (s *Schema) MergeSchema() string {
	var schema string
	basepath, _ := os.Getwd()
	for _, localpath := range s.SchemaPaths {
		schemaText, err := s.readFile(path.Join(basepath, localpath))
		if err != nil {
			fmt.Println(err)
		}
		schema += schemaText
	}
	return schema
}

func (s *Schema) readFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

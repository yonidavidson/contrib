package entgql

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	tparse "text/template/parse"

	"entgo.io/ent/entc/gen"
)

// RemoveOldGeneratedFiles removes old generated files without gql
func RemoveOldGeneratedFiles(prefix string) gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			for _, rootTemplate := range AllTemplates {
				for _, template := range rootTemplate.Templates() {
					if tparse.IsEmptyTree(template.Root) {
						continue
					}
					if !strings.HasPrefix(template.Name(), "gql_") {
						continue
					}
					fmt.Println(rootTemplate.Name(), "->", template.Name())
					deleteByName(g, strings.TrimPrefix(template.Name(), "gql_"))
				}
			}
			return next.Generate(g)
		})
	}
}

func deleteByName(g *gen.Graph, name string) error {
	for _, n := range g.Nodes {
		if n.Package() == name {
			return nil
		}
	}
	err := os.Remove(filepath.Join(g.Target, name+".go"))
	if !os.IsNotExist(err) {
		println("deleted:", name)
		return err
	}
	return nil
}


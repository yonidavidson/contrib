package entgql

import (
	"os"
	"path/filepath"
	"strings"
	tparse "text/template/parse"

	"entgo.io/ent/entc/gen"
	"github.com/prometheus/common/log"
)

func removeOldGeneratedFiles() gen.Hook {
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
					name := strings.TrimPrefix(template.Name(), "gql_")
					if deleteByName(g, name) != nil {
						log.Info("deleted gql file with old name: ", name)
					}
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
		return err
	}
	//if err != nil {
	//	log.Fatal(err)
	//}
	return nil
}

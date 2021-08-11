package blog

import (
	"entgo.io/contrib/entprom/internal/ent/hook"

	"entgo.io/contrib/entprom/internal/ent"
)

func main() {
	client, _ := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	// Add a hook only on user mutations.
	client.File.Use(exampleHook())

	//// Add a hook only on update operations.
	client.Use(hook.On(exampleHook(), ent.OpUpdate|ent.OpUpdateOne))
	//
	//// Reject delete operations.
	client.Use(hook.Reject(ent.OpDelete | ent.OpDeleteOne))
}

package blog

import (
	"context"

	"entgo.io/ent"
)

func exampleHook() ent.Hook {
	//use this to init your hook
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			// Do something before mutation
			v, err := next.Mutate(ctx, m)
			if err != nil {
				// Do something if error after mutation
			}
			// Do something always after mutation
			return v, err
		})
	}
}

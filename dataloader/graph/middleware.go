package graph

import (
	"context"
	"net/http"
	"time"

	"github.com/dev-sota/gqlgen-gorm/graph/model"
	"gorm.io/gorm"
)

const loadersKey = "dataLoaders"

type Loaders struct {
	UserById       UserLoader
	TodosByUserIDs TodoLoader
}

func Middleware(db *gorm.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			UserById: UserLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []string) ([]*model.User, []error) {
					placeholders := make([]string, len(ids))
					args := make([]interface{}, len(ids))
					for i := 0; i < len(ids); i++ {
						placeholders[i] = "?"
						args[i] = i
					}

					res, err := db.Raw("SELECT * FROM users WHERE id IN ?", ids).Rows()
					if err != nil {
						panic(err)
					}
					defer res.Close()

					userById := map[string]*model.User{}
					for res.Next() {
						user := model.User{}
						err := res.Scan(&user.ID, &user.Name)
						if err != nil {
							panic(err)
						}
						userById[user.ID] = &user
					}

					users := make([]*model.User, len(ids))
					for i, id := range ids {
						users[i] = userById[id]
					}

					return users, nil
				},
			},
			TodosByUserIDs: TodoLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(userIDs []string) ([][]*model.Todo, []error) {
					if len(userIDs) == 0 {
						return nil, nil
					}

					res, err := db.Raw("SELECT * FROM todos WHERE user_id IN ?", userIDs).Rows()
					if err != nil {
						panic(err)
					}
					defer res.Close()

					todoByUserId := map[string][]*model.Todo{}
					for res.Next() {
						todo := model.Todo{}
						err := res.Scan(&todo.ID, &todo.Text, &todo.Done, &todo.UserID)
						if err != nil {
							panic(err)
						}
						todoByUserId[todo.UserID] = append(todoByUserId[todo.UserID], &todo)
					}

					results := make([][]*model.Todo, len(userIDs))
					for i, id := range userIDs {
						results[i] = todoByUserId[id]
					}

					return results, nil
				},
			},
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

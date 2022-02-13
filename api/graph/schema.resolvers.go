package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/S-H-GAMELINKS/go-lang-graphql-todo-app/api/graph/generated"
	model1 "github.com/S-H-GAMELINKS/go-lang-graphql-todo-app/api/graph/model"
	"github.com/S-H-GAMELINKS/go-lang-graphql-todo-app/api/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input *model1.CreateTodoInput) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, id *int, input *model1.UpdateTodoInput) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CompleteTodo(ctx context.Context, id *int) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTodos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) ID(ctx context.Context, obj *model.Todo) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }

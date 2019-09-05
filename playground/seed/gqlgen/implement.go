package gqlgen

import (
	"context"

	prisma "github.com/playground/seed/client"
)

func (r *mutationResolver) CreateNewUser(ctx context.Context, email *string, name *string) (*prisma.Traveler, error) {
	panic("not implemented")
}

func (r *queryResolver) Hello(ctx context.Context, who string) (string, error) {
	panic("not implemented")
}

func (r *queryResolver) Ping(ctx context.Context, number *int) (string, error) {
	panic("not implemented")
}


package gqlgen

import (
	prisma "github.com/playground/seed/client"
)

type Resolver struct {
	Client *prisma.Client
}


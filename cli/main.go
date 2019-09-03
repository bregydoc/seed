package main

import (
	"github.com/bregydoc/seed"
)

func main() {
	err := seed.Plant("seed.graphql", "github.com/playground")
	if err != nil {
		panic(err)
	}

}

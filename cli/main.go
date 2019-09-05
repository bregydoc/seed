package main

import (
	"github.com/bregydoc/seed"
)

func main() {
	project, err := seed.NewDefault("github.com/playground")
	if err != nil {
		panic(err)
	}

	err = project.Plant()
	if err != nil {
		panic(err)
	}

}

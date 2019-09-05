package seed

import (
	"os/exec"
	"strings"
)

// Project represents a seed project
type Project struct {
	Name            string
	AbsoluteWorkDir string
	WorkDir         string
	SeedFilename    string
}

func newDefaultProject(name string) (*Project, error) {
	pwd, err := exec.Command("pwd").Output()
	if err != nil {
		return nil, err
	}

	abs := strings.Trim(string(pwd), " \n\t")

	return &Project{
		Name:            name,
		AbsoluteWorkDir: abs,
		WorkDir:         ".",
		SeedFilename:    "seed.graphql",
	}, nil
}

// NewDefault returns a new default project with your project name
func NewDefault(name string) (*Project, error) {
	return newDefaultProject(name)
}

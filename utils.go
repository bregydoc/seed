package seed

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func goModExists(project *Project) bool {
	_, err := os.Open(path.Join(project.AbsoluteWorkDir, goModFilename))
	if os.IsExist(err) {
		return true
	}

	return false
}

func createGoModFile(project *Project) error {
	filename := path.Join(project.AbsoluteWorkDir, goModFilename)
	body := fmt.Sprintf("module %s\n", project.Name)
	return ioutil.WriteFile(filename, []byte(body), 0644)
}

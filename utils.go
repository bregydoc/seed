package seed

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func goModExists(workDir string) bool {
	_, err := os.Open(path.Join(workDir, goModFilename))
	if os.IsExist(err) {
		return true
	}

	return false
}

func createGoModFile(workDir string, projectName string) error {
	filename := path.Join(workDir, goModFilename)
	body := fmt.Sprintf("module %s\n", projectName)
	return ioutil.WriteFile(filename, []byte(body), 0644)
}

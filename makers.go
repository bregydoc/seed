package seed

import (
	"io/ioutil"
	"os"
	"path"
)

func createSeedDir(workDir string) error {
	full := path.Join(workDir, seedDir)
	err := os.Mkdir(full, os.ModePerm)
	if os.IsExist(err) {
		return nil
	}
	return err
}

func createPrismaDir(workDir string) error {
	full := path.Join(workDir, seedDir, prismaDir)
	err := os.Mkdir(full, os.ModePerm)
	if os.IsExist(err) {
		return nil
	}
	return err
}

func createGQLDir(workDir string) error {
	full := path.Join(workDir, seedDir, gqlDir)
	err := os.Mkdir(full, os.ModePerm)
	if os.IsExist(err) {
		return nil
	}
	return err
}

func writeToPrisma(workDir string, types []GQLType) error {
	reconstructed := ""
	for _, t := range types {
		reconstructed += t.Body + "\n"
	}

	p := path.Join(workDir, seedDir, prismaDir, prismaFile)

	return ioutil.WriteFile(p, []byte(reconstructed), 0644)
}

func writeToGQLGen(workDir string, types []GQLType) error {
	reconstructed := ""
	for _, t := range types {
		reconstructed += t.Body + "\n"
	}

	g := path.Join(workDir, seedDir, gqlDir, gqlGenFile)

	return ioutil.WriteFile(g, []byte(reconstructed), 0644)
}

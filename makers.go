package seed

import (
	"io/ioutil"
	"os"
	"path"
)

func createSeedDir(project *Project) error {
	full := path.Join(project.AbsoluteWorkDir, seedDir)
	err := os.Mkdir(full, os.ModePerm)
	if os.IsExist(err) {
		return nil
	}
	return err
}

func createPrismaDir(project *Project) error {
	full := path.Join(project.AbsoluteWorkDir, seedDir, prismaDir)
	err := os.Mkdir(full, os.ModePerm)
	if os.IsExist(err) {
		return nil
	}
	return err
}

func createGQLDir(project *Project) error {
	full := path.Join(project.AbsoluteWorkDir, seedDir, gqlDir)
	err := os.Mkdir(full, os.ModePerm)
	if os.IsExist(err) {
		return nil
	}
	return err
}

func writeToPrisma(project *Project, types []GQLType) error {
	reconstructed := ""
	for _, t := range types {
		reconstructed += t.Body + "\n"
	}

	p := path.Join(project.AbsoluteWorkDir, seedDir, prismaDir, prismaFile)

	return ioutil.WriteFile(p, []byte(reconstructed), 0644)
}

func writeToGQLGen(project *Project, types []GQLType) error {
	reconstructed := ""
	for _, t := range types {
		reconstructed += t.Body + "\n"
	}

	g := path.Join(project.AbsoluteWorkDir, seedDir, gqlDir, gqlGenFile)

	return ioutil.WriteFile(g, []byte(reconstructed), 0644)
}

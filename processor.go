package seed

import (
	"io/ioutil"
	"path"

	"log"

	"os"

	"os/exec"

	"strings"

	"gopkg.in/yaml.v2"
)

type pGenerator struct {
	Generator string `yaml:"generator"`
	Output    string `yaml:"output"`
}

type prismaProcessor struct {
	Endpoint  string       `yaml:"endpoint"`
	Secret    *string      `yaml:"secret,omitempty"`
	Datamodel []string     `yaml:"datamodel"`
	Generate  []pGenerator `yaml:"generate"`
}

func (processor *prismaProcessor) writeConfig(project *Project) error {
	data, err := yaml.Marshal(processor)
	if err != nil {
		return err
	}

	finalPath := path.Join(project.AbsoluteWorkDir, seedDir, prismaDir, prismaConfigFile)

	return ioutil.WriteFile(finalPath, data, 0644)
}

func (processor *prismaProcessor) generate(project *Project) error {
	if err := os.Chdir(path.Join(project.AbsoluteWorkDir, seedDir, prismaDir)); err != nil {
		return err
	}

	if _, err := exec.Command("prisma", "generate").Output(); err != nil {
		return err
	}

	return os.Chdir(project.AbsoluteWorkDir)
}

type gqlGenField struct {
	Filename string  `yaml:"filename"`
	Type     *string `yaml:"type,omitempty"`
}

type gqlGenModel struct {
	Model string `yaml:"model"`
}

type gqlGenProcessor struct {
	Schema   []string               `yaml:"schema"`
	Exec     gqlGenField            `yaml:"exec"`
	Model    gqlGenField            `yaml:"model"`
	Resolver gqlGenField            `yaml:"resolver"`
	Models   map[string]gqlGenModel `yaml:"models"`
}

func (processor *gqlGenProcessor) processAndWrite(project *Project, types []GQLType) error {
	reconstructed := ""
	for _, t := range types {
		tx, err := fixGQLTypeDeclarationToGQLGen(t)
		if err != nil {
			return err
		}
		reconstructed += tx.Body + "\n"
	}

	g := path.Join(project.AbsoluteWorkDir, seedDir, gqlDir, gqlLiteralsGenFile)

	err := ioutil.WriteFile(g, []byte(reconstructed), 0644)
	if err != nil {
		return err
	}

	finalPath := path.Join(project.AbsoluteWorkDir, seedDir, gqlDir, gqlConfigFile)
	for _, t := range types {
		processor.Models[t.Name] = gqlGenModel{
			Model: path.Join(project.Name, "seed", "client."+t.Name),
		}
	}

	data, err := yaml.Marshal(processor)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(finalPath, data, 0644)
}

func (processor *gqlGenProcessor) generateCode(project *Project) error {
	if err := os.Chdir(path.Join(project.AbsoluteWorkDir, seedDir, gqlDir)); err != nil {
		return err
	}

	data, err := exec.Command("go", "run", "github.com/99designs/gqlgen").Output()
	if err != nil {
		log.Println(string(data))
		return err
	}

	return os.Chdir(project.AbsoluteWorkDir)
}

func (processor *gqlGenProcessor) fixResolver(project *Project) error {
	filenameRes := path.Join(project.AbsoluteWorkDir, seedDir, gqlDir, resolverTempFilename)

	data, err := ioutil.ReadFile(filenameRes)
	if err != nil {
		return err
	}

	newData, err := implementTheUnimplemented(string(data))
	if err != nil {
		return err
	}

	newData, err = removeGenesisResolver(newData)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(filenameRes, []byte(newData), 0644); err != nil {
		return err
	}

	newResolverPath := path.Join(project.AbsoluteWorkDir, seedDir, gqlDir, resolverFilename)

	packageImport := path.Join(project.Name, seedDir, "client")

	newResolver := strings.Replace(resolverFile, "{{Package}}", packageImport, 1)

	return ioutil.WriteFile(newResolverPath, []byte(newResolver), 0644)
}

func (processor *gqlGenProcessor) cleanUnimplemented(project *Project) error {
	filenameRes := path.Join(project.AbsoluteWorkDir, seedDir, gqlDir, resolverTempFilename)

	data, err := ioutil.ReadFile(filenameRes)
	if err != nil {
		return err
	}

	newData, unresolved, err := getUnimplementedMethods(string(data))
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(filenameRes, []byte(newData), 0644); err != nil {
		return err
	}

	packageImport := path.Join(project.Name, seedDir, "client")

	toImplementFIle := strings.Replace(toImplementFile, "{{Package}}", packageImport, 1)

	for _, r := range unresolved {
		toImplementFIle += r + "\n"
	}

	completeToImplementFilename := path.Join(project.AbsoluteWorkDir, seedDir, gqlDir, toImplementFilename)

	if err = ioutil.WriteFile(completeToImplementFilename, []byte(toImplementFIle), 0644); err != nil {
		return err
	}

	return nil
}

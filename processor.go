package seed

import (
	"io/ioutil"
	"path"

	"log"

	"encoding/json"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
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

func (processor *prismaProcessor) writeConfig(workDir string) error {
	data, err := yaml.Marshal(processor)
	if err != nil {
		return err
	}

	finalPath := path.Join(workDir, seedDir, prismaDir, prismaConfigFile)

	return ioutil.WriteFile(finalPath, data, 0644)
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

func (processor *gqlGenProcessor) processAndWrite(workDir string, projectName string, types []GQLType) error {
	reconstructed := ""
	for _, t := range types {
		tx, err := fixGQLTypeDeclarationToGQLGen(t)
		if err != nil {
			return err
		}
		reconstructed += tx.Body + "\n"
	}

	g := path.Join(workDir, seedDir, gqlDir, gqlLiteralsGenFile)

	err := ioutil.WriteFile(g, []byte(reconstructed), 0644)
	if err != nil {
		return err
	}

	finalPath := path.Join(workDir, seedDir, gqlDir, gqlConfigFile)
	for _, t := range types {
		processor.Models[t.Name] = gqlGenModel{
			Model: path.Join(projectName, "seed", "client."+t.Name),
		}
	}

	data, err := yaml.Marshal(processor)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(finalPath, data, 0644)
}

func (processor *gqlGenProcessor) generateCode(workDir string) error {
	configFilename := path.Join(workDir, seedDir, gqlDir, gqlConfigFile)
	var cfg *config.Config
	var err error
	cfg, err = config.LoadConfig(configFilename)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "  ")
	e.Encode(cfg)

	if err = api.Generate(cfg); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

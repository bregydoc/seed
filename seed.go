package seed

import (
	"io/ioutil"
	"strings"
)

type GQLType struct {
	Name string
	Body string // []byte
}

func readSeed(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func segment(project *Project) ([]GQLType, error) {
	types := make([]GQLType, 0)

	seed, err := readSeed(project.SeedFilename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(seed, "\n")

	chunk := ""
	cur := 0
	l := lines[cur]
	for cur < len(lines) {
		l = lines[cur]
		if (strings.Contains(l, "type") || strings.Contains(l, "input") || strings.Contains(l, "enum")) && strings.Contains(l, "{") {
			chunk = ""
			name := strings.ReplaceAll(l, "{", "")
			name = strings.ReplaceAll(name, "type", "")
			name = strings.ReplaceAll(name, "input", "")
			name = strings.ReplaceAll(name, "enum", "")
			name = strings.TrimSpace(name)
			for !strings.Contains(l, "}") && cur < len(lines) {
				l = lines[cur]
				chunk += l + "\n"
				cur++
			}
			types = append(types, GQLType{
				Name: name,
				Body: chunk,
			})
		}
		cur++
	}

	return types, nil
}

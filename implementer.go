package seed

import (
	"bytes"
	"regexp"
	"strings"

	"errors"
)

var isFuncRegex = regexp.MustCompile(`func \(.*(\{|\))`)
var resolverRegex = regexp.MustCompile(`\(r \*.+Resolver\)`)
var methodNameRegex = regexp.MustCompile(`[A-Z]\w+\(`)
var typeNameRegex = regexp.MustCompile(`\*prisma\.\w+\)`)
var returnTypeRegex = regexp.MustCompile(`\((\[\])*\*prisma\.\w+`)

func implementTheUnimplemented(toImplement string) (string, error) {
	lines := strings.Split(toImplement, "\n")
	var err error
	buf := bytes.NewBufferString("")
	for i, line := range lines {
		if strings.Contains(line, notImplementedPanicLine) {
			beforeLine := lines[i-1]
			if !isFuncRegex.MatchString(beforeLine) {
				return "", errors.New("your resolver is malformed")
			}

			resRaw := resolverRegex.FindString(beforeLine)

			// log.Println("resRaw=>", resRaw)
			if strings.Contains(resRaw, "query") || strings.Contains(resRaw, "mutation") || strings.Contains(resRaw, "subscription") {
				if _, err = buf.WriteString(line + "\n"); err != nil {
					return "", err
				}
				continue
			}

			methodNameRaw := methodNameRegex.FindString(beforeLine)
			methodName := strings.ReplaceAll(methodNameRaw, " ", "")
			methodName = strings.ReplaceAll(methodName, "(", "")

			// log.Println("methodNameRaw=>", methodNameRaw)
			// log.Println("methodName=>", methodName)

			typeNameRaw := typeNameRegex.FindString(beforeLine)
			typeName := strings.ReplaceAll(typeNameRaw, " ", "")
			typeName = strings.ReplaceAll(typeName, "*prisma", "")
			typeName = strings.ReplaceAll(typeName, ".", "")
			typeName = strings.ReplaceAll(typeName, ")", "")

			// log.Println("typeNameRaw=>", typeNameRaw)
			// log.Println("typeName=>", typeName)

			returnTypeRaw := returnTypeRegex.FindString(beforeLine)
			returnTypeIsArray := strings.Contains(returnTypeRaw, "[]")
			returnType := strings.ReplaceAll(returnTypeRaw, "([]*prisma.", "")

			// log.Println("returnTypeRaw=>", returnTypeRaw)
			// log.Println("returnTypeIsArray=>", returnTypeIsArray)
			// log.Println("returnType=>", returnType)

			fetchLine := ""
			if returnTypeIsArray {
				fetchLine = strings.ReplaceAll(arrayFetch, "{{Parent}}", typeName)
				fetchLine = strings.ReplaceAll(fetchLine, "{{Branch}}", methodName)
				fetchLine = strings.ReplaceAll(fetchLine, "{{BranchType}}", returnType)
			} else {
				fetchLine = strings.ReplaceAll(singleFetch, "{{Parent}}", typeName)
				fetchLine = strings.ReplaceAll(fetchLine, "{{Branch}}", methodName)
			}

			if _, err = buf.WriteString(fetchLine + "\n"); err != nil {
				return "", err
			}
		} else {
			if _, err = buf.WriteString(line + "\n"); err != nil {
				return "", err
			}
		}

	}

	return buf.String(), nil
}

func removeGenesisResolver(in string) (string, error) {
	new := strings.Replace(in, `type Resolver struct{}`, "", 1)
	return new, nil
}

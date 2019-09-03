package seed

import (
	"regexp"
	"strings"
)

var decoratorReg = regexp.MustCompile(`@\w*\(*.*\)*`)

func removeDecorators(in string) string {
	return decoratorReg.ReplaceAllString(in, "")
}

// improve this funtion to understand when a DateTime is a name and not a type
func fixDateTimeToNative(in string) string {
	return strings.ReplaceAll(in, "DateTime", "Time")
}

func fixGQLTypeDeclarationToGQLGen(t GQLType) (GQLType, error) {
	newBody := t.Body
	newBody = removeDecorators(newBody)
	newBody = fixDateTimeToNative(newBody)

	return GQLType{
		Name: t.Name,
		Body: newBody,
	}, nil
}

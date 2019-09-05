package seed

import (
	"bytes"
	"strings"
)

func getUnimplementedMethods(in string) (string, []string, error) {
	lines := strings.Split(in, "\n")

	buf := bytes.NewBufferString("")

	unimplementedMethods := []string{}
	i := 0

	for i < len(lines) {
		l := lines[i]
		if strings.Contains(l, notImplementedPanicLine) {

			j := i
			extracted := false
			internalBuf := bytes.NewBufferString("")
			isUp := true
			totalLines := 0
			for !extracted && j < len(lines) && j >= 0 {

				if isUp {
					j--
					if strings.Contains(lines[j], "func") && strings.Contains(lines[j], "{") {
						buf = bytes.NewBufferString(strings.Replace(buf.String(), lines[j], "", 1))
						isUp = false
					}
				} else {
					if strings.HasPrefix(lines[j], "}") { // Prefix cause I want a first level bracket ('}')
						extracted = true
					}
					if _, err := internalBuf.WriteString(lines[j] + "\n"); err != nil {
						return "", nil, err
					}
					j++
					totalLines++
				}
			}
			i += totalLines - 1 // ! see that
			unimplementedMethods = append(unimplementedMethods, internalBuf.String())
		} else {
			if _, err := buf.WriteString(l + "\n"); err != nil {
				return "", nil, err
			}
		}
		i++
	}

	return buf.String(), unimplementedMethods, nil
}

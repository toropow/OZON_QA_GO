package main

import (
	"strings"
)

func main() {

}

func fixString(line string) (fixLine string) {
	lines := strings.Split(line, ". ")
	var result []string
	for _, oneLine := range lines {
		cleanLine := strings.TrimSpace(oneLine)
		cleanLineUpper := strings.ToUpper(string(cleanLine[0])) + cleanLine[1:]
		result = append(result, cleanLineUpper)
	}

	fixLine = strings.Join(result, ". ")
	if string(fixLine[len(fixLine)-1]) != "." {
		fixLine = fixLine + "."
	}

	return fixLine
}

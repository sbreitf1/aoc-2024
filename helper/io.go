package helper

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadLines(file string) []string {
	data, err := os.ReadFile(file)
	ExitOnError(err)
	lines := strings.Split(string(data), "\n")
	for i := range lines {
		lines[i] = strings.Trim(lines[i], "\r")
	}
	return lines
}

func ReadNonEmptyLines(file string) []string {
	lines := ReadLines(file)
	nonEmptyLines := make([]string, 0, len(lines))
	for _, line := range lines {
		if len(line) > 0 {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}
	return nonEmptyLines
}

func ReadString(file string) string {
	data, err := os.ReadFile(file)
	ExitOnError(err)
	return string(data)
}

func ParseInts(str string) []int {
	pattern := regexp.MustCompile(`\d+`)

	matches := pattern.FindAllString(str, -1)
	ints := make([]int, len(matches))
	for i := range matches {
		ints[i], _ = strconv.Atoi(matches[i])
	}
	return ints
}

func SplitAndTrim(str string, separator string) []string {
	parts := strings.Split(str, separator)
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

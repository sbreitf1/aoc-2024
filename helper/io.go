package helper

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadString(file string) string {
	data, err := os.ReadFile(file)
	ExitOnError(err)
	return string(data)
}

func ReadLines(file string) []string {
	lines := strings.Split(ReadString(file), "\n")
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

func ExtractInts[T Integer](str string) []T {
	pattern := regexp.MustCompile(`-?\d+`)

	matches := pattern.FindAllString(str, -1)
	ints := make([]T, len(matches))
	for i := range matches {
		val, _ := strconv.ParseInt(matches[i], 10, 64)
		ints[i] = T(val)
	}
	return ints
}

func ExtractPositiveInts[T Integer](str string) []T {
	pattern := regexp.MustCompile(`\d+`)

	matches := pattern.FindAllString(str, -1)
	ints := make([]T, len(matches))
	for i := range matches {
		val, _ := strconv.ParseInt(matches[i], 10, 64)
		ints[i] = T(val)
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

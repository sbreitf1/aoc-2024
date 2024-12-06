package helper

import (
	"sort"
	"strings"
)

func GetReversedSlice[T any](arr []T) []T {
	arr2 := make([]T, len(arr))
	copy(arr2, arr)
	ReverseSlice(arr2)
	return arr2
}

func ReverseSlice[T any](arr []T) {
	for i := 0; i < len(arr)/2; i++ {
		tmp := arr[i]
		arr[i] = arr[len(arr)-i-1]
		arr[len(arr)-i-1] = tmp
	}
}

func RemoveIndex[T any](src []T, removeIndex int) []T {
	dst := make([]T, len(src)-1)
	copy(dst[:removeIndex], src[:removeIndex])
	copy(dst[removeIndex:], src[removeIndex+1:])
	return dst
}

func CloneMap[K comparable, V any](src map[K]V) map[K]V {
	dst := make(map[K]V, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func IterateMapInKeyOrder[K Ordered, V any](m map[K]V, f func(k K, v V)) {
	keys := GetKeySlice(m)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, k := range keys {
		f(k, m[k])
	}
}

func GetKeySlice[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func LinesToRunes(lines []string) [][]rune {
	return MapValues(lines, func(line string) []rune {
		return []rune(line)
	})
}

func RunesToLines(runeLines [][]rune) []string {
	return MapValues(runeLines, func(line []rune) string {
		return string(line)
	})
}

func TrimSpaces(lines []string) []string {
	return MapValues(lines, strings.TrimSpace)
}

func MapValues[IN, OUT any](values []IN, mapFunc func(IN) OUT) []OUT {
	out := make([]OUT, len(values))
	for i := range values {
		out[i] = mapFunc(values[i])
	}
	return out
}

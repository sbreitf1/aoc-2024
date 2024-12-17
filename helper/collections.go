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

func ReverseString(str string) string {
	return string(GetReversedSlice([]rune(str)))
}

func RemoveIndex[T any](src []T, removeIndex int) []T {
	dst := make([]T, len(src)-1)
	copy(dst[:removeIndex], src[:removeIndex])
	copy(dst[removeIndex:], src[removeIndex+1:])
	return dst
}

func InitSlice[T any](num int, val T) []T {
	arr := make([]T, num)
	for i := range arr {
		arr[i] = Clone(val)
	}
	return arr
}

func InitSlice2D[T any](w, h int, val T) [][]T {
	arr := make([][]T, h)
	for y := range arr {
		arr[y] = InitSlice(w, val)
	}
	return arr
}

func Combine[T any](arr []T, add ...T) []T {
	dst := make([]T, len(arr)+len(add))
	copy(dst[:len(arr)], arr)
	copy(dst[len(arr):], add)
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

func GetFirstMapKey[K comparable, V any](m map[K]V) K {
	for k := range m {
		return k
	}
	panic("map is empty")
}

func GetUniqueValues[T comparable](values []T) []T {
	m := make(map[T]bool)
	for _, v := range values {
		m[v] = true
	}
	out := make([]T, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

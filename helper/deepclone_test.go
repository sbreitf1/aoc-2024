package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeepCloneOrdinals(t *testing.T) {
	require.Equal(t, 5, Clone(5))
	require.Equal(t, 32.5, Clone(32.5))
	require.Equal(t, true, Clone(true))
	require.Equal(t, "täst", Clone("täst"))
	require.Equal(t, 'b', Clone('b'))
}

func TestDeepCloneStruct(t *testing.T) {
	type Struct struct {
		Int     int
		Float64 float64
		Bool    bool
		String  string
		Rune    rune
		private string
	}

	require.Equal(t, Struct{Int: 5, Float64: 32.5, Bool: true, String: "täst", Rune: 'b', private: ""}, Clone(Struct{Int: 5, Float64: 32.5, Bool: true, String: "täst", Rune: 'b', private: "not copied"}))
}

func TestDeepCloneArray(t *testing.T) {
	original := [2]int{1, 2}
	require.Panics(t, func() { Clone(original) })
}

func TestDeepCloneSlice(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	cloned := Clone(original)
	require.Equal(t, []int{1, 2, 3, 4, 5}, cloned)
	cloned[3] = 42
	require.Equal(t, []int{1, 2, 3, 42, 5}, cloned)
	require.Equal(t, []int{1, 2, 3, 4, 5}, original)
}

func TestDeepCloneSlice2D(t *testing.T) {
	original := [][]int{{1, 2}, {3, 4}, {5, 6}}
	cloned := Clone(original)
	require.Equal(t, [][]int{{1, 2}, {3, 4}, {5, 6}}, cloned)
	cloned[1][1] = 42
	require.Equal(t, [][]int{{1, 2}, {3, 42}, {5, 6}}, cloned)
	require.Equal(t, [][]int{{1, 2}, {3, 4}, {5, 6}}, original)
	cloned[1] = []int{13, 37}
	require.Equal(t, [][]int{{1, 2}, {13, 37}, {5, 6}}, cloned)
	require.Equal(t, [][]int{{1, 2}, {3, 4}, {5, 6}}, original)
}

func TestDeepCloneMap(t *testing.T) {
	original := map[int]string{1: "one", 2: "two", 3: "three"}
	cloned := Clone(original)
	require.Equal(t, map[int]string{1: "one", 2: "two", 3: "three"}, cloned)
	cloned[3] = "drei"
	require.Equal(t, map[int]string{1: "one", 2: "two", 3: "drei"}, cloned)
	require.Equal(t, map[int]string{1: "one", 2: "two", 3: "three"}, original)
}

func TestDeepCloneMapOfSlices(t *testing.T) {
	original := map[int][]int{1: {2, 3}, 4: {5, 6}, 7: {8, 9}}
	cloned := Clone(original)
	require.Equal(t, map[int][]int{1: {2, 3}, 4: {5, 6}, 7: {8, 9}}, cloned)
	cloned[4][1] = 42
	require.Equal(t, map[int][]int{1: {2, 3}, 4: {5, 42}, 7: {8, 9}}, cloned)
	require.Equal(t, map[int][]int{1: {2, 3}, 4: {5, 6}, 7: {8, 9}}, original)
	cloned[4] = []int{13, 37}
	require.Equal(t, map[int][]int{1: {2, 3}, 4: {13, 37}, 7: {8, 9}}, cloned)
	require.Equal(t, map[int][]int{1: {2, 3}, 4: {5, 6}, 7: {8, 9}}, original)
}

func TestDeepCloneMapOfMaps(t *testing.T) {
	original := map[int]map[int]string{1: {2: "zwei"}, 4: {5: "fünf"}, 7: {8: "acht"}}
	cloned := Clone(original)
	require.Equal(t, map[int]map[int]string{1: {2: "zwei"}, 4: {5: "fünf"}, 7: {8: "acht"}}, cloned)
	cloned[4][5] = "foobar"
	require.Equal(t, map[int]map[int]string{1: {2: "zwei"}, 4: {5: "foobar"}, 7: {8: "acht"}}, cloned)
	require.Equal(t, map[int]map[int]string{1: {2: "zwei"}, 4: {5: "fünf"}, 7: {8: "acht"}}, original)
	cloned[4] = map[int]string{7: "sieben"}
	require.Equal(t, map[int]map[int]string{1: {2: "zwei"}, 4: {7: "sieben"}, 7: {8: "acht"}}, cloned)
	require.Equal(t, map[int]map[int]string{1: {2: "zwei"}, 4: {5: "fünf"}, 7: {8: "acht"}}, original)
}

func TestDeepClonePointerToInt(t *testing.T) {
	originalValue := 5
	original := &originalValue
	cloned := Clone(original)
	require.Equal(t, 5, *cloned)
	*cloned = 42
	require.Equal(t, 42, *cloned)
	require.Equal(t, 5, *original)
	require.Equal(t, 5, originalValue)
}

func TestDeepClonePointerToStruct(t *testing.T) {
	type Struct struct {
		Int   int
		Slice []string
	}
	originalValue := Struct{Int: 42, Slice: []string{"foo", "bar"}}
	original := &originalValue
	cloned := Clone(original)
	require.Equal(t, Struct{Int: 42, Slice: []string{"foo", "bar"}}, *cloned)
	cloned.Int = 1337
	require.Equal(t, Struct{Int: 1337, Slice: []string{"foo", "bar"}}, *cloned)
	require.Equal(t, Struct{Int: 42, Slice: []string{"foo", "bar"}}, *original)
	require.Equal(t, Struct{Int: 42, Slice: []string{"foo", "bar"}}, originalValue)
	cloned.Slice[0] = "test"
	require.Equal(t, Struct{Int: 1337, Slice: []string{"test", "bar"}}, *cloned)
	require.Equal(t, Struct{Int: 42, Slice: []string{"foo", "bar"}}, *original)
	require.Equal(t, Struct{Int: 42, Slice: []string{"foo", "bar"}}, originalValue)
}

func TestDeepCloneNested(t *testing.T) {
	type LeafStruct struct {
		Int   int
		Slice []string
	}
	type RootStruct struct {
		SliceOfInts  []int
		SliceOfSlice [][]string
		MapOfSlices  map[string][]int
		Leaf         LeafStruct
		SliceOfLeafs []LeafStruct
	}

	original := RootStruct{SliceOfInts: []int{1, 2, 3}, SliceOfSlice: [][]string{{"eins"}, {"zwei", "drei"}}, MapOfSlices: map[string][]int{"one": {1}, "two": {3, 4}}, Leaf: LeafStruct{Int: 42, Slice: []string{"oans", "zwoa"}}, SliceOfLeafs: []LeafStruct{{Int: 13, Slice: []string{"foobar"}}, {Int: 37, Slice: []string{}}}}
	cloned := Clone(original)
	require.Equal(t, RootStruct{SliceOfInts: []int{1, 2, 3}, SliceOfSlice: [][]string{{"eins"}, {"zwei", "drei"}}, MapOfSlices: map[string][]int{"one": {1}, "two": {3, 4}}, Leaf: LeafStruct{Int: 42, Slice: []string{"oans", "zwoa"}}, SliceOfLeafs: []LeafStruct{{Int: 13, Slice: []string{"foobar"}}, {Int: 37, Slice: []string{}}}}, cloned)
	cloned.SliceOfInts[2] = 42
	require.Equal(t, RootStruct{SliceOfInts: []int{1, 2, 42}, SliceOfSlice: [][]string{{"eins"}, {"zwei", "drei"}}, MapOfSlices: map[string][]int{"one": {1}, "two": {3, 4}}, Leaf: LeafStruct{Int: 42, Slice: []string{"oans", "zwoa"}}, SliceOfLeafs: []LeafStruct{{Int: 13, Slice: []string{"foobar"}}, {Int: 37, Slice: []string{}}}}, cloned)
	require.Equal(t, RootStruct{SliceOfInts: []int{1, 2, 3}, SliceOfSlice: [][]string{{"eins"}, {"zwei", "drei"}}, MapOfSlices: map[string][]int{"one": {1}, "two": {3, 4}}, Leaf: LeafStruct{Int: 42, Slice: []string{"oans", "zwoa"}}, SliceOfLeafs: []LeafStruct{{Int: 13, Slice: []string{"foobar"}}, {Int: 37, Slice: []string{}}}}, original)
	cloned.SliceOfSlice[0][0] = "first"
	require.Equal(t, RootStruct{SliceOfInts: []int{1, 2, 42}, SliceOfSlice: [][]string{{"first"}, {"zwei", "drei"}}, MapOfSlices: map[string][]int{"one": {1}, "two": {3, 4}}, Leaf: LeafStruct{Int: 42, Slice: []string{"oans", "zwoa"}}, SliceOfLeafs: []LeafStruct{{Int: 13, Slice: []string{"foobar"}}, {Int: 37, Slice: []string{}}}}, cloned)
	require.Equal(t, RootStruct{SliceOfInts: []int{1, 2, 3}, SliceOfSlice: [][]string{{"eins"}, {"zwei", "drei"}}, MapOfSlices: map[string][]int{"one": {1}, "two": {3, 4}}, Leaf: LeafStruct{Int: 42, Slice: []string{"oans", "zwoa"}}, SliceOfLeafs: []LeafStruct{{Int: 13, Slice: []string{"foobar"}}, {Int: 37, Slice: []string{}}}}, original)
	cloned.MapOfSlices["two"][1] = 8
	require.Equal(t, RootStruct{SliceOfInts: []int{1, 2, 42}, SliceOfSlice: [][]string{{"first"}, {"zwei", "drei"}}, MapOfSlices: map[string][]int{"one": {1}, "two": {3, 8}}, Leaf: LeafStruct{Int: 42, Slice: []string{"oans", "zwoa"}}, SliceOfLeafs: []LeafStruct{{Int: 13, Slice: []string{"foobar"}}, {Int: 37, Slice: []string{}}}}, cloned)
	require.Equal(t, RootStruct{SliceOfInts: []int{1, 2, 3}, SliceOfSlice: [][]string{{"eins"}, {"zwei", "drei"}}, MapOfSlices: map[string][]int{"one": {1}, "two": {3, 4}}, Leaf: LeafStruct{Int: 42, Slice: []string{"oans", "zwoa"}}, SliceOfLeafs: []LeafStruct{{Int: 13, Slice: []string{"foobar"}}, {Int: 37, Slice: []string{}}}}, original)
	cloned.Leaf.Slice[0] = "EINS!"
	require.Equal(t, RootStruct{SliceOfInts: []int{1, 2, 42}, SliceOfSlice: [][]string{{"first"}, {"zwei", "drei"}}, MapOfSlices: map[string][]int{"one": {1}, "two": {3, 8}}, Leaf: LeafStruct{Int: 42, Slice: []string{"EINS!", "zwoa"}}, SliceOfLeafs: []LeafStruct{{Int: 13, Slice: []string{"foobar"}}, {Int: 37, Slice: []string{}}}}, cloned)
	require.Equal(t, RootStruct{SliceOfInts: []int{1, 2, 3}, SliceOfSlice: [][]string{{"eins"}, {"zwei", "drei"}}, MapOfSlices: map[string][]int{"one": {1}, "two": {3, 4}}, Leaf: LeafStruct{Int: 42, Slice: []string{"oans", "zwoa"}}, SliceOfLeafs: []LeafStruct{{Int: 13, Slice: []string{"foobar"}}, {Int: 37, Slice: []string{}}}}, original)
	cloned.SliceOfLeafs[1].Int = 42
	require.Equal(t, RootStruct{SliceOfInts: []int{1, 2, 42}, SliceOfSlice: [][]string{{"first"}, {"zwei", "drei"}}, MapOfSlices: map[string][]int{"one": {1}, "two": {3, 8}}, Leaf: LeafStruct{Int: 42, Slice: []string{"EINS!", "zwoa"}}, SliceOfLeafs: []LeafStruct{{Int: 13, Slice: []string{"foobar"}}, {Int: 42, Slice: []string{}}}}, cloned)
	require.Equal(t, RootStruct{SliceOfInts: []int{1, 2, 3}, SliceOfSlice: [][]string{{"eins"}, {"zwei", "drei"}}, MapOfSlices: map[string][]int{"one": {1}, "two": {3, 4}}, Leaf: LeafStruct{Int: 42, Slice: []string{"oans", "zwoa"}}, SliceOfLeafs: []LeafStruct{{Int: 13, Slice: []string{"foobar"}}, {Int: 37, Slice: []string{}}}}, original)
	cloned.SliceOfLeafs[0].Slice[0] = "blub"
	require.Equal(t, RootStruct{SliceOfInts: []int{1, 2, 42}, SliceOfSlice: [][]string{{"first"}, {"zwei", "drei"}}, MapOfSlices: map[string][]int{"one": {1}, "two": {3, 8}}, Leaf: LeafStruct{Int: 42, Slice: []string{"EINS!", "zwoa"}}, SliceOfLeafs: []LeafStruct{{Int: 13, Slice: []string{"blub"}}, {Int: 42, Slice: []string{}}}}, cloned)
	require.Equal(t, RootStruct{SliceOfInts: []int{1, 2, 3}, SliceOfSlice: [][]string{{"eins"}, {"zwei", "drei"}}, MapOfSlices: map[string][]int{"one": {1}, "two": {3, 4}}, Leaf: LeafStruct{Int: 42, Slice: []string{"oans", "zwoa"}}, SliceOfLeafs: []LeafStruct{{Int: 13, Slice: []string{"foobar"}}, {Int: 37, Slice: []string{}}}}, original)
}

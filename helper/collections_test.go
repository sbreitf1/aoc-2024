package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetReversedSlice(t *testing.T) {
	t.Run("Ints0", func(t *testing.T) {
		in := []int{}
		require.Equal(t, []int{}, GetReversedSlice(in))
		require.Equal(t, []int{}, in)
	})
	t.Run("Ints1", func(t *testing.T) {
		in := []int{1}
		require.Equal(t, []int{1}, GetReversedSlice(in))
		require.Equal(t, []int{1}, in)
	})
	t.Run("Ints5", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		require.Equal(t, []int{5, 4, 3, 2, 1}, GetReversedSlice(in))
		require.Equal(t, []int{1, 2, 3, 4, 5}, in)
	})
	t.Run("Ints6", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6}
		require.Equal(t, []int{6, 5, 4, 3, 2, 1}, GetReversedSlice(in))
		require.Equal(t, []int{1, 2, 3, 4, 5, 6}, in)
	})
	t.Run("Strings4", func(t *testing.T) {
		in := []string{"foo", "yolo", "bar", "öäü"}
		require.Equal(t, []string{"öäü", "bar", "yolo", "foo"}, GetReversedSlice(in))
		require.Equal(t, []string{"foo", "yolo", "bar", "öäü"}, in)
	})
}

func TestReverseSlice(t *testing.T) {
	t.Run("Ints0", func(t *testing.T) {
		in := []int{}
		ReverseSlice(in)
		require.Equal(t, []int{}, in)
	})
	t.Run("Ints1", func(t *testing.T) {
		in := []int{1}
		ReverseSlice(in)
		require.Equal(t, []int{1}, in)
	})
	t.Run("Ints5", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		ReverseSlice(in)
		require.Equal(t, []int{5, 4, 3, 2, 1}, in)
	})
	t.Run("Ints6", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6}
		ReverseSlice(in)
		require.Equal(t, []int{6, 5, 4, 3, 2, 1}, in)
	})
	t.Run("Strings4", func(t *testing.T) {
		in := []string{"foo", "yolo", "bar", "öäü"}
		ReverseSlice(in)
		require.Equal(t, []string{"öäü", "bar", "yolo", "foo"}, in)
	})
}

func TestReverseString(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		require.Equal(t, "", ReverseString(""))
	})
	t.Run("OneChar", func(t *testing.T) {
		require.Equal(t, "a", ReverseString("a"))
	})
	t.Run("ManyChars", func(t *testing.T) {
		require.Equal(t, "hjs#üä4jh", ReverseString("hj4äü#sjh"))
	})
}

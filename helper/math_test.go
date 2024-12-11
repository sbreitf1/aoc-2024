package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDigits(t *testing.T) {
	require.Equal(t, 1, Digits(1))
	require.Equal(t, 1, Digits(9))
	require.Equal(t, 2, Digits(10))
	require.Equal(t, 2, Digits(99))
	require.Equal(t, 7, Digits(9999999))
	require.Equal(t, 8, Digits(10000000))
	require.Equal(t, 11, Digits(99999999999))
	require.Equal(t, 12, Digits(100000000000))
}

package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadString(t *testing.T) {
	tmpDir := t.TempDir()
	file := filepath.Join(tmpDir, t.Name())
	require.NoError(t, os.WriteFile(file, []byte("some cool stuff"), os.ModePerm))
	require.Equal(t, "some cool stuff", ReadString(file))
}

func TestReadLines(t *testing.T) {
	testCases := []struct {
		FileContent string
		Expected    []string
	}{
		{"", []string{""}},
		{"\n", []string{"", ""}},
		{"foobar", []string{"foobar"}},
		{"foo\nbar", []string{"foo", "bar"}},
		{"\nfoo\n\nbar\n", []string{"", "foo", "", "bar", ""}},
		{"\r", []string{""}},
		{"\rfoo\r", []string{"foo"}},
		{"foo\r\nbar", []string{"foo", "bar"}},
		{"foo\r\n\r\nbar", []string{"foo", "", "bar"}},
	}

	tmpDir := t.TempDir()

	for i, tc := range testCases {
		file := filepath.Join(tmpDir, fmt.Sprintf("%s-%d", t.Name(), i))
		require.NoError(t, os.WriteFile(file, []byte(tc.FileContent), os.ModePerm))
		require.Equal(t, tc.Expected, ReadLines(file))
	}
}

func TestNonEmptyReadLines(t *testing.T) {
	testCases := []struct {
		FileContent string
		Expected    []string
	}{
		{"", []string{}},
		{"\n", []string{}},
		{"foobar", []string{"foobar"}},
		{"foo\nbar", []string{"foo", "bar"}},
		{"\nfoo\n\nbar\n", []string{"foo", "bar"}},
		{"\r", []string{}},
		{"\rfoo\r", []string{"foo"}},
		{"foo\r\nbar", []string{"foo", "bar"}},
		{"foo\r\n\r\nbar", []string{"foo", "bar"}},
	}

	tmpDir := t.TempDir()

	for i, tc := range testCases {
		file := filepath.Join(tmpDir, fmt.Sprintf("%s-%d", t.Name(), i))
		require.NoError(t, os.WriteFile(file, []byte(tc.FileContent), os.ModePerm))
		require.Equal(t, tc.Expected, ReadNonEmptyLines(file))
	}
}

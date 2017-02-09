package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func setup(t *testing.T, size int) (string, string, func()) {
	sourceDir, err := ioutil.TempDir("", "source_test")
	destDir, err := ioutil.TempDir("", "dest_test")
	if err != nil {
		fatal(err)
	}

	createTempFiles(sourceDir, size)

	return sourceDir, destDir, func() {
		os.RemoveAll(sourceDir)
		os.RemoveAll(destDir)
	}
}

func createTempFiles(dir string, size int) {
	content := []byte("temporary file's content")

	for i := 0; i < size; i++ {
		fileName := filepath.Join(dir, fmt.Sprintf("temp_file_%d", i))
		if err := ioutil.WriteFile(fileName, content, 0444); err != nil {
			fatal(err)
		}
	}
}

func check(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestMoveFiles(t *testing.T) {
	total := 100
	toMove := 10

	sourceDir, destDir, teardown := setup(t, total)
	defer teardown()

	err := moveRandomly(sourceDir, destDir, toMove)
	if err != nil {
		fatal(err)
	}

	destFiles, err := listFiles(destDir)
	check(t, err)
	if toMove != len(destFiles) {
		t.Errorf("files note copied to destination")
	}

	sourceFiles, err := listFiles(sourceDir)
	check(t, err)
	if (total - toMove) != len(sourceFiles) {
		t.Errorf("files note moved to source from destination")
	}
}

func TestReturnErrorWhenThereisNoEnoughFilesToMove(t *testing.T) {
	total := 10
	toMove := 100

	sourceDir, destDir, teardown := setup(t, total)
	defer teardown()

	err := moveRandomly(sourceDir, destDir, toMove)
	if err == nil {
		t.Errorf("No enough files expected but no error returned")
	}
}

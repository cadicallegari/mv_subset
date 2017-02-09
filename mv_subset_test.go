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

func countFiles(dir string) int {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fatal(err)
	}

	return len(files)
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

	if toMove != countFiles(destDir) {
		t.Errorf("files note copied to destination")
	}

	if (total - toMove) != countFiles(sourceDir) {
		t.Errorf("files note moved to source from destination")
	}
}

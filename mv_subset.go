package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
)

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("mv_subset source_dir dest_dir amount")
	os.Exit(1)
}

func parseargs() (string, string, int) {

	return "", "", 10
}

func listFiles(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fatal(err)
	}

	return files
}

func moveRandomly(source string, dest string, amount int) error {
	files := listFiles(source)

	if len(files) < amount {
		return errors.New("There is no enough files to move")
	}

	randons := rand.Perm(len(source))

	for i := 0; i < amount; i++ {
		randomFile := files[randons[i]]
		move(randomFile.Name(), source, dest)
	}

	return nil
}

func move(fileName, source, dest string) {
	err := os.Rename(concatPath(source, fileName), concatPath(dest, fileName))

	if err != nil {
		fatal(err)
	}
}

func concatPath(sourcePath, fileName string) string {
	return filepath.Clean(fmt.Sprintf("%s/%s", sourcePath, fileName))
}

func main() {
	source, dest, amount := parseargs()
	err := moveRandomly(source, dest, amount)

	if err != nil {
		fatal(err)
	}
}

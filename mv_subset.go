package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
)

func listFiles(dir string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func moveRandomly(source string, dest string, amount int) error {
	files, err := listFiles(source)

	if err != nil {
		return err
	}

	if len(files) < amount {
		return errors.New("There is no enough files to move")
	}

	randons := rand.Perm(len(source))

	for i := 0; i < amount; i++ {
		randomFile := files[randons[i]]
		err := move(randomFile.Name(), source, dest)
		if err != nil {
			return err
		}
	}

	return nil
}

func move(fileName, source, dest string) error {
	return os.Rename(concatPath(source, fileName), concatPath(dest, fileName))
}

func concatPath(sourcePath, fileName string) string {
	return filepath.Clean(fmt.Sprintf("%s/%s", sourcePath, fileName))
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("\tmv_subset <source_dir> <dest_dir> <amount>\n")
	os.Exit(1)
}

func parseargs() (string, string, int) {
	flag.Parse()

	if len(flag.Args()) < 3 {
		usage()
	}

	amount, err := strconv.Atoi(flag.Args()[2])
	if err != nil {
		fatal(err)
	}

	return flag.Args()[0], flag.Args()[1], amount
}

func main() {
	source, dest, amount := parseargs()
	err := moveRandomly(source, dest, amount)

	if err != nil {
		fatal(err)
	}
}

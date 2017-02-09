package main

import (
	"fmt"
	"os"
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

func moveRandomly(source string, dest string, amount int) error {

	return nil
}

func main() {
	source, dest, amount := parseargs()
	err := moveRandomly(source, dest, amount)

	if err != nil {
		fatal(err)
	}
}

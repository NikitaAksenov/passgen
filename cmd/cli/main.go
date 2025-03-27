package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/NikitaAksenov/passgen/pkg/passgen"
)

func main() {
	length := flag.Int("length", 8, "Length of password (must be in [1; 255] range)")
	flag.Parse()

	password, err := passgen.Generate(*length)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(password)
}

package main

import (
	"fmt"
	"os"

	"github.com/alisdairrankine/gravmet/data"
	yaml "gopkg.in/yaml.v2"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("invalid filename")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	if file == nil {
		fmt.Println("file  empty")
	}
	defer file.Close()

	team := &data.Team{}

	dec := yaml.NewDecoder(file)

	dec.Decode(team)

	fmt.Printf("%+v\n", team)
}

package main

import (
	"fmt"
	project_root_directory "github.com/golang-infrastructure/go-project-root-directory"
)

func main() {

	file, err := project_root_directory.ReadRootFile("go.mod")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(file))
	// Output:
	// module github.com/golang-infrastructure/go-project-root-directory
	//
	// go 1.19
	//
	// require (
	//        github.com/golang-infrastructure/go-how-run v0.0.0-20230107060855-56163adc7748
	//        github.com/stretchr/testify v1.8.1
	// )
	//
	// require (
	//        github.com/davecgh/go-spew v1.1.1 // indirect
	//        github.com/pmezard/go-difflib v1.0.0 // indirect
	//        gopkg.in/yaml.v3 v3.0.1 // indirect
	// )

}

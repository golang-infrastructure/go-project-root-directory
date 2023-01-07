package main

import (
	"fmt"
	project_root_directory "github.com/golang-infrastructure/go-project-root-directory"
)

func main() {

	path, err := project_root_directory.GetRootFilePath("go.mod")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(path)
	// Output:
	// D:\workspace\go-project-root-directory\go.mod

}

package main

import (
	"fmt"
	project_root_directory "github.com/golang-infrastructure/go-project-root-directory"
)

func main() {

	directory, err := project_root_directory.GetRootDirectory()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("项目根路径： " + directory)

}

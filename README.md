# Project Root Directory

# 一、为什么会有这个东西？

在Golang的单元测试`foo_test.go`中读取文件和在`main.go`中读取文件的初始路径根本就不一样，如果你使用的是相对路径读取配置文件之类的，就可能会被恶心到，此项目就是解决类似问题的，提供一个固定的项目根目录的绝对路径，你可以从这里出发去读取你的文件，为从不同地方运行访问文件都提供一致性的行为。

# 二、安装

```bash
go get -u github.com/golang-infrastructure/go-project-root-directory
```

# 三、示例

## 3.1 获取项目的根目录

在main.go中： 

```go
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
	fmt.Println(directory)
	// Output:
	// D:\workspace\go-project-root-directory

}
```

在单元测试中，获取到的路径都是相同的：

```go
package test

import (
	project_root_directory "github.com/golang-infrastructure/go-project-root-directory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_foo(t *testing.T) {
	directory, err := project_root_directory.GetRootDirectory()
	assert.Nil(t, err)
	t.Log(directory)
	// Output: 
	// D:\workspace\go-project-root-directory
}
```

## 3.2 获取项目根路径下的文件的路径 

```go
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
```

## 3.3 读取项目根路径下的文件的内容

```go
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
```






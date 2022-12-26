package project_root_directory

import (
	"os"
	"path/filepath"
)

func Get() string {
	return ""
}

type RunType int

const (
	RunTypeUintTest RunType = iota
	RunTypeRelease
)

func judgeCurrentRunType() {

}

// 获取可执行文件的路径
func getExecutable() string {
	ex, err := os.Executable()
	if err != nil {
		return ""
	}
	exPath := filepath.Dir(ex)
	realPath, err := filepath.EvalSymlinks(exPath)
	if err != nil {
		return ""
	}
	return realPath
	//return filepath.Dir(realPath)
}

// 源代码测试

// 源代码Example

// 源代码基准测试

// 源代码main方法运行

// 编译后的执行

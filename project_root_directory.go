package project_root_directory

import (
	how_run "github.com/golang-infrastructure/go-how-run"
	"os"
	"path/filepath"
)

// GetRootDirectory 获取当前项目的根路径在哪里
func GetRootDirectory() (string, error) {
	runType, err := how_run.GetRunType()
	if err != nil {
		return "", err
	}
	switch runType {
	// RunTypeSourceUnknown 咱也不知道咋运行的
	case how_run.RunTypeUnknown:
		return "", ErrProjectRootDirectoryUnknown
	// RunTypeSourceCode 是从源代码中运行的
	case how_run.RunTypeSourceCode:
		return GetSourceCodeRootDirectory()
	// RunTypeReleaseBinary 发布的二进制文件运行
	case how_run.RunTypeReleaseBinary:
		return GetExecutableRootDirectory()
	default:
		return "", ErrProjectRootDirectoryUnknown
	}
}

// GetSourceCodeRootDirectory 如果是从源代码运行的，则按此方式寻找项目的根目录
func GetSourceCodeRootDirectory() (string, error) {
	searchDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}
	// 从当前路径往上找，第一个拥有go.mod文件的目录认为是项目的根路径
	for searchDirectory != "" {
		goModPath := filepath.Join(searchDirectory, "go.mod")
		stat, err := os.Stat(goModPath)
		if err == nil && stat != nil && !stat.IsDir() {
			return searchDirectory, nil
		}
		searchDirectory = filepath.Dir(searchDirectory)
	}
	return "", ErrProjectRootDirectoryUnknown
}

// GetExecutableRootDirectory 获取可执行文件的root路径
func GetExecutableRootDirectory() (string, error) {
	// 对于可执行文件而言，其所在的路径就是项目的根目录
	return os.Getwd()
}

// GetRootFilePath 返回根路径下的文件路径
func GetRootFilePath(filename string) (string, error) {
	directory, err := GetRootDirectory()
	if err != nil {
		return "", err
	}
	return filepath.Join(directory, filename), nil
}

// ReadRootFile 读取项目根目录下的文件
func ReadRootFile(filename string) ([]byte, error) {
	path, err := GetRootFilePath(filename)
	if err != nil {
		return nil, err
	}
	return os.ReadFile(path)
}

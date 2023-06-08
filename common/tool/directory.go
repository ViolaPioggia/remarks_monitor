package tool

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetFileNum(path string) int64 {
	// 指定文件夹路径
	folderPath := path

	// 遍历文件夹
	fileCount := 0
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 判断是否为文件
		if !info.IsDir() {
			fileCount++
		}

		return nil
	})

	if err != nil {
		fmt.Println("遍历文件夹失败：", err)
		return 0
	}

	return int64(fileCount)
}
func GetFilePaths(path string) []string {
	// 指定文件夹路径
	folderPath := path

	// 打开文件夹
	folder, _ := os.Open(folderPath)
	defer folder.Close()

	// 读取文件列表
	fileNames, err := folder.Readdirnames(-1)
	if err != nil {
		fmt.Println("读取文件列表失败：", err)
		return fileNames
	}

	// 遍历文件列表
	var paths []string
	for _, fileName := range fileNames {
		paths = append(paths, path+fileName)
	}
	return paths
}
func GetWD() string {
	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录失败：", err)
		return ""
	}
	return wd
}

package xfile

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// IsFileOrDirExists 判断文件或文件夹是否存在
func IsFileOrDirExists(src string) bool {
	_, err := os.Stat(src)
	if err != nil {
		return os.IsExist(err)
	}

	return true
}

// IsFile 是否是文件
func IsFile(src string) bool {
	f, err := os.Stat(src)
	if err != nil {
		return false
	}

	return !f.IsDir()
}

// IsDir 是否是目录
func IsDir(src string) bool {
	f, err := os.Stat(src)
	if err != nil {
		return false
	}

	return f.IsDir()
}

// ReadFile 读取文件
func ReadFile(path string) (b []byte, err error) {
	if filepath.IsAbs(path) {
		path = filepath.Clean(path)
	} else {
		path, err = filepath.Abs(path)
		if err != nil {
			return
		}
	}
	b, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}
	return
}

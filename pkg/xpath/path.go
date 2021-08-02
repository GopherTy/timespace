package xpath

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gopherty/timespace/pkg/xfile"
)

// BasePath  获取项目的绝对路径
func BasePath() (basePath string, err error) {
	path, err := exec.LookPath(os.Args[0])
	if err != nil {
		return
	}
	path, err = filepath.Abs(path)
	if err != nil {
		return
	}
	basePath = filepath.Dir(path)
	return
}

// CreatePath 根据指定路径创建，若存在就不做任何操作。
func CreatePath(path string) (err error) {
	if path == "" {
		return
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return
	}

	absPath = filepath.Dir(absPath)
	if xfile.IsFileOrDirExists(absPath) {
		return
	}

	err = os.Mkdir(absPath, os.ModePerm)
	if err != nil {
		return
	}

	return
}

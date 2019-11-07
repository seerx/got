package got

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//GetMyPath 获取可执行程序所在路径
func GetMyPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	// return s
	i := strings.LastIndex(s, string(os.PathSeparator))
	// fmt.Println(i)
	path := string(s[0 : i+1])
	return path
}

//DirecotoryExists 路径是否存在
func DirecotoryExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// JoinPath 计算真实路径，去掉 .. 之类的目录层次
func JoinPath(basePath string, relatePath string) string {
	paths := strings.Split(relatePath, "/")
	return filepath.Join(append([]string{basePath}, paths...)...)
}

//MakeDirecotories 连续创建路径
func MakeDirecotories(path string) bool {
	exists, err := DirecotoryExists(path)
	if err != nil {
		return false
	}

	if !exists {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false
		}
	}

	return true
}

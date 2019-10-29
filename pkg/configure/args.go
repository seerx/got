package configure

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// ParseConfigureFile 解析配置文件路径
func ParseConfigureFile(argName string) (string, error) {
	var file string
	flag.StringVar(&file, argName, "", "请指定配置文件路径")
	flag.Parse()

	if file == "" {
		return file, errors.New("请使用 -c 参数指定配置文件路径")
	}

	st, err := os.Lstat(file)
	if os.IsNotExist(err) {
		return file, fmt.Errorf("配置文件 %s 不存在", file)
	}
	if err != nil {
		return file, fmt.Errorf("配置文件 %s 无法访问: %s", file, err.Error())
	}
	if st.IsDir() {
		return file, fmt.Errorf("配置文件 %s 是目录", file)
	}

	return file, nil
}

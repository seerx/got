package tools

import (
	"net"
	"os"
	"os/exec"
	"strings"
)

/*GetIP 获取本机 IP 地址
 */
func GetIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "error"
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	panic("Unable to determine local IP address (non loopback). Exiting.")
}

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

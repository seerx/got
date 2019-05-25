package got

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//ParseJSONFile 解析 json 文件
func ParseJSONFile(file string, v interface{}) error {
	jsonFile, err := os.Open(file)
	if err != nil {
		return err
	}
	err = json.NewDecoder(jsonFile).Decode(v)
	if err != nil {
		return err
	}
	return nil //json.NewDecoder(jsonFile).Decode(v)
}

//WriteJSONFile 写入 json 文件
func WriteJSONFile(file string, v interface{}) error {
	infoStr, err := json.Marshal(v)
	if err != nil {
		return err
	}
	var infoData = []byte(infoStr)
	err = ioutil.WriteFile(file, infoData, 0666)
	return err
}

package got

import (
	"bytes"
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
	infoData, err := json.Marshal(v)
	if err != nil {
		return err
	}

	var out bytes.Buffer
	err = json.Indent(&out, infoData, "", "\t")
	if err != nil {
		return err
	}

	// out.WriteTo(file)
	err = ioutil.WriteFile(file, out.Bytes(), 0666)
	return err
}

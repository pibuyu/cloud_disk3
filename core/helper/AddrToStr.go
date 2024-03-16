package helper

import (
	"bytes"
	"encoding/json"
	"log"
)

// 将对象的地址转化为对象的string形式返回
func AddToStr(data interface{}) (string, error) {
	//地址->bytes->bytes buffer->string
	b, err := json.Marshal(data)
	if err != nil {
		log.Println("json marshal failed,err" + err.Error())
		return "", err
	}

	buff := new(bytes.Buffer)
	err = json.Indent(buff, b, "", "\t")
	if err != nil {
		log.Println("json indent failed,err" + err.Error())
		return "", err
	}

	return buff.String(), nil
}

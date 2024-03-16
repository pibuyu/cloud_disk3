package helper

import uuid "github.com/google/uuid"

func GetUUID() string {
	res, err := uuid.NewV6()
	if err != nil {
		return "生成uuid出错"
	}
	return res.String()
}

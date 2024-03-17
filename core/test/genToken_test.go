package test

import (
	"cloud_disk3/core/helper"
	"testing"
)

func TestGenToken(t *testing.T) {
	generateToken, _ := helper.GenerateToken(1, "user_1", "hhf")
	print(generateToken)
	userClaim, _ := helper.AnalyseToken(generateToken)
	println(userClaim.Identity)
}

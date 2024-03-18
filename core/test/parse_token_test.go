package test

import (
	"cloud_disk3/core/helper"
	"testing"
)

//var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MCwiSWRlbnRpdHkiOiIiLCJOYW1lIjoiIiwiZXhwIjoxNzExMzI4MDc4fQ.XdHkt4mtMqQJTj80b5e6iOqC6cSqtVjrQA3-uBVq63c"

func TestParseToken(t *testing.T) {
	token, _ := helper.GenerateToken(1, "hhf", "123456")
	userClaim, _ := helper.AnalyseToken(token)
	println(userClaim.Identity)
}

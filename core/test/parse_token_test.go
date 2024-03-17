package test

import (
	"cloud_disk3/core/helper"
	"testing"
)

func TestParseToken(t *testing.T) {
	userClaim, _ := helper.AnalyseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MCwiSWRlbnRpdHkiOiIiLCJOYW1lIjoiIn0.y1PPwP_Zwu2k5YyB3vsh8TM73ouFpNNrFPaSKJHJt6s")
	println(userClaim)
}

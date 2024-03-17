package test

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	println(os.Setenv("TENCENT_SECRET_ID", "AKID2QrIVaVQEwTnt592z3wgRIOTTMbZQ6aF"))
}

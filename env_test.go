package env

import (
	"os"
	"testing"
)

func TestSetEnvVariables(t *testing.T) {
	path := "./test.env"
	SetEnvVariables(path)
	t1 := os.Getenv("TEST1")
	if t1 == "" {
		t.Fatal("SetEnvVariables failed")
	}
	t.Log("SetEnvVariables succeeded")
}

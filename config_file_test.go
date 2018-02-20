package main

import (
	"io/ioutil"
	"os"
	"testing"
)

var filePath = "/tmp/config"
var text = []byte(`some stuff`)

func TestEnsureConfigFile(t *testing.T) {
	ioutil.WriteFile(filePath, text, os.FileMode(int(0664)))
	defer os.Remove(filePath)

	check, err := EnsureConfigFile(filePath)
	if err != nil {
		t.Error("Error looking for file")
	}
	t.Logf("successfully looked for file: %v", check)
}

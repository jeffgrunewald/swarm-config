package main

import (
	"reflect"
	"testing"
)

func TestVersion(t *testing.T) {
	vType := reflect.TypeOf(Version).Kind()
	if vType != reflect.String {
		t.Error("Version is not properly defined")
	}
	t.Logf("successfully validated version is of type: %v", vType)
}

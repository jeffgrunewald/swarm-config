package main

import "testing"

func TestGetUserHome(t *testing.T) {
	usrHome, err := getUserHome()
	if err != nil {
		t.Error("Couldn't retrieve user home directory")
	}
	t.Logf("successfully validated user home dir at: %s", usrHome)
}

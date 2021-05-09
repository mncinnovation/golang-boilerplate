package main

import (
	"testing"
)

func TestGreetingOK(t *testing.T) {
	err := Greeting("Arman")
	if err != nil {
		t.Error(err)
	}
}

func TestGreetingFail(t *testing.T) {
	err := Greeting("")
	if err != nil {
		t.Error(err)
	}
}

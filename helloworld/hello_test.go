package helloworld

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	actual := Hello("hoge")
	expected := "hello world, hoge"
	if actual != expected {
		t.Errorf("actual:%s want:%s", actual, expected)
	}
}

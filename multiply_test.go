package gomodmultiplies_test

import (
	"testing"

	"github.com/bxcodec/gomodmultiplies/v2"
)

func TestMultiply(t *testing.T) {
	a, b := 10, 20
	res := gomodmultiplies.Multiply(a, b)
	if res != 200 {
		t.Errorf("got: %v want:%v \n", res, 200)
	}
}

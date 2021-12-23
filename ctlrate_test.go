package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomWeight(t *testing.T) {

	RandomWeight()
	a := assert.New(t)

	arr := []string{"a", "b", "c", "d", "e"}
	a.Equal(true, containString(arr, "a"), "a is in arr")
	a.Equal(true, containString(arr, "e"), "e is in arr")
	a.Equal(false, containString(arr, "f"), "f is not in arr")
}

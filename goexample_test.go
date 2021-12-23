package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExample(t *testing.T) {

	a := assert.New(t)
	arr := []string{"a", "b", "c", "d", "e"}
	a.Equal(true, containString(arr, "a"), "a is in arr")

}

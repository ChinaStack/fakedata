package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFakeAuth(t *testing.T) {

	events := FakeShoppingBehavior()
	fmt.Println(events[0])
	a := assert.New(t)
	arr := []string{"a", "b", "c", "d", "e"}
	a.Equal(true, containString(arr, "a"), "a is in arr")

}
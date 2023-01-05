package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumar(t *testing.T) {

	c := require.New(t)
	result := Add(20, 2)

	expect := 22
	c.Equal(expect, result)
}

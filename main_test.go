package main

import (
	_ "image/png"

	//"runtime/trace"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHasAMine(t *testing.T) {

	var cell = true
	require.True(t, hasAMine(cell))

	cell = false
	require.False(t, hasAMine(cell))
}

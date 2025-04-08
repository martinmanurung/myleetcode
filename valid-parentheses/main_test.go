package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidBracket(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(isValid("(}"), false)
	assert.Equal(isValid("(}[){]"), false)
	assert.Equal(isValid("(]"), false)
	assert.Equal(isValid("([)]"), false)
	assert.Equal(isValid("{[)}"), false)
}

func TestValidBracket(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(isValid("()"), true)
	assert.Equal(isValid("()[]{}"), true)
	assert.Equal(isValid("{[]}"), true)
}

package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolveText(t *testing.T) {
	assert := assert.New(t)

	s := "$name$ and $name$"
	c := &Content{Text: s, Params: []string{"name"}}

	v := make(map[string]string)
	v["name"] = "test"

	n, err := c.ResolveText(v)
	assert.NoError(err)
	assert.EqualValues("test and test", n)
}

func TestResolveTextNoValue(t *testing.T) {
	assert := assert.New(t)

	s := "$name$ and $name$"
	c := &Content{Text: s, Params: []string{"name"}}

	v := make(map[string]string)

	_, err := c.ResolveText(v)
	assert.Error(err)
}

func TestResolveTextMultiple(t *testing.T) {
	assert := assert.New(t)

	s := "$name$ and $other$"
	c := &Content{Text: s, Params: []string{"name", "other"}}

	v := make(map[string]string)
	v["name"] = "test"
	v["other"] = "second"

	n, err := c.ResolveText(v)
	assert.NoError(err)
	assert.EqualValues("test and second", n)
}

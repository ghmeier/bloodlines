package models

import (
	"testing"

	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
)

func TestToList(t *testing.T) {
	assert := assert.New(t)

	s := "a,b,c"
	l := ToList(s)

	assert.EqualValues("a", l[0])
	assert.EqualValues("b", l[1])
	assert.EqualValues("c", l[2])
	assert.Equal(3, len(l))

	s = ""
	l = ToList(s)
	assert.Equal(0, len(l))
}

func TestToMap(t *testing.T) {
	assert := assert.New(t)

	s := []string{"a:a", "b:b", "c:c"}
	l := toMap(s)

	assert.EqualValues("a", l["a"])
	assert.EqualValues("b", l["b"])
	assert.EqualValues("c", l["c"])
}

func TestToUUIDList(t *testing.T) {
	assert := assert.New(t)

	s := ""
	nids := toUUIDList(s)
	assert.Equal(0, len(nids))

	ids := []uuid.UUID{uuid.NewUUID(), uuid.NewUUID()}
	s = ids[0].String() + "," + ids[1].String()

	nids = toUUIDList(s)

	assert.Equal(2, len(nids))
	assert.Equal(ids[0], nids[0])
	assert.Equal(ids[1], nids[1])
}

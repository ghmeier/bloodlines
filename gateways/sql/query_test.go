package sql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testQuery struct {
	*BaseSearch
	base  string
	one   *SortTerm
	two   *SortTerm
	three *SortTerm
}

func TestBasicQuery(t *testing.T) {
	assert := assert.New(t)

	q1 := newTestQuery(0, 1, 0)
	q2 := newTestQuery(-1, 0, 0)
	q3 := newTestQuery(1, 1, -1)

	s1 := q1.ToQuery()
	s2 := q2.ToQuery()
	s3 := q3.ToQuery()
	e1 := "init_query WHERE one=1 OR two LIKE \"%two%\" OR three LIKE \"%three%\" ORDER BY one ASC, two DESC, three ASC LIMIT ?,?"
	e2 := "init_query WHERE one=1 OR two LIKE \"%two%\" OR three LIKE \"%three%\" ORDER BY two ASC, three ASC LIMIT ?,?"
	e3 := "init_query WHERE one=1 OR two LIKE \"%two%\" OR three LIKE \"%three%\" ORDER BY one DESC, two DESC LIMIT ?,?"

	assert.EqualValues(e1, s1)
	assert.EqualValues(e2, s2)
	assert.EqualValues(e3, s3)
}

func newTestQuery(one, two, three int) Search {
	return &testQuery{
		BaseSearch: &BaseSearch{},
		base:       "init_query",
		one:        NewSortTerm("one", "1", one, false),
		two:        NewSortTerm("two", "two", two, true),
		three:      NewSortTerm("three", "three", three, true),
	}
}

func (t *testQuery) ToQuery() string {
	return t.Limit(t.Order(t.Where(t.base, t.one, t.two, t.three), t.one, t.two, t.three))
}

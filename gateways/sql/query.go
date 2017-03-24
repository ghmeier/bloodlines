package sql

import (
	"fmt"
)

/*Search defines an interface for structs that can be used in sql queries*/
type Search interface {
	ToQuery() string
	Where(string, *SortTerm) string
	Order(string, ...*SortTerm) string
	And(string, string) string
	Limit(string) string
}

type BaseSearch struct{}

type SortTerm struct {
	key   string
	value string
	order int
	like  bool
}

func NewSortTerm(key, value string, order int, like bool) *SortTerm {
	if order != 0 && order != 1 {
		order = -1
	}

	return &SortTerm{
		key:   key,
		value: value,
		order: order,
		like:  like,
	}
}

func (*BaseSearch) Where(q string, sort *SortTerm) string {
	if sort.key == "" || sort.value == "" {
		return q
	}

	if sort.like {
		return fmt.Sprintf("%s WHERE %s LIKE '%s'", q, sort.key, sort.value)
	}

	return fmt.Sprintf("%s WHERE %s=%s", q, sort.key, sort.value)
}

/*Order returns q + "ORDER BY" the sort key and and order based on the int value
  of sort.order.*/
func (b *BaseSearch) Order(q string, sort ...*SortTerm) string {
	prefix := false

	for _, s := range sort {
		if s.order < 0 {
			continue
		}

		if !prefix {
			q = fmt.Sprintf("%s ORDER BY %s %s", q, s.key, orderString(s.order))
			prefix = true
		} else {
			q = fmt.Sprintf("%s, %s %s", q, s.key, orderString(s.order))
		}
	}

	return q
}

func orderString(i int) string {
	switch i {
	case 0:
		return "ASC"
	case 1:
		return "DESC"
	default:
		return "ASC"
	}
}

func (*BaseSearch) And(q, s string) string {
	if s == "" {
		return q
	}

	return fmt.Sprintf("%s AND %s", q, s)
}

func (*BaseSearch) Limit(q string) string {
	return fmt.Sprintf("%s LIMIT ?,?", q)
}

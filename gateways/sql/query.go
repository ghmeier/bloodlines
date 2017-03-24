package sql

import (
	"fmt"
)

/*Search defines an interface for structs that can be used in sql queries*/
type Search interface {
	ToQuery() string
	Where(string, ...*SortTerm) string
	Order(string, ...*SortTerm) string
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

func (*BaseSearch) Where(q string, sort ...*SortTerm) string {
	prefix := false

	for _, s := range sort {
		if s.key == "" || s.value == "" || s.order < 0 {
			continue
		}

		if !prefix {
			q = fmt.Sprintf("%s WHERE", q)
			prefix = true
		} else {
			q = fmt.Sprintf("%s OR", q)
		}

		if s.like {
			q = fmt.Sprintf("%s %s LIKE \"%%%s%%\"", q, s.key, s.value)
		} else {
			q = fmt.Sprintf("%s %s=%s", q, s.key, s.value)
		}
	}

	return q
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

func (*BaseSearch) Limit(q string) string {
	return fmt.Sprintf("%s LIMIT ?,?", q)
}

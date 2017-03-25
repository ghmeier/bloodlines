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

/*BaseSearch implements Where, Order and Limit, generically for searches*/
type BaseSearch struct{}

/*SortTerm is used to define search terms as well as the ordering for
  queries using that search*/
type SortTerm struct {
	key   string
	value string
	order int
	like  bool
}

/*NewSortTerm returns a configured SortTerm where any order value other than
0 or 1 will not be used in any queries*/
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

/*Where adds WHERE term=value OR ... to the q string for each sort term
  in the sort slice. If SortTerm.like = true, it will use a `like` clause
  instead of `=`*/
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

/*Limit returns q + "LIMIT ?,?" for paging*/
func (*BaseSearch) Limit(q string) string {
	return fmt.Sprintf("%s LIMIT ?,?", q)
}

func orderString(i int) string {
	switch i {
	case 1:
		return "DESC"
	default:
		return "ASC"
	}
}

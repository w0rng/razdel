package set

import (
	"fmt"
	"strings"
)

type Set[T comparable] map[T]struct{}

func New[T comparable](els []T) Set[T] {
	res := make(Set[T])
	for _, elem := range els {
		res[elem] = struct{}{}
	}
	return res
}

func (s Set[T]) Add(set Set[T]) Set[T] {
	for elem := range set {
		s[elem] = struct{}{}
	}
	return s
}

func (s Set[T]) String() string {
	var elems []string
	for elem := range s {
		elems = append(elems, fmt.Sprintf("\"%v\"", elem))
	}
	return fmt.Sprintf("{%s}", strings.Join(elems, ", "))
}

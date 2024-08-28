package andor

import "github.com/cgxarrie-go/andor.git/elementtype"

type andor[T any] struct {
	element   element[T]
	matchFunc func(item T) bool
}

type element[T any] struct {
	Type     elementtype.Type
	elements []element[T]
	item     T
}

func (ao andor[T]) Match() bool {

	return ao.matchElement(ao.element)

}

func (ao andor[T]) matchElement(e element[T]) bool {

	switch e.Type {
	case elementtype.And:
		return ao.matchAnd(e.elements)
	case elementtype.Or:
		return ao.matchOr(e.elements)
	case elementtype.Item:
		return ao.matchFunc(e.item)
	default:
		return false
	}
}

func (ao andor[T]) matchAnd(elements []element[T]) bool {

	for _, e := range elements {
		if match := ao.matchElement(e); !match {
			return false
		}
	}

	return true

}

func (ao andor[T]) matchOr(elements []element[T]) bool {
	for _, e := range elements {
		if match := ao.matchElement(e); match {
			return true
		}
	}

	return false

}

package andor

import (
	"errors"
	"fmt"

	"github.com/cgxarrie-go/andor/elementtype"
)

type andor[T any] struct {
	element   element
	matchFunc func(item T) bool
}

type element struct {
	elementType elementtype.Type
	elements    []element
	item        any
}

func New[T any](matchFunc func(T) bool, element element) andor[T] {

	return andor[T]{
		element:   element,
		matchFunc: matchFunc,
	}
}

func And(items ...any) element {

	e := element{
		elementType: elementtype.And,
		elements:    []element{},
	}

	if len(items) == 0 {
		return e
	}

	andElements := []element{}

	for _, item := range items {
		switch item.(type) {
		case element:
			andElements = append(andElements, item.(element))
		default:
			i := element{
				elementType: elementtype.Item,
				item:        item,
			}
			andElements = append(andElements, i)
		}
	}

	e.elements = andElements
	return e
}

func Or(items ...any) element {
	e := element{
		elementType: elementtype.Or,
		elements:    []element{},
	}

	if len(items) == 0 {
		return e
	}

	orElements := []element{}

	for _, item := range items {
		switch item.(type) {
		case element:
			orElements = append(orElements, item.(element))
		default:
			i := element{
				elementType: elementtype.Item,
				item:        item,
			}
			orElements = append(orElements, i)
		}
	}

	e.elements = orElements
	return e
}

func Item(item any) element {
	return element{
		elementType: elementtype.Item,
		item:        item,
	}
}

func (ao andor[T]) Match() (bool, error) {

	if err := ao.validate(ao.element); err != nil {
		return false, err
	}

	return ao.matchElement(ao.element), nil

}

func (ao andor[T]) validate(e any) error {

	if e == nil {
		return errors.New("nil element")
	}

	switch e.(type) {
	case element:
		switch e.(element).elementType {
		case elementtype.And:
			for _, elem := range e.(element).elements {
				if err := ao.validate(elem); err != nil {
					return err
				}
			}
			return nil
		case elementtype.Or:
			for _, elem := range e.(element).elements {
				if err := ao.validate(elem); err != nil {
					return err
				}
			}
			return nil
		default:
			return ao.validate(e.(element).item)
		}
	case T:
		return nil
	default:
		return fmt.Errorf("invalid item type. got %T, item %+v", e, e)
	}
}

func (ao andor[T]) matchElement(e element) bool {

	switch e.elementType {
	case elementtype.And:
		return ao.matchAnd(e.elements)
	case elementtype.Or:
		return ao.matchOr(e.elements)
	default:
		i := e.item.(T)
		return ao.matchFunc(i)
	}
}

func (ao andor[T]) matchAnd(elements []element) bool {

	for _, e := range elements {
		if match := ao.matchElement(e); !match {
			return false
		}
	}

	return true

}

func (ao andor[T]) matchOr(elements []element) bool {
	for _, e := range elements {
		if match := ao.matchElement(e); match {
			return true
		}
	}

	return false

}

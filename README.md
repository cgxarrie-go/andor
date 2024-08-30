# andor
Simple module to combine and and or conditions 

## Add AndOr to your project

go get github.com/cgxarrie-go/andor

## How to use AndOr

- Declare a function to validte your items
- Declare your elements to validate as And, Or or item
- Call the MAth function of yout AndOr condition

The response has 2 elements
- boolean:
    * false when 
        * error is not nil
        * the match function returns false for any item
        * the match function return error for any item
    * true in any other case
- error: 
    * nil when all the added elements match the AndOr generic type
    * not nil when any of the added elements does not match the AndOr generic type

### Types of elements
- Item : An element of the type to be validated individually
- And : All elements must match the validation function to return true
- Or : Any of the elements must match the validation function to return true

Any And or Or element can include Items, And element or Or elements

### Example

```Go

import "github.com/cgxarrie-go/andor"

func main() {

    matchFn := func(i int) (bool, error) {
        return i%2 == 0, nil
    }
    
    condition := andor.New[int](matchFn,
        andor.And(
            2,
            andor.Or(
                4,
                5,
                andor.And(
                    6,
                    8,
                    10,
                ),
            ),
        ))

    result, error := condition.Match()

    
```

one line example
```Go

import "github.com/cgxarrie-go/andor"

func main() {

    matchFn := func(i int) bool {
        return i%2 == 0
    }
    
    condition := andor.New[int](matchFn,
        andor.And(2,andor.Or(4,5,andor.And(6,8,10))))

    result, error := condition.Match()
    
```

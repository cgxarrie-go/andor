# andor
Simple module to combine and and or conditions 

## Add AndOr to your project

go get github.com/cgxarrie-go/andor.git

## How to use AndOr

- Declare a function to validte your items
- Declare your elements to validate as And, Or or item

### Types of elements
- Item : An element of the type to be validated individually
- And : All elements must match the validation function to return true
- Or : Any of the elements must match the validation function to return true

Any And or Or element can include Items, And element or Or elements

### Example

```Go

import "github.com/cgxarrie-go/andor"

func main() {

    matchFn := func(i int16) bool {
        return i%2 == 0
    }
    
    condition := andor.New[int16](fn,
        andor.And[int16](
            andor.Item[int16](2),
                andor.Or[int16](
                    andor.Item[int16](4),
                    andor.Item[int16](5),
                    andor.And[int16](
                        andor.Item[int16](6),
                        andor.Item[int16](8),
                        andor.Item[int16](10),
                    ),
                ),
            ))

    result := condition.Match()

}


    
```

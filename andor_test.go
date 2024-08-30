package andor

import (
	"testing"

	"github.com/cgxarrie-go/andor/elementtype"
	"github.com/stretchr/testify/assert"
)

func Test_MatchItem(t *testing.T) {
	tests := []struct {
		name        string
		funcRetrurn bool
		want        bool
	}{
		{
			name:        "expect true",
			funcRetrurn: true,
			want:        true,
		},
		{
			name:        "expect false",
			funcRetrurn: false,
			want:        false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			andor := andor[int]{
				element: element{
					elementType: elementtype.Item,
					item:        1,
				},
				matchFunc: func(item int) bool {
					return test.funcRetrurn
				},
			}

			// Act
			got, err := andor.Match()

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_MatchAnd_AllTrue_ShouldReturnTrue(t *testing.T) {
	// Arrange
	andor := andor[int]{
		matchFunc: func(item int) bool {
			return item%2 == 0
		},
		element: element{
			elementType: elementtype.And,
			elements: []element{
				{
					elementType: elementtype.Item,
					item:        2,
				},
				{
					elementType: elementtype.Item,
					item:        4,
				},
			},
		},
	}

	// Act
	got, err := andor.Match()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, true, got)
}

func Test_MatchAnd_OneFalse_ShouldReturnFalse(t *testing.T) {
	// Arrange
	andor := andor[int]{
		matchFunc: func(item int) bool {
			return item%2 == 0
		},
		element: element{
			elementType: elementtype.And,
			elements: []element{
				{
					elementType: elementtype.Item,
					item:        2,
				},
				{
					elementType: elementtype.Item,
					item:        3,
				},
			},
		},
	}

	// Act
	got, err := andor.Match()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, false, got)
}

func Test_MatchOr_OneTrue_ShouldReturnTrue(t *testing.T) {
	// Arrange
	andor := andor[int]{
		matchFunc: func(item int) bool {
			return item%2 == 0
		},
		element: element{
			elementType: elementtype.Or,
			elements: []element{
				{
					elementType: elementtype.Item,
					item:        1,
				},
				{
					elementType: elementtype.Item,
					item:        2,
				},
			},
		},
	}

	// Act
	got, err := andor.Match()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, true, got)
}

func Test_MatchOr_AllFalse_ShouldReturnFalse(t *testing.T) {
	// Arrange
	andor := andor[int]{
		matchFunc: func(item int) bool {
			return item%2 == 0
		},
		element: element{
			elementType: elementtype.Or,
			elements: []element{
				{
					elementType: elementtype.Item,
					item:        1,
				},
				{
					elementType: elementtype.Item,
					item:        3,
				},
			},
		},
	}

	// Act
	got, err := andor.Match()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, false, got)
}
func Test_Validate_NilElement_ShouldReturnError(t *testing.T) {
	// Arrange
	ao := andor[int]{}

	// Act
	err := ao.validate(nil)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "nil element", err.Error())
}

func Test_Validate_InvalidItemType_ShouldReturnError(t *testing.T) {
	// Arrange
	ao := andor[int]{}
	item := "invalid"

	// Act
	err := ao.validate(item)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "invalid item type. got string, item invalid", err.Error())
}

func Test_Validate_ValidElement_ShouldNotReturnError(t *testing.T) {
	// Arrange
	ao := andor[int]{}
	element := element{
		elementType: elementtype.Item,
		item:        1,
	}

	// Act
	err := ao.validate(element)

	// Assert
	assert.NoError(t, err)
}

func Test_Validate_ValidAndElement_ShouldNotReturnError(t *testing.T) {
	// Arrange
	ao := andor[int]{}
	element := element{
		elementType: elementtype.And,
		elements: []element{
			{
				elementType: elementtype.Item,
				item:        1,
			},
			{
				elementType: elementtype.Item,
				item:        2,
			},
		},
	}

	// Act
	err := ao.validate(element)

	// Assert
	assert.NoError(t, err)
}

func Test_Validate_InvalidAndElement_ShouldNotReturnError(t *testing.T) {
	// Arrange
	ao := andor[int]{}
	element := element{
		elementType: elementtype.And,
		elements: []element{
			{
				elementType: elementtype.Item,
				item:        1,
			},
			{
				elementType: elementtype.Item,
				item:        "invalid",
			},
		},
	}

	// Act
	err := ao.validate(element)

	// Assert
	assert.Error(t, err)
}

func Test_Validate_ValidOrElement_ShouldNotReturnError(t *testing.T) {
	// Arrange
	ao := andor[int]{}
	element := element{
		elementType: elementtype.Or,
		elements: []element{
			{
				elementType: elementtype.Item,
				item:        1,
			},
			{
				elementType: elementtype.Item,
				item:        2,
			},
		},
	}

	// Act
	err := ao.validate(element)

	// Assert
	assert.NoError(t, err)
}

func Test_Validate_InvalidOrElement_ShouldNotReturnError(t *testing.T) {
	// Arrange
	ao := andor[int]{}
	element := element{
		elementType: elementtype.Or,
		elements: []element{
			{
				elementType: elementtype.Item,
				item:        1,
			},
			{
				elementType: elementtype.Item,
				item:        "invalid-element",
			},
		},
	}

	// Act
	err := ao.validate(element)

	// Assert
	assert.Error(t, err)
}

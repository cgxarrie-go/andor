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
			andor := andor[int16]{
				element: element[int16]{
					elementType: elementtype.Item,
					item:        1,
				},
				matchFunc: func(item int16) bool {
					return test.funcRetrurn
				},
			}

			// Act
			got := andor.Match()

			// Assert
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_MatchAnd_AllTrue_ShouldReturnTrue(t *testing.T) {
	// Arrange
	andor := andor[int16]{
		matchFunc: func(item int16) bool {
			return item%2 == 0
		},
		element: element[int16]{
			elementType: elementtype.And,
			elements: []element[int16]{
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
	got := andor.Match()

	// Assert
	assert.Equal(t, true, got)
}

func Test_MatchAnd_OneFalse_ShouldReturnFalse(t *testing.T) {
	// Arrange
	andor := andor[int16]{
		matchFunc: func(item int16) bool {
			return item%2 == 0
		},
		element: element[int16]{
			elementType: elementtype.And,
			elements: []element[int16]{
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
	got := andor.Match()

	// Assert
	assert.Equal(t, false, got)
}

func Test_MatchOr_OneTrue_ShouldReturnTrue(t *testing.T) {
	// Arrange
	andor := andor[int16]{
		matchFunc: func(item int16) bool {
			return item%2 == 0
		},
		element: element[int16]{
			elementType: elementtype.Or,
			elements: []element[int16]{
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
	got := andor.Match()

	// Assert
	assert.Equal(t, true, got)
}

func Test_MatchOr_AllFalse_ShouldReturnFalse(t *testing.T) {
	// Arrange
	andor := andor[int16]{
		matchFunc: func(item int16) bool {
			return item%2 == 0
		},
		element: element[int16]{
			elementType: elementtype.Or,
			elements: []element[int16]{
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
	got := andor.Match()

	// Assert
	assert.Equal(t, false, got)
}

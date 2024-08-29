package andor_test

import (
	"testing"

	"github.com/cgxarrie-go/andor"
	"github.com/stretchr/testify/assert"
)

func Test_And_True(t *testing.T) {
	// Arrange
	fn := func(i int16) bool {
		return i%2 == 0
	}

	rootCondition := andor.New[int16](fn,
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

	// Act
	got := rootCondition.Match()

	// Assert
	assert.True(t, got)
}

func Test_And_False(t *testing.T) {
	// Arrange
	fn := func(i int16) bool {
		return i%2 == 0
	}

	rootCondition := andor.New[int16](fn,
		andor.And[int16](
			andor.Item[int16](2),
			andor.Or[int16](
				andor.Item[int16](3),
				andor.Item[int16](5),
				andor.And[int16](
					andor.Item[int16](1),
					andor.Item[int16](8),
					andor.Item[int16](10),
				),
			),
		))

	// Act
	got := rootCondition.Match()

	// Assert
	assert.False(t, got)
}

func Test_Or_True(t *testing.T) {
	// Arrange
	fn := func(i int16) bool {
		return i%2 == 0
	}

	rootCondition := andor.New[int16](fn,
		andor.Or[int16](
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

	// Act
	got := rootCondition.Match()

	// Assert
	assert.True(t, got)
}

func Test_Or_False(t *testing.T) {
	// Arrange
	fn := func(i int16) bool {
		return i%2 == 0
	}

	rootCondition := andor.New[int16](fn,
		andor.Or[int16](
			andor.Item[int16](1),
			andor.Or[int16](
				andor.Item[int16](1),
				andor.Item[int16](3),
				andor.And[int16](
					andor.Item[int16](5),
					andor.Item[int16](7),
					andor.Item[int16](8),
				),
			),
		))

	// Act
	got := rootCondition.Match()

	// Assert
	assert.False(t, got)
}

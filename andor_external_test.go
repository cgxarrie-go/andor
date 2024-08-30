package andor_test

import (
	"testing"

	"github.com/cgxarrie-go/andor"
	"github.com/stretchr/testify/assert"
)

func Test_And_True(t *testing.T) {
	// Arrange
	fn := func(i int) bool {
		return i%2 == 0
	}

	rootCondition := andor.New[int](fn,
		andor.And(2, andor.Or(4, 5, andor.And(6, 8, 10))))

	// Act
	got, err := rootCondition.Match()

	// Assert
	assert.NoError(t, err)
	assert.True(t, got)
}

func Test_And_False(t *testing.T) {
	// Arrange
	fn := func(i int) bool {
		return i%2 == 0
	}

	rootCondition := andor.New[int](fn,
		andor.And(2, andor.Or(3, 5, andor.And(1, 8, 10))))

	// Act
	got, err := rootCondition.Match()

	// Assert
	assert.NoError(t, err)
	assert.False(t, got)
}

func Test_Or_True(t *testing.T) {
	// Arrange
	fn := func(i int) bool {
		return i%2 == 0
	}

	rootCondition := andor.New[int](fn,
		andor.Or(2, andor.Or(4, 5, andor.And(6, 8, 10))))

	// Act
	got, err := rootCondition.Match()

	// Assert
	assert.NoError(t, err)
	assert.True(t, got)
}

func Test_Or_False(t *testing.T) {
	// Arrange
	fn := func(i int) bool {
		return i%2 == 0
	}

	rootCondition := andor.New[int](fn,
		andor.Or(1, andor.Or(1, 3, andor.And(5, 7, 8))))

	// Act
	got, err := rootCondition.Match()

	// Assert
	assert.NoError(t, err)
	assert.False(t, got)
}

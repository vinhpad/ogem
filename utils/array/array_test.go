package array

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Run("Map numbers to their squares", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Map(input, func(x int) int { return x * x })
		assert.Equal(t, []int{1, 4, 9, 16, 25}, result)
	})
	
	t.Run("Map strings to their lengths", func(t *testing.T) {
		input := []string{"a", "sadsd", "c"}
		result := Map(input, func(x string) int { return len(x) })
		assert.Equal(t, []int{1, 5, 1}, result)
	})

	t.Run("Empty array", func(t *testing.T) {
		input := []int{}
		result := Map(input, func(x int) int { return x * 2 })
		assert.Equal(t, []int{}, result)
	})
}

func TestContains(t *testing.T) {
	t.Run("Contains existing element", func(t *testing.T) {
		assert.True(t, Contains([]int{1, 2, 3, 4, 5}, 3))
	})

	t.Run("Does not contain element", func(t *testing.T) {
		assert.False(t, Contains([]int{1, 2, 3, 4, 5}, 6))
	})

	t.Run("Empty array", func(t *testing.T) {
		assert.False(t, Contains([]int{}, 1))
	})
}

func TestFind(t *testing.T) {
	t.Run("Find existing element", func(t *testing.T) {
		arr := []int{1, 2, 3, 3, 4, 5}
		elem, found := Find(arr, func(x int) bool { return x == 3 })
		assert.Equal(t, 3, elem)
		assert.True(t, found)
	})
	
	t.Run("Find non-existing element", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		elem, found := Find(arr, func(x int) bool { return x == 6 })
		assert.Equal(t, 0, elem)
		assert.False(t, found)
	})

	t.Run("Empty array", func(t *testing.T) {
		arr := []int{}
		elem, found := Find(arr, func(x int) bool { return x == 1 })
		assert.Equal(t, 0, elem)
		assert.False(t, found)
	})

	t.Run("Find element in slice of interfaces with nil", func(t *testing.T) {
		arr := []any{"a", "b", "c"}
		elem, found := Find(arr, func(x any) bool { return x == nil })
		assert.Equal(t, nil, elem)
		assert.False(t, found)
	})
} 
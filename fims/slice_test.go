package fims

import (
	"reflect"
	"strconv"
	"testing"
)

func TestApply(t *testing.T) {
	tests := []struct {
		name       string
		collection []any
		iteratee   func(any) any
		expected   []any
	}{
		{
			name:       "integers to strings",
			collection: []any{1, 2, 3},
			iteratee: func(i any) any {
				return strconv.Itoa(i.(int))
			},
			expected: []any{"1", "2", "3"},
		},
		{
			name:       "empty collection",
			collection: []any{},
			iteratee: func(i any) any {
				return ""
			},
			expected: []any{},
		},
		{
			name:       "strings to their lengths",
			collection: []any{"testing", "hello", "world!"},
			iteratee: func(s any) any {
				return len(s.(string))
			},
			expected: []any{7, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Apply(tt.collection, tt.iteratee)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Map() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	t.Run("sum integers", func(t *testing.T) {
		collection := []int{1, 2, 3, 4}
		iteratee := func(acc int, item int, index int) int {
			return acc + item
		}
		expected := 10
		got := Reduce(collection, iteratee)

		if got != expected {
			t.Errorf("Reduce() = %v, want %v", got, expected)
		}
	})

	t.Run("concatenate strings", func(t *testing.T) {
		collection := []string{"a", "b", "c"}
		iteratee := func(acc string, item string, index int) string {
			return acc + item
		}
		expected := "abc"
		got := Reduce(collection, iteratee)

		if got != expected {
			t.Errorf("Reduce() = %v, want %v", got, expected)
		}
	})

	t.Run("empty collection", func(t *testing.T) {
		collection := []int{}
		iteratee := func(acc int, item int, index int) int {
			return acc + item
		}
		expected := 0
		got := Reduce(collection, iteratee)

		if got != expected {
			t.Errorf("Reduce() = %v, want %v", got, expected)
		}
	})
}

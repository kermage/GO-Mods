package fims

import (
	"reflect"
	"testing"
)

func TestColumns(t *testing.T) {
	type TestStruct struct {
		ID   int
		Name string
	}

	tests := []struct {
		name       string
		collection []TestStruct
		key        string
		iteratee   func(TestStruct) any
		expected   map[any]TestStruct
	}{
		{
			name: "normal case",
			collection: []TestStruct{
				{ID: 1, Name: "one"},
				{ID: 2, Name: "two"},
				{ID: 3, Name: "three"},
			},
			key: "ID",
			iteratee: func(item TestStruct) any {
				return item.ID
			},
			expected: map[any]TestStruct{
				1: {ID: 1, Name: "one"},
				2: {ID: 2, Name: "two"},
				3: {ID: 3, Name: "three"},
			},
		},
		{
			name:       "empty collection",
			collection: []TestStruct{},
			key:        "ID",
			iteratee: func(item TestStruct) any {
				return item.ID
			},
			expected: map[any]TestStruct{},
		},
		{
			name:       "nil collection",
			collection: nil,
			key:        "ID",
			iteratee: func(item TestStruct) any {
				return item.ID
			},
			expected: map[any]TestStruct{},
		},
		{
			name: "duplicate keys",
			collection: []TestStruct{
				{ID: 1, Name: "one"},
				{ID: 2, Name: "two"},
				{ID: 3, Name: "one"},
			},
			key: "Name",
			iteratee: func(item TestStruct) any {
				return item.Name
			},
			expected: map[any]TestStruct{
				"one": {ID: 3, Name: "one"},
				"two": {ID: 2, Name: "two"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Columns(tt.collection, tt.key, tt.iteratee)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Columns() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFlip(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected map[int]string
	}{
		{
			name:     "non-empty map",
			input:    map[string]int{"a": 1, "b": 2},
			expected: map[int]string{1: "a", 2: "b"},
		},
		{
			name:     "map with duplicate values",
			input:    map[string]int{"a": 1, "b": 2, "c": 1},
			expected: map[int]string{1: "c", 2: "b"},
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: map[int]string{},
		},
		{
			name:     "nil map",
			input:    nil,
			expected: map[int]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Flip(tt.input)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Invert() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFilterMap(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]int
		predicate func(key string, value int) bool
		expected  map[string]int
	}{
		{
			name:  "filter even values",
			input: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			predicate: func(key string, value int) bool {
				return value%2 == 0
			},
			expected: map[string]int{"b": 2, "d": 4},
		},
		{
			name:  "filter by key",
			input: map[string]int{"apple": 1, "banana": 2, "avocado": 3},
			predicate: func(key string, value int) bool {
				return key[0] == 'a'
			},
			expected: map[string]int{"apple": 1, "avocado": 3},
		},
		{
			name:      "empty map",
			input:     map[string]int{},
			predicate: func(key string, value int) bool { return true },
			expected:  map[string]int{},
		},
		{
			name:      "nil map",
			input:     nil,
			predicate: func(key string, value int) bool { return true },
			expected:  map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterMap(tt.input, tt.predicate)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("FilterMap() = %v, want %v", got, tt.expected)
			}
		})
	}
}

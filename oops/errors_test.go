package oops

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewError(t *testing.T) {
	tests := []struct {
		data     any
		message  string
		expected string
	}{
		{
			data:     nil,
			message:  "",
			expected: `{"data":{},"message":""}`,
		},
		{
			data:     "demo",
			message:  "string",
			expected: `{"data":{"error":"demo"},"message":"string"}`,
		},
		{
			data:     911,
			message:  "int",
			expected: `{"data":{"error":911},"message":"int"}`,
		},
		{
			data:     1.1,
			message:  "float",
			expected: `{"data":{"error":1.1},"message":"float"}`,
		},
		{
			data:     false,
			message:  "bool",
			expected: `{"data":{"error":false},"message":"bool"}`,
		},
		{
			data:     fmt.Errorf("err msg"),
			message:  "error",
			expected: `{"data":{"error":"err msg"},"message":"error"}`,
		},
		{
			data:     NewError("inner self", testing.T{}),
			message:  "self",
			expected: `{"data":{"data":{"error":{}},"message":"inner self"},"message":"self"}`,
		},
		{
			data: map[string]any{
				"test": "value",
				"error": NewError("inner errors", map[string]any{
					"test2": []string{"value1", "value2"},
					"error2": NewError("inner errors", map[int]any{
						1: nil,
						2: map[any]any{
							3:                    fmt.Errorf("error3"),
							fmt.Errorf("error4"): "value4",
							"test4":              []int{1, 2, 3},
						},
					}),
				}),
			},
			message:  "nested errors",
			expected: `{"data":{"error":{"data":{"error2":{"data":{"1":null,"2":{"3":"error3","error4":"value4","test4":[1,2,3]}},"message":"inner errors"},"test2":["value1","value2"]},"message":"inner errors"},"test":"value"},"message":"nested errors"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			err := NewError(test.message, test.data)
			result, _ := json.Marshal(err)

			if string(result) != test.expected {
				t.Errorf("NewError(%q, %v)\nResult: %v\nWanted: %v", test.message, test.data, string(result), test.expected)
			}
		})
	}
}

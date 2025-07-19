package oops

import (
	"fmt"
)

type Error struct {
	Data    map[string]any `json:"data"`
	Message string         `json:"message"`
}

func NewError(message string, data any) *Error {
	return &Error{
		Data:    errorsData(data),
		Message: message,
	}
}

// implements the error interface by returning the error message.
func (e *Error) Error() string {
	return e.Message
}

// convert various data types into a structured map format for error storage.
// handles nil values, maps, errors, and other types appropriately.
func errorsData(data any) map[string]any {
	switch v := data.(type) {
	case nil:
		return map[string]any{}
	case map[string]any:
		return resolveErrorsData(v)
	case map[int]any:
		return resolveErrorsData(v)
	case map[any]any:
		return resolveErrorsData(v)
	case *Error:
		return map[string]any{
			"message": v.Message,
			"data":    v.Data,
		}
	case error:
		return map[string]any{
			"status": v.Error(),
		}
	default:
		return map[string]any{
			"status": v,
		}
	}
}

// recursively resolves nested map structures into a flat map[string]any format.
// handles maps with different key types (string, int, any) and converts them uniformly.
func resolveErrorsData[T comparable](data map[T]any) map[string]any {
	result := map[string]any{}

	for name, err := range data {
		k := fmt.Sprintf("%v", name)

		switch v := err.(type) {
		case map[string]any:
			result[k] = resolveErrorsData(v)
		case map[int]any:
			result[k] = resolveErrorsData(v)
		case map[any]any:
			result[k] = resolveErrorsData(v)
		default:
			result[k] = resolveErrorItem(err)
		}
	}

	return result
}

// converts individual error items into a consistent format.
// handles Error instances and standard errors by extracting their relevant information.
func resolveErrorItem(err any) any {
	switch v := err.(type) {
	case *Error:
		return map[string]any{
			"message": v.Message,
			"data":    v.Data,
		}
	case error:
		return v.Error()
	}

	return err
}

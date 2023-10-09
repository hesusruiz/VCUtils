package yaml

import (
	"fmt"
	"strconv"
)

// Map returns a map[string]any according to a dotted path or default or map[string]any.
func GetMap(data any, path string, defaults ...map[string]any) map[string]any {
	value, err := get_map_strict(data, path)

	if err == nil {
		return value
	}

	for _, def := range defaults {
		return def
	}
	return map[string]any{}
}

// map_strict returns a map[string]any according to a dotted path.
func get_map_strict(data any, path string) (map[string]any, error) {
	n, err := Get(data, path)
	if err != nil {
		return nil, err
	}
	if value, ok := n.(map[string]any); ok {
		return value, nil
	}
	return nil, typeMismatch("map[string]any", n)
}

// string_strict returns a string according to a dotted path.
func get_string_strict(data any, path string) (string, error) {
	n, err := Get(data, path)
	if err != nil {
		return "", err
	}
	switch n := n.(type) {
	case bool, float64, int:
		return fmt.Sprint(n), nil
	case string:
		return n, nil
	}
	return "", typeMismatch("bool, float64, int or string", n)
}

// String returns a string according to a dotted path or default or "".
func GetString(data any, path string, defaults ...string) string {
	value, err := get_string_strict(data, path)

	if err == nil {
		return value
	}

	for _, def := range defaults {
		return def
	}
	return ""
}

// bool returns a bool according to a dotted path.
func get_bool(data any, path string) (bool, error) {
	n, err := Get(data, path)
	if err != nil {
		return false, err
	}
	switch n := n.(type) {
	case bool:
		return n, nil
	case string:
		return strconv.ParseBool(n)
	}
	return false, typeMismatch("bool or string", n)
}

// Bool returns a bool according to a dotted path or default value or false.
func GetBool(data any, path string, defaults ...bool) bool {
	value, err := get_bool(data, path)

	if err == nil {
		return value
	}

	for _, def := range defaults {
		return def
	}
	return false
}

// float64 returns a float64 according to a dotted path.
func get_float64(data any, path string) (float64, error) {
	n, err := Get(data, path)
	if err != nil {
		return 0, err
	}
	switch n := n.(type) {
	case float64:
		return n, nil
	case int:
		return float64(n), nil
	case string:
		return strconv.ParseFloat(n, 64)
	}
	return 0, typeMismatch("float64, int or string", n)
}

// Float64 returns a float64 according to a dotted path or default value or 0.
func GetFloat64(data any, path string, defaults ...float64) float64 {
	value, err := get_float64(data, path)

	if err == nil {
		return value
	}

	for _, def := range defaults {
		return def
	}
	return float64(0)
}

// int returns an int according to a dotted path.
func get_int(data any, path string) (int, error) {
	n, err := Get(data, path)
	if err != nil {
		return 0, err
	}
	switch n := n.(type) {
	case float64:
		// encoding/json unmarshals numbers into floats, so we compare
		// the string representation to see if we can return an int.
		if i := int(n); fmt.Sprint(i) == fmt.Sprint(n) {
			return i, nil
		} else {
			return 0, fmt.Errorf("value can't be converted to int: %v", n)
		}
	case int:
		return n, nil
	case string:
		if v, err := strconv.ParseInt(n, 10, 0); err == nil {
			return int(v), nil
		} else {
			return 0, err
		}
	}
	return 0, typeMismatch("float64, int or string", n)
}

// Int returns an int according to a dotted path or default value or 0.
func GetInt(data any, path string, defaults ...int) int {
	value, err := get_int(data, path)

	if err == nil {
		return value
	}

	for _, def := range defaults {
		return def
	}
	return 0
}

// list returns a []any according to a dotted path.
func get_list(data any, path string) ([]any, error) {
	n, err := Get(data, path)
	if err != nil {
		return nil, err
	}
	if value, ok := n.([]any); ok {
		return value, nil
	}
	return nil, typeMismatch("[]any", n)
}

// List returns a []any according to a dotted path or defaults or []any.
func GetList(data any, path string, defaults ...[]any) []any {
	value, err := get_list(data, path)

	if err == nil {
		return value
	}

	for _, def := range defaults {
		return def
	}
	return make([]any, 0)
}

// ListString is for the very common case of a list of strings
func GetListString(data any, path string, defaults ...[]string) []string {
	value, err := get_list(data, path)

	if err == nil {
		return ToListString(value)
	}

	for _, def := range defaults {
		return def
	}
	return make([]string, 0)
}

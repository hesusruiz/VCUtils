package yaml

import "fmt"

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

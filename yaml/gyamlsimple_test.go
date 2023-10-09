package yaml

import "testing"

func TestGetString(t *testing.T) {
	tests := []struct {
		name string
		data any
		path string
		want string
	}{
		{
			name: "If path is '.', return src, 1",
			data: map[string]any{
				"uno": 1,
				"dos": 2,
			},
			path: ".",
			want: "",
		},
		{
			name: "If path is '.', return src, 3",
			data: map[string]any{
				"uno": 111,
				"dos": 222,
			},
			path: "dos",
			want: "222",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetString(tt.data, tt.path); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

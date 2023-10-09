package yaml

import (
	"reflect"
	"testing"
)

func TestYAML_Get(t *testing.T) {
	tests := []struct {
		name    string
		src     any
		path    string
		want    any
		wantErr bool
	}{
		{
			name:    "If path is empty, return src, 1",
			src:     "Hola",
			path:    "",
			want:    "Hola",
			wantErr: false,
		},
		{
			name:    "If path is '.', return src, 2",
			src:     "Hola",
			path:    ".",
			want:    "Hola",
			wantErr: false,
		},
		{
			name: "If path is '.', return src, 3",
			src: map[string]any{
				"uno": 1,
				"dos": 2,
			},
			path: ".",
			want: map[string]any{
				"uno": 1,
				"dos": 2,
			},
			wantErr: false,
		},
		{
			name:    "Error if src is not a map or array",
			src:     "Hola",
			path:    "juan",
			want:    nil,
			wantErr: true,
		},
		{
			name: "Src is a Map, path does not match, 1",
			src: map[string]any{
				"uno": 1,
				"dos": 2,
			},
			path:    "juan",
			want:    nil,
			wantErr: true,
		},
		{
			name: "Src is a Map, path matches, 2",
			src: map[string]any{
				"uno": 1,
				"dos": 2,
			},
			path:    "dos",
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.src, tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("YAML.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YAML.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

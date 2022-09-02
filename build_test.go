package micro_proto

import (
	"testing"
)

func TestBuildProto(t *testing.T) {
	tests := []struct {
		name string
		want error
	}{
		{
			name: "",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildProto(); got != tt.want {
				t.Errorf("BuildProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

package errors

import (
	"testing"
)

func TestEcode_Int32(t *testing.T) {
	tests := []struct {
		name string
		e    Ecode
		want int32
	}{
		{
			name: "ok_1",
			e:    Ecode(0),
			want: 0,
		}, {
			name: "ok_2",
			e:    Ecode(-1),
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Int32(); got != tt.want {
				t.Errorf("Ecode.Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEcode_Int64(t *testing.T) {
	tests := []struct {
		name string
		e    Ecode
		want int64
	}{
		{
			name: "ok_1",
			e:    Ecode(0),
			want: 0,
		}, {
			name: "ok_2",
			e:    Ecode(-1),
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Int64(); got != tt.want {
				t.Errorf("Ecode.Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

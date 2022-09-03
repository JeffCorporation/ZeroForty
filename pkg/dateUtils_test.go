package pkg

import (
	"reflect"
	"testing"
	"time"
)

func TestStartOfWeek(t *testing.T) {

	tests := []struct {
		name string
		args time.Time
		want time.Time
	}{
		{
			name: "2022-06-06",
			args: time.Date(2022, 6, 6, 6, 6, 6, 6, time.Local),
			want: time.Date(2022, 6, 5, 0, 0, 0, 0, time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartOfWeek(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartOfWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndOfWeek(t *testing.T) {
	tests := []struct {
		name string
		args time.Time
		want time.Time
	}{
		{
			name: "2022-06-06",
			args: time.Date(2022, 6, 6, 6, 6, 6, 6, time.Local),
			want: time.Date(2022, 6, 11, 23, 59, 59, 999999999, time.Local),
		},
		{
			name: "nanosecs",
			args: time.Date(2022, 6, 6, 6, 6, 6, 6, time.Local),
			want: time.Date(2022, 6, 12, 0, 0, 0, 0, time.Local).Add(-1 * time.Nanosecond),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndOfWeek(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EndOfWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}

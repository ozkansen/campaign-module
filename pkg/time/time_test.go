package time

import (
	"reflect"
	"testing"
	"time"
)

func TestIncrease(t *testing.T) {
	type args struct {
		h     int
		tFunc func()
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "increase valid",
			args: args{
				h: 5,
				tFunc: func() {
					Now = func() time.Time {
						return time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
					}
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.tFunc()
			Increase(tt.args.h)
			if !reflect.DeepEqual(Now(), time.Date(2000, 1, 1, 5, 0, 0, 0, time.UTC)) {
				t.Error("increase date not valid")
			}
		})
	}
}

func TestSet(t *testing.T) {
	type args struct {
		t Time
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "valid test",
			args: args{
				t: Time{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Set(tt.args.t)
			if !reflect.DeepEqual(tt.args.t, Now()) {
				t.Error("time set error")
			}
		})
	}
}

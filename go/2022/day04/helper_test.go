package main

import "testing"

// In how many assignment pairs does one range fully contain the other?

func Test_Contains(t *testing.T) {
	type fields struct {
		left  int
		right int
	}
	type args struct {
		pair task
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{{
		name:   "should be true if the task is fully contained in the other",
		fields: fields{left: 4, right: 6},
		args:   args{pair: task{left: 6, right: 6}},
		want:   true,
	}, {
		name:   "should return false",
		fields: fields{left: 6, right: 8},
		args:   args{pair: task{left: 2, right: 4}},
		want:   false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := task{
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := a.Contains(tt.args.pair); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Overlaps(t *testing.T) {
	type fields struct {
		left, right int
	}
	type args struct {
		pair task
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{{
		name:   "should return true if it's not even in range",
		fields: fields{left: 4, right: 8},
		args:   args{pair: task{left: 10, right: 12}},
		want:   false,
	}, {
		name:   "should return true if partially out of range",
		fields: fields{left: 4, right: 8},
		args:   args{pair: task{left: 6, right: 9}},
		want:   true,
	}, {
		name:   "should return true if partially out of range (i.e. lower)",
		fields: fields{left: 4, right: 8},
		args:   args{pair: task{left: 2, right: 5}},
		want:   true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := task{
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := a.Overlaps(tt.args.pair); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

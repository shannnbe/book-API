package main

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name  string
		args  args
		wantZ int
	}{
		{
			name: "keven",
			args: args{x: 1, y: 2},
			wantZ: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotZ := add(tt.args.x, tt.args.y); gotZ != tt.wantZ {
				t.Errorf("add() = %v, want %v", gotZ, tt.wantZ)
			}
		})
	}
}

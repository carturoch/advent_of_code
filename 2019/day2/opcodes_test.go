package main

import (
	"reflect"
	"testing"
)

func TestExec(t *testing.T) {
	type args struct {
		instructions []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Process opcode 1",
			args{[]int{1,0,0,0,99}},
			[]int{2,0,0,0,99},
		},
		{
			"Process opcode 2",
			args{[]int{2,3,0,3,99}},
			[]int{2,3,0,6,99},
		},
		{
			"Process offset instructions",
			args{[]int{2,4,4,5,99,0}},
			[]int{2,4,4,5,99,9801},
		},
		{
			"Process multiple instructions",
			args{[]int{1,1,1,4,99,5,6,0,99}},
			[]int{30,1,1,4,2,5,6,0,99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exec(tt.args.instructions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

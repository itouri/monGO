package handlers

import (
	"reflect"
	"testing"
)

func TestMakeOpSlice(t *testing.T) {
	exprs := []string{
		"1+2",
		"1*2",
		"1+2*3-4/5",
		"1+44+55/666",
		"(1+2)*3",
		"((1+2)*3)",
	}
	ans := [][]int{
		{1, Plus, 2},
		{1, Times, 2},
		{1, Plus, 2, Times, 3, Minus, 4, Divide, 5},
		{1, Plus, 44, Plus, 55, Divide, 666},
		{PaLeft, 1, Plus, 2, PaRight, Times, 3},
		{PaLeft, PaLeft, 1, Plus, 2, PaRight, Times, 3, PaRight},
	}
	for i, expr := range exprs {
		check := makeOpSlice(expr)
		if !reflect.DeepEqual(ans[i], check) {
			t.Errorf("Miss matched %s: %#v <=> %#v", expr, ans[i], check)
		}
	}
}

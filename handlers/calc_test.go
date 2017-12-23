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

func TestConvertToRPN(t *testing.T) {
	exprs := [][]int{
		{1, Plus, 2},
		{1, Times, 2},
		{1, Plus, 2, Times, 3, Minus, 4, Divide, 5},
		{1, Plus, 44, Plus, 55, Divide, 666},
		{PaLeft, 1, Plus, 2, PaRight, Times, 3},
		//{PaLeft, PaLeft, 1, Plus, 2, PaRight, Times, 3, PaRight},
	}
	rpn := [][]int{
		{1, 2, Plus},
		{1, 2, Times},
		{1, 2, 3, Times, 4, 5, Divide, Minus, Plus},
		{1, 44, 55, 666, Divide, Plus, Plus},
		{1, 2, Plus, 3, Times},
		//{PaLeft, PaLeft, 1, Plus, 2, PaRight, Times, 3, PaRight},
	}
	for i, expr := range exprs {
		check := convertToRPN(expr)
		if !reflect.DeepEqual(rpn[i], check) {
			t.Errorf("Miss matched %#v <=> %#v", rpn[i], check)
		}
	}
}

func TestCalcRPN(t *testing.T) {
	rpns := [][]int{
		{1, 2, Plus},
		{1, 2, Times},
		{1, 2, 3, Times, 4, 2, Divide, Minus, Plus},
		{1, 44, 56, 2, Divide, Plus, Plus},
		{1, 2, Plus, 3, Times},
		//{PaLeft, PaLeft, 1, Plus, 2, PaRight, Times, 3, PaRight},
	}
	ans := []int{
		3,
		2,
		5,
		73,
		9,
	}
	for i, rpn := range rpns {
		check := clacRPN(rpn)
		if !reflect.DeepEqual(ans[i], check) {
			t.Errorf("Miss matched %#v <=> %#v", ans[i], check)
		}
	}
}

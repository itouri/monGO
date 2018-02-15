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
		"-7+3",
		"3+(-7)",
		"(1+2)*3",
		"((1+2)*3)",
	}
	ans := [][]Ops{
		{1, '+', 2},
		{1, '*', 2},
		{1, '+', 2, '*', 3, '-', 4, '/', 5},
		{1, '+', 44, '+', 55, '/', 666},
		{-7, '+', 3},
		{3, '+', -7},
		{'(', 1, '+', 2, ')', '*', 3},
		{'(', '(', 1, '+', 2, ')', '*', 3, ')'},
	}
	for i, expr := range exprs {
		check := makeOpSlice(expr)
		if !reflect.DeepEqual(ans[i], check) {
			t.Errorf("Miss matched %s: %#v <=> %#v", expr, ans[i], check)
		}
	}
}

func TestConvertToRPN(t *testing.T) {
	exprs := [][]Ops{
		{1, '+', 2},
		{1, '*', 2},
		{1, '+', 2, '*', 3, '-', 4, '/', 5},
		{1, '+', 44, '+', 55, '/', 666},
		{-7, '+', 3},
		{3, '+', -7},
		{'(', 1, '+', 2, ')', '*', 3},
		//{'(', '(', 1, '+', 2, ')', '*', 3, ')'},
	}
	rpn := [][]Ops{
		{1, 2, '+'},
		{1, 2, '*'},
		{1, 2, 3, '*', 4, 5, '/', '-', '+'},
		{1, 44, 55, 666, '/', '+', '+'},
		{-7, 3, '+'},
		{3, -7, '+'},
		{1, 2, '+', 3, '*'},
		//{'(', '(', 1, '+', 2, ')', '*', 3, ')'},
	}
	for i, expr := range exprs {
		check := convertToRPN(expr)
		if !reflect.DeepEqual(rpn[i], check) {
			t.Errorf("Miss matched %#v <=> %#v", rpn[i], check)
		}
	}
}

func TestCalcRPN(t *testing.T) {
	rpns := [][]Ops{
		{1, 2, '+'},
		{1, 2, '*'},
		{1, 2, 3, '*', 4, 2, '/', '-', '+'},
		{1, 44, 56, 2, '/', '+', '+'},
		{1, 2, '+', 3, '*'},
		{-7, 3, '+'},
		{3, -7, '+'},
		//{'(', '(', 1, '+', 2, ')', '*', 3, ')'},
	}
	ans := []int{
		3,
		2,
		5,
		73,
		9,
		-4,
		-4,
	}
	for i, rpn := range rpns {
		check := clacRPN(rpn)
		if !reflect.DeepEqual(ans[i], check) {
			t.Errorf("Miss matched %#v <=> %#v", ans[i], check)
		}
	}
}

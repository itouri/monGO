package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

const (
	Times   = -1
	Divide  = -2
	Plus    = -3
	Minus   = -4
	PaLeft  = -5
	PaRight = -6
)

// change priority
func cP(in int) int {
	switch in {
	case -1:
	case -2:
		return 1
	case -3:
	case -4:
		return 2
	case -5:
	case -6:
		return 3
	default:
		return 0
	}
}

func a(sli []int, v int, str string) []int {
	//TODO error handling
	if str != "" {
		val, _ := strconv.Atoi(str)
		sli = append(sli, val)
	}
	sli = append(sli, v)
	return sli
}

func makeOpSlice(str string) []int {
	ops := []int{}
	var numStr string
	for _, r := range str {
		isNum := false
		switch {
		case '0' <= r && r <= '9':
			numStr += string(r)
			isNum = true
		case r == '+':
			ops = a(ops, Plus, numStr)
		case r == '-':
			ops = a(ops, Minus, numStr)
		case r == '*':
			ops = a(ops, Times, numStr)
		case r == '/':
			ops = a(ops, Divide, numStr)
		case r == '(':
			ops = a(ops, PaLeft, numStr)
		case r == ')':
			ops = a(ops, PaRight, numStr)
		default:
			// return ERROR
		}
		if !isNum {
			numStr = ""
		}
	}
	if numStr != "" {
		val, _ := strconv.Atoi(numStr)
		ops = append(ops, val)
	}
	return ops
}

func clac(str string) {
	// ops := makeOpSlice(str)
	// rpn := convertToRPN(ops)
}

func GetCalc(c echo.Context) error {
	str := c.QueryString()

	return c.String(http.StatusOK, str)
}

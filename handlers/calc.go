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
	case -1, -2:
		return 1
	case -3, -4:
		return 2
	case -5, -6:
		return 3
	default:
		return 0
	}
	return -1
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

func convertToRPN(ops []int) []int {
	p := []int{}
	s := []int{}
	for _, op := range ops {
		if op == PaLeft {
			s = append(s, op)
			continue
		}
		if op == PaRight {
			for s[len(s)-1] != PaLeft {
				p = append(p, s[len(s)-1])
				s = s[:len(s)-1]
			}
			// discard '('
			s = s[:len(s)-1]
			// skip adding ')'
			continue
		}
		if len(s) != 0 {
			//fmt.Printf("%d vs %d\n", cP(op), cP(s[len(s)-1]))
			for cP(op) > cP(s[len(s)-1]) {
				p = append(p, s[len(s)-1])
				s = s[:len(s)-1]
				if len(s) == 0 {
					break
				}
			}
		}
		s = append(s, op)
	}
	for i := len(s) - 1; i >= 0; i-- {
		p = append(p, s[i])
	}
	return p
}

func clacRPN(ops []int) int {
	for i := 0; i < len(ops); i++ {
		// +-/*
		if ops[i] < 0 {
			// val = val1 - val2
			var val int
			val1 := ops[i-2]
			val2 := ops[i-1]
			switch ops[i] {
			case Times:
				val = val1 * val2
			case Divide:
				val = val1 / val2
			case Plus:
				val = val1 + val2
			case Minus:
				val = val1 - val2
			}
			// [5, 1, 2, +(i), -] -> [5, 3, -]
			ops[i] = val
			ops = append(ops[:i-2], ops[i:]...)
			i -= 2
		}
		if len(ops) == 1 {
			break
		}
	}
	return ops[0]
}

func clac(str string) int {
	ops := makeOpSlice(str)
	rpn := convertToRPN(ops)
	ans := clacRPN(rpn)
	return ans
}

func GetCalc(c echo.Context) error {
	str := c.QueryString()

	return c.String(http.StatusOK, str)
}

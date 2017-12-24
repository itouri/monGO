package handlers

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

const (
	Times   = -991
	Divide  = -992
	Plus    = -993
	Minus   = -994
	PaLeft  = -995
	PaRight = -996
)

// change priority
func cP(in int) int {
	switch in {
	case -991, -992:
		return 1
	case -993, -994:
		return 2
	case -995, -996:
		return 3
	default:
		return 0
	}
	return -1
}

func a(sli []int, v int, str string, isMinus bool) []int {
	//TODO error handling
	if str != "" {
		val, _ := strconv.Atoi(str)
		if isMinus {
			val *= -1
		}
		sli = append(sli, val)
	}
	if v == PaRight && isMinus {
		return sli
	}
	sli = append(sli, v)
	return sli
}

func makeOpSlice(str string) []int {
	ops := []int{}
	var numStr string
	isSkip := false
	isMinus := false
	for i, r := range str {
		isNum := false
		if isSkip {
			isSkip = false
			continue
		}
		switch {
		case '0' <= r && r <= '9':
			numStr += string(r)
			isNum = true
		case r == '+':
			ops = a(ops, Plus, numStr, isMinus)
		case r == '-':
			if i == 0 {
				isMinus = true
				continue
			}
			ops = a(ops, Minus, numStr, isMinus)
		case r == '*':
			ops = a(ops, Times, numStr, isMinus)
		case r == '/':
			ops = a(ops, Divide, numStr, isMinus)
		case r == '(':
			if str[i+1] == '-' {
				isMinus = true
				isSkip = true
				continue
			}
			ops = a(ops, PaLeft, numStr, isMinus)
		case r == ')':
			ops = a(ops, PaRight, numStr, isMinus)
		default:
			// return ERROR
		}
		if !isNum {
			numStr = ""
			isMinus = false
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
		if -996 <= ops[i] && ops[i] <= -991 {
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
	if strings.Count(str, "(") != strings.Count(str, ")") {
		return c.String(http.StatusOK, "ERROR")
		//return fmt.Errorf("Don't match '(' and ')' numbers")
	}
	r := regexp.MustCompile(`[^/+\-\*\/()0-9]`)
	if r.MatchString(str) {
		return c.String(http.StatusOK, "ERROR")
		//return fmt.Errorf("string include not valid char")
	}

	return c.String(http.StatusOK, strconv.Itoa(clac(str)))
}

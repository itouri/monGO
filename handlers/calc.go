package handlers

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

type Ops interface{}

// change priority
func cP(in Ops) int {
	switch in.(type) {
	case rune:
		switch in {
		case '*', '/':
			return 1
		case '+', '-':
			return 2
		case '(', ')':
			return 3
		default:
			return -1
		}
	case int:
		return 0
	default:
		return -1
	}
	return -1
}

func a(sli []Ops, v rune, str string, isMinus bool) []Ops {
	//TODO error handling
	if str != "" {
		val, _ := strconv.Atoi(str)
		if isMinus {
			val *= -1
		}
		sli = append(sli, val)
	}
	if v == ')' && isMinus {
		return sli
	}
	sli = append(sli, v)
	return sli
}

func makeOpSlice(str string) []Ops {
	var ops []Ops
	var numStr string
	isSkip := false
	isMinus := false
	for i, r := range str {
		isNum := false
		if isSkip {
			isSkip = false
			continue
		}
		switch r {
		case '+', '*', '/', ')':
			ops = a(ops, r, numStr, isMinus)
		case '-':
			if i == 0 {
				isMinus = true
				continue
			}
			ops = a(ops, r, numStr, isMinus)
		case '(':
			if str[i+1] == '-' {
				isMinus = true
				isSkip = true
				continue
			}
			ops = a(ops, r, numStr, isMinus)
		default: // r is int
			numStr += string(r)
			isNum = true
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

func convertToRPN(ops []Ops) []Ops {
	p := []Ops{}
	s := []Ops{}
	for _, op := range ops {
		if op == '(' {
			s = append(s, op)
			continue
		}
		if op == ')' {
			for s[len(s)-1] != '(' {
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

func clacRPN(ops []Ops) int {
	for i := 0; i < len(ops); i++ {
		switch ops[i].(type) {
		case rune:
			// val = val1 - val2
			var val int
			val1, ok := ops[i-2].(int)
			if !ok {
				return -1
			}
			val2, ok := ops[i-1].(int)
			if !ok {
				return -1
			}
			switch ops[i] {
			case '*':
				val = val1 * val2
			case '/':
				val = val1 / val2
			case '+':
				val = val1 + val2
			case '-':
				val = val1 - val2
			}
			// [5, 1, 2, +(i), -] -> [5, 3, -]
			ops[i] = val
			ops = append(ops[:i-2], ops[i:]...)
			i -= 2
		case int:
		default:
			return -1
		}

		if len(ops) == 1 {
			break
		}
	}
	ret, ok := ops[0].(int)
	if !ok {
		return -1
	}
	return ret
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

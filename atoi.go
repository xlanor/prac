package atoi

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(s string) int {
	t := strings.TrimSpace(s)
	res := 0
	if t == "" {
		return res
	}
	start := 0
	sign := 1
	overflow := false

	sr := []rune(s)
	for id, ch := range sr {
		if id == start && string(ch) == " " {
			start += 1
			continue
		}
		if id == start && string(ch) == "-" {
			sign = -1
			continue
		}
		if id == start && string(ch) == "+" {
			sign = 1
			continue
		}
		castInt := int(byte(ch) - '0')
		if castInt < 10 && castInt >= 0 {
			if castInt > math.MaxInt32-(res*10) {
				// integer overflow.
				overflow = true
				break
			} else {

				res = res*10 + castInt
			}
		} else {
			break
		}
	}
	if overflow {
		if sign == -1 {

			res = math.MinInt32
		} else {
			res = math.MaxInt32
		}
	} else {
		res = res * sign
	}
	if res <= (-1 << 31) {
		res = -1 << 31
	}
	if res >= (1<<31 - 1) {
		res = 1<<31 - 1
	}

	return res

}

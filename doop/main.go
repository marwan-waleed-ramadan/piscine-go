package main

import (
	"os"
)

const (
	maxInt64 = 9223372036854775807
	minInt64 = -9223372036854775808
)

func atoi(s string) (int64, bool) {
	var n int64
	sign := int64(1)
	i := 0

	if s == "" {
		return 0, false
	}

	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		digit := int64(s[i] - '0')

		// Check for overflow
		if sign == 1 && (n > (maxInt64-digit)/10) {
			return 0, false
		}
		if sign == -1 && (-n < (minInt64+digit)/10) {
			return 0, false
		}
		n = n*10 + digit
	}
	return n * sign, true
}

func putnbr(n int64) {
	if n == 0 {
		os.Stdout.Write([]byte{'0', '\n'})
		return
	}
	if n < 0 {
		os.Stdout.Write([]byte{'-'})
		n = -n
	}
	var digits [20]byte
	i := len(digits)
	for n > 0 {
		i--
		digits[i] = byte(n%10) + '0'
		n /= 10
	}
	os.Stdout.Write(digits[i:])
	os.Stdout.Write([]byte{'\n'})
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok1 := atoi(os.Args[1])
	b, ok2 := atoi(os.Args[3])
	if !ok1 || !ok2 {
		return
	}

	op := os.Args[2]
	switch op {
	case "+":
		if (b > 0 && a > maxInt64-b) || (b < 0 && a < minInt64-b) {
			return
		}
		putnbr(a + b)
	case "-":
		if (b < 0 && a > maxInt64+b) || (b > 0 && a < minInt64+b) {
			return
		}
		putnbr(a - b)
	case "*":
		if (a > 0 && b > 0 && a > maxInt64/b) ||
			(a < 0 && b < 0 && a < maxInt64/b) ||
			(a > 0 && b < 0 && b < minInt64/a) ||
			(a < 0 && b > 0 && a < minInt64/b) {
			return
		}
		putnbr(a * b)
	case "/":
		if b == 0 {
			os.Stdout.Write([]byte("No division by 0\n"))
			return
		}
		putnbr(a / b)
	case "%":
		if b == 0 {
			os.Stdout.Write([]byte("No modulo by 0\n"))
			return
		}
		putnbr(a % b)
	default:
		return
	}
}

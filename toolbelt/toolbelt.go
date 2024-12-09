package toolbelt

import (
	"fmt"
	"strconv"
	"unicode"
)

func Must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ":" + err.Error())
	}
}

func RuneToIntIfDigit(r rune) (int, bool) {
	if unicode.IsDigit(r) {
		return 0, false
	}
	return int(r - '0'), true
}

func ByteToIntIfDigit(b byte) (int, bool) {
	if b >= '0' && b <= '9' {
		return int(b - '0'), true
	}
	return 0, false
}

func EzIntParse(digits string) int {
	num, err := strconv.Atoi(digits)
	Must("parse number", err)
	return num
}

func StrArrToInts(ss []string) []int {
	ns := make([]int, 0, len(ss))
	for _, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err.Error())
		}
		ns = append(ns, n)
	}
	return ns
}

func RemoveFirst[T comparable](arr []T, toRemove T) []T {
	for i, el := range arr {
		if el == toRemove {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

func RemoveAllStrArr(arr []string, toRemove string) (newArr []string) {
	for _, el := range arr {
		if el != toRemove {
			newArr = append(newArr, el)
		}
	}
	return newArr
}

// Deprecated. Use generic All
func AllInStrArr(arr []string, val string) bool {
	for _, arrElement := range arr {
		if arrElement != val {
			return false
		}
	}
	return true
}

func All[E comparable](s []E, target E) bool {
	for _, e := range s {
		if e != target {
			return false
		}
	}
	return true
}

func StrArrToIntMap(arr []string) map[int]int {
	result := make(map[int]int)
	for _, digit := range arr {
		digitInt, err := strconv.Atoi(digit)
		if err != nil {
			fmt.Println("Error converting digit to int")
		}
		result[digitInt] += 1
	}

	return result
}

func MapValSum[T comparable](_map map[T]int) (sum int) {
	for _, val := range _map {
		sum += val
	}

	return sum
}

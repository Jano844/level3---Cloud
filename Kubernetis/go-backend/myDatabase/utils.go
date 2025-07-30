package myDatabase

import (
	"fmt"
	"strconv"
)

func ToAlphaBase(n int) string {
	if n < 0 {
		return "-" + ToAlphaBase(-n)
	}
	result := ""
	for {
		remainder := n % 26
		result = string('a'+remainder) + result
		n = n/26 - 1
		if n < 0 {
			break
		}
	}
	return result
}

func ToAlphaBaseStr(input string) (string, error) {
	num, err := strconv.Atoi(input)
	if err != nil {
		return "", fmt.Errorf("invalid number input: %v", err)
	}
	return ToAlphaBase(num), nil
}
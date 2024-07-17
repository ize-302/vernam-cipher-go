package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func GenerateRandomBinary () string {
	result := ""; 
	for index:=0; index < 8; index++ {
		var randomNum = rand.Intn(2)
		result += fmt.Sprintf("%d", randomNum);
	} 
	return result
}

// Finds the bitwise value (xor) between two binary numbers
func FindXOR(bin1, bin2 string) (string, error) {
	if len(bin1) != len(bin2) {
		return "", fmt.Errorf("binary strings must be of equal length")
	}
	result := ""
	for i := 0; i < len(bin1); i++ {
		bit1, err1 := strconv.Atoi(string(bin1[i]))
		bit2, err2 := strconv.Atoi(string(bin2[i]))
		if err1 != nil || err2 != nil {
			return "", fmt.Errorf("invalid binary string")
		}
		xorBit := bit1 ^ bit2
		result += strconv.Itoa(xorBit)
	}
	return result, nil
}
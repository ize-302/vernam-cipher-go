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

// ConvertBinaryToHex converts a binary string to a hexadecimal string.
func ConvertBinaryToHex(binaryString string) (string, error) {
	// Convert the binary string to an integer
	number, err := strconv.ParseUint(binaryString, 2, 64)
	if err != nil {
		return "", err
	}
	// Convert the integer to a hexadecimal string
	hexString := fmt.Sprintf("%X", number)
	return hexString, nil
}

// ConvertHexToBinary converts a hexadecimal string to a binary string.
func ConvertHexToBinary(hexString string) (string, error) {
	// Convert the hexadecimal string to an integer
	number, err := strconv.ParseUint(hexString, 16, 64)
	if err != nil {
			return "", err
	}
	// Convert the integer to a binary string
	binaryString := fmt.Sprintf("%08b", number)
	return binaryString, nil
}
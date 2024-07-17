package main

import (
	"fmt"
	"strconv"
	"strings"
)


func encryptMessage (message string) {
	// Generate key for each character
	var charCountInMessage = len(message)
	// stores generated secret keys
	var cipherKeys string
	// stores encrypted message i.e ciphertext
	var cipherTexts string

	/*
	* Produce random key and corresponding cipherText for each character in the message
	*/
	for i:=0; i < charCountInMessage; i++ {
		randomKey := GenerateRandomBinary()

		// Doing this to avoid additing space before the first cipher Key
		if i == 0 {
			cipherKeys += randomKey
			} else {
			cipherKeys += " " + randomKey
		}

		// Gets the corresponding ascii value of a current character
		asciiValue := byte(message[i])
		// Convert to binary character
		binaryRepresentation := fmt.Sprintf("%08b", asciiValue)
		// find the bitwise (xor) of the secret key and binary representation of current character
		xorResult, err := FindXOR(randomKey, binaryRepresentation)

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			if i == 0 {
				cipherTexts += xorResult
				} else {
				cipherTexts += " " + xorResult
			}
		}
	}
	fmt.Println("Message encrypted:")
	fmt.Printf("Cipher Key => : %s\n",cipherKeys)
	fmt.Printf("Encrypted Message =>: %s\n",cipherTexts)
}

func decryptMessage(cipherText, cipherKey string) {
	// stores the decrypted message
	var message string

	// Split the data string into individual binary strings
	cipherTextAsArray := strings.Split(cipherText, " ")
	cipherKeyAsArray := strings.Split(cipherKey, " ")

	for i:=0; i<len(cipherTextAsArray); i++ {
		xorResult, err := FindXOR(cipherKeyAsArray[i], cipherTextAsArray[i])

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			// convert the bitwise(xor) result to whole number
			number, err := strconv.ParseUint(xorResult, 2, 8)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				// get the corresponding ascii character
				char := rune(number)
				message += string(char)
			}
		}
	}
	fmt.Println("Message Decrypted:")
	fmt.Println("message =>",message)
}

func main() {
	// Encrypting a message
	encryptMessage("HELLO, WORLD!")

	fmt.Println("-------------------------------------")

	// Decrypting a message using the secret key and encrypted message (both in binary format)
	cipherText := "10000000 01101111 11101110 01101010 10101011 10101001 01100111 00011100 00111001 00010000 10111111 11111110 00010100"
	cipherKey := "11001000 00101010 10100010 00100110 11100100 10000101 01000111 01001011 01110110 01000010 11110011 10111010 00110101"
	decryptMessage(cipherText, cipherKey)
}

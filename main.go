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
		hexValue, err := ConvertBinaryToHex(randomKey)

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			// Doing this to avoid additing space before the first cipher Key
			if i == 0 {
				cipherKeys += hexValue
				} else {
					cipherKeys += " " + hexValue
			}
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
			// convert binary to hex
			hexValue, err := ConvertBinaryToHex(xorResult)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				if i == 0 {
					cipherTexts += hexValue
					} else {
					cipherTexts += " " + hexValue
				}
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

	// Split the data string into individual hex strings
	cipherTextAsArray := strings.Split(cipherText, " ")
	cipherKeyAsArray := strings.Split(cipherKey, " ")

	for i:=0; i<len(cipherTextAsArray); i++ {
		// convert the current hex values for text and key into binary
		TextInBinaryFormat, textErr := ConvertHexToBinary(cipherTextAsArray[i])
		KeyInBinaryFormat, keyErr := ConvertHexToBinary(cipherKeyAsArray[i])

		if textErr != nil && keyErr != nil {
			fmt.Println("Error:", textErr)	
		} else {
			xorResult, err := FindXOR(KeyInBinaryFormat, TextInBinaryFormat)
	
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

	}
	fmt.Println("Message Decrypted:")
	fmt.Println("message =>",message)
}

func main() {
	// Encrypting a message
	encryptMessage("HELLO, WORLD!")

	fmt.Println("-------------------------------------")

	// Decrypting a message using the secret key and encrypted message (both in hex format)
	cipherText := "EE E3 9 A3 56 7 D6 5C 52 60 12 93 11"
	cipherKey := "A6 A6 45 EF 19 2B F6 B 1D 32 5E D7 30"
	decryptMessage(cipherText, cipherKey)
}

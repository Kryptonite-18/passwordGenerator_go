package main

import (
	"fmt"
	"math/rand"
	"time"
)

const lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
const uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const specialChars = "!@#$%^&*-{};'.<>/?`"

func main() {
	var useLetters, useNumbers, useSpecialChars, useUppercaseLetters bool

	var useLetters_, useNumbers_, useSpecialChars_, useUppercaseLetters_ string

	fmt.Println("Enter the size of your password")
	var length int
	fmt.Scanln(&length)

	fmt.Println("do you want lowerCase in your password : y/n")
	fmt.Scanln(&useLetters_)
	if useLetters_ == "y" || useLetters_ == "Y" {
		useLetters = true
	} else {
		useLetters = false
	}

	fmt.Println("do you want numbers in your password : y/n")
	fmt.Scanln(&useNumbers_)
	if useNumbers_ == "y" || useNumbers_ == "Y" {
		useNumbers = true
	} else {
		useNumbers = false
	}
	fmt.Println("do you want special characters in your password : y/n")

	fmt.Scanln(&useSpecialChars_)
	if useSpecialChars_ == "y" || useSpecialChars_ == "Y" {
		useSpecialChars = true
	}

	fmt.Println("do you want upperCase letters in your password: y/n")

	fmt.Scanln(&useUppercaseLetters_)
	if useUppercaseLetters_ == "y" || useSpecialChars_ == "Y" {
		useUppercaseLetters = true
	}

	password := passwordGenerate(length, useLetters, useUppercaseLetters, useNumbers, useSpecialChars)
	fmt.Println("Generated password:", password)

}

func passwordGenerate(length int, useLetter bool, uppercaseletters bool, useNumbers bool, useSpecial bool) string {
	var characterSet string

	// var includedLetters, includedNumbers, includedSpecials bool
	if useLetter {
		characterSet += lowercaseLetters
		if uppercaseletters {
			characterSet += uppercaseLetters
		}
	}
	if useNumbers {
		characterSet += numbers
	}

	if useSpecial {
		characterSet += specialChars
	}

	if len(characterSet) == 0 {
		panic("Please select at least one character set") // Ensure at least one set is chosen
	}

	bytes := make([]byte, length)

	bt := tryingPasswords(bytes, characterSet)

	if length <= 2 {
		tryingPasswords(bytes, characterSet)
	}

	if useLetter == true && useNumbers == true && useSpecial == true && uppercaseletters == true {
		for checkLetters(bt) == false || checkNumbers(bt) == false || checkSpecials(bt) == false || uppercaseletters == false {
			tryingPasswords(bytes, characterSet)
		}
	}
	if useLetter == true && useNumbers == true && uppercaseletters == true {
		for checkLetters(bt) == false || checkNumbers(bt) == false || uppercaseletters == false {
			tryingPasswords(bytes, characterSet)
		}
	}

	if useLetter == true && useSpecial == true && uppercaseletters == true {
		for checkLetters(bt) == false || checkSpecials(bt) == false || uppercaseletters == false {
			tryingPasswords(bytes, characterSet)
		}
	}
	if useNumbers == true && useSpecial == true && uppercaseletters == true {
		for checkNumbers(bt) == false || checkSpecials(bt) == false || uppercaseletters == false {
			tryingPasswords(bytes, characterSet)
		}
	}

	if useSpecial == true && uppercaseletters == true {
		for checkSpecials(bt) == false || uppercaseletters == false {
			tryingPasswords(bytes, characterSet)
		}
	}

	if useNumbers == true && uppercaseletters == true {
		for checkNumbers(bt) == false || checkUpperCase(bt) {
			tryingPasswords(bytes, characterSet)
		}
	}

	return string(bytes)
}

func tryingPasswords(bytes []byte, characterSet string) []byte {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range bytes {
		randomIndex := byte(r.Intn(len(characterSet) - 1))
		bytes[i] = characterSet[randomIndex]

	}
	return bytes
}

func checkNumbers(char []byte) bool {
	for i := range numbers {
		for j := 0; j < len(char); j++ {
			if numbers[i] == char[j] {
				return true
			}
		}
	}

	return false
}

func checkLetters(char []byte) bool {
	for i := range lowercaseLetters {
		for j := 0; j < len(char); j++ {
			if lowercaseLetters[i] == char[j] {
				return true
			}
		}
	}

	return false
}

func checkSpecials(char []byte) bool {
	for i := range specialChars {
		for j := 0; j < len(char); j++ {
			if specialChars[i] == char[j] {
				return true
			}
		}
	}

	return false
}

func checkUpperCase(char []byte) bool {
	for i := range uppercaseLetters {
		for j := 0; j < len(char); j++ {
			if uppercaseLetters[i] == char[j] {
				return true
			}
		}
	}

	return false
}

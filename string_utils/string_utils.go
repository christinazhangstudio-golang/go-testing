package string_utils

import (
	"strings"
	"fmt"
)

//lower case, private function, but can be used in strings_utils_test.go
func isPalindrome(text string) bool {
	if len(text) <= 1 {
		return true
	}
	
	if len(text) == 2 && text[0] == text[1]{
		return true
	}

	if text[0] != text[len(text) - 1] {
		return false
	}

	return isPalindrome(text[1 : len(text) - 1])
}

//upper case, public function
func IsPalindrome(text string) bool {
	fmt.Println("calling string utils library")
	
	//passing the space-removed text improves runtime, benchmarking helps confirm this
	return isPalindrome(strings.Replace(text, " ", "", -1))
}
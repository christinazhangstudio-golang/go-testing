package main

import (
	"fmt"
	"github.tesla.com/chrzhang/go-testing/string_utils"
	"github.tesla.com/chrzhang/go-testing/services"
)

func main(){
	//since capital letter, is public function
	fmt.Println(string_utils.IsPalindrome("anna"))

	//using the interface
	fmt.Println(services.StringsService.IsPalindrome("ana"))
}
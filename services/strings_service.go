package services

import (
	"github.tesla.com/chrzhang/go-testing/string_utils"
)

//public variable
var (
	//instance of stringsService
	//StringsService = stringsService{}
	StringsService stringServiceInterface
)

func init(){
	//new instance of stringsService
	StringsService = &stringsService{}
}

type stringsService struct{
	Name string
}

type stringServiceInterface interface {
	//for any struct to be considered stringsService, must implement method
	IsPalindrome(someText string) bool
	//now stringsService struct is implementing stringServiceInterface
	//so now we can make StringsService a type of stringServiceInterface
	//any struct that implements an interface can be a service

	//you can add more methods if needed here, 
	//e.g. Method1(text string) bool
	//Method2(text string) bool
	//provided you define them as with IsPalindrome
}

//func IsPalindrome(text string) bool {
//	return string_utils.IsPalindrome(text)
//}
//change this to a method, which defines behavior of stringsService struct
func (s *stringsService) IsPalindrome(text string) bool {
	return string_utils.IsPalindrome(text)
}


//the method like above defines a service, can call on db or API, but mock to isolate artifact 
package services

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

var (
	isPalindromeFunction func(t string) bool
)

type serviceMock struct{}

func (s *serviceMock) IsPalindrome(t string) bool {
	return isPalindromeFunction(t)
}

func TestIsPalindrome(t *testing.T){
	//"calling string utils library" prints to show that we're doing an integration test (it accesses another artifact i.e. another library)
	//so this is not a unit test, but integration test!
	//unit tests shouldn't be accessing any database,
	//assert.True(t, IsPalindrome("ana"))
	
	//so in the situation this is supposed to test something accessing the db
	//then we use structs...
	//s := stringsService{}
	//assert.True(t, s.IsPalindrome("ana"))

	//let's say we need different instance of same service on every call
	//in other languages, services typically should be stateless
	//calling a method on a public variable
	//assert.True(t, StringsService.IsPalindrome("ana"))

	//problem is, still calling strings utils library
	//so create an interface

	//create an instance of the mock
	StringsService = &serviceMock{}
	isPalindromeFunction = func(t string) bool {
		return true
	}
	assert.True(t, StringsService.IsPalindrome("ana"))
}
package string_utils

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)

//need to call test cases, m.Run means run all the test cases in the package string_utils
func TestMain(m *testing.M){
	//Do whatever you need to do before running test cases

	//e.g. mock response instead of actually calling API
	//httpmock.Activate()
	
	os.Exit(m.Run())
}

//unit testing, since 4 test cases (including recursive case)...
//you may be tempted to put all the asserts in one single test case, but don't do that! mixing together area returns
//just make single test case for each return
func TestIsPalindromeLenOne(t *testing.T) {
	//if IsPalindrome("a") != true {
	//	t.Error("single character should be a palindrome.")
	//} 

	//can do this instead, which is cleaner to read
	//asserts will continue even if test fails
	assert.True(t, IsPalindrome("a"), "single character should be a palindrome.")
}
func TestIsPalindromeTwoSameCharacters(t *testing.T) {
	assert.True(t, IsPalindrome("aa"), "two same characters should be a palindrome.")
}
func TestIsPalindromeNotPalindrome(t *testing.T) {
	assert.False(t, IsPalindrome("abc"), "should not be a palindrome.")
}
func TestIsPalindromeNoError(t *testing.T) {
	assert.True(t, IsPalindrome("anitalavalatina"), "should be a palindrome.")
}
func TestIsPalindromeWithSpacesNoReplaces(t *testing.T) {
	assert.False(t, isPalindrome("anita lava la tina"), "should not be a palindrome.")
}
func TestIsPalindromeWithSpaces(t *testing.T) {
	assert.True(t, IsPalindrome("anita lava la tina"), "should be a palindrome.")
}

//benchmark test, see performance and whether best approach is taken
//not run with go test
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("anitalavalatina")
	}
}
package testing

import (
	"github.com/stretchr/testify/assert"
	"golang-unitTest/dummy"
	"testing"
)

// To run the test = go run
// To run the test and see each function = go run -v
// To run testing on a specific function = go run -v -run functionName
func TestGreetings(t *testing.T) {
	result := dummy.Greetings("Kevin")

	if result != "Hello Kevin" {
		t.Error("Test Error")
	}
}

//golang assersion testify
//go get github.com/stretchr/testify
func TestGreetingsAssertion(t *testing.T) {
	result := dummy.Greetings("Kevin")
	assert.Equal(t, "Hello Kevin", result, "Test Result Not Match")
}

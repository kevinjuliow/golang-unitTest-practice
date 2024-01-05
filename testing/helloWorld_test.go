package testing

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang-unitTest/dummy"
	"testing"
)

// Benchmark Project
// to run the benchmark = go test -v -bench=.
func BenchmarkGreetings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.Greetings("Testing")
	}
}

// this is the Main func of testing . the name must be TestMain
func TestMain(m *testing.M) {
	//before
	fmt.Println("before testing")

	m.Run() //this will run the test

	//after
	fmt.Println("after testing")
}

// To run the test = go run
// To run the test and see each function = go run -v
// To run testing on a specific function = go run -v -run functionName
func TestGreetings(t *testing.T) {
	result := dummy.Greetings("Kevin")
	if result != "Hello Kevin" {
		t.Error("Test Error")
	}
}

// golang assersion testify
// go get github.com/stretchr/testify
func TestGreetingsAssertion(t *testing.T) {
	result := dummy.Greetings("Kevin")
	assert.Equal(t, "Hello Kevin", result, "Test Result Not Match")
	//The print code will be executed because it's the sama as t.Fail
	//fmt.Println("Printed")
	//if use require instead of assert , it's the same as t.Failnow
}

package dfa

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestDFA_Check(t *testing.T) {
	sensitive := []string{"王八蛋", "王八羔子"}

	dfaInstance := GetDFA()
	dfaInstance.AddBadWords(sensitive)
	//fmt.Println(dfaInstance)
	str := "你个王#八……羔子， 你就是个王*八/蛋"
	words, hitWords, check := dfaInstance.Check(str)
	fmt.Println(words, hitWords, check)

	assert.Equal(t, true, check)
}
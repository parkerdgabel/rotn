package main

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	r := NewRotn(1)
	testStr := "ABC"
	expected := "BCD"
	actual := encrypt(testStr, r)
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v\n", expected, actual)
	}
}

func TestEncryptLowerCase(t *testing.T) {
	r := NewRotn(1)
	testStr := "abc"
	expected := "bcd"
	actual := encrypt(testStr, r)
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v\n", expected, actual)
	}
}

func TestEncryptWhitespace(t *testing.T) {
	r := NewRotn(1)
	testStr := "a b c\n\ta"
	expected := "b c d\n\tb"
	actual := encrypt(testStr, r)
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v\n", expected, actual)
	}
}

func TestEncryptPunct(t *testing.T) {
	r := NewRotn(1)
	testStr := "'abc'!.?"
	expected := "'bcd'!.?"
	actual := encrypt(testStr, r)
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v\n", expected, actual)
	}
}

func TestEncryptDigits(t *testing.T) {
	r := NewRotn(1)
	testStr := "abc123"
	expected := "bcd123"
	actual := encrypt(testStr, r)
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v\n", expected, actual)
	}
}

func TestEncryptProjectSpecExample(t *testing.T) {
	r := NewRotn(-13)
	testStr := "Doctor: \"The tests show that your DNA is ... backwards.\"\nPatient: \"And?\""
	expected := "Qbpgbe: \"Gur grfgf fubj gung lbhe QAN vf ... onpxjneqf.\"\nCngvrag: \"Naq?\""
	actual := encrypt(testStr, r)
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v\n", expected, actual)
	}
}

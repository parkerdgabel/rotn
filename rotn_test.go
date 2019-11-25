package main

import (
	"testing"
)

func TestValidRotation(t *testing.T) {

	r := NewRotn(1)
	actual := r.rotate('A')
	expected := byte('B')
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	actual = r.rotate('Z')
	expected = byte('A')
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestNegativeN(t *testing.T) {
	r := NewRotn(-1)

	actual := r.rotate('A')
	expected := byte('Z')
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	actual = r.rotate('B')
	expected = byte('A')
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestSymmetry(t *testing.T) {
	r := NewRotn(13)

	actual := r.rotate('A')
	actual = r.rotate(actual)
	expected := byte('A')

	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

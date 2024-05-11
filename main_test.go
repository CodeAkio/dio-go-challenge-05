package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(5, 3)
	if result != 8 {
		t.Errorf("add(5, 3) = %v; want 8", result)
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(5, 3)
	if result != 2 {
		t.Errorf("subtract(5, 3) = %v; want 2", result)
	}
}

func TestMultiply(t *testing.T) {
	result := multiply(5, 3)
	if result != 15 {
		t.Errorf("multiply(5, 3) = %v; want 15", result)
	}
}

func TestDivide(t *testing.T) {
	result, err := divide(5, 3)
	if err != nil || result != 5.0/3.0 {
		t.Errorf("divide(5, 3) = %v; want %v", result, 5.0/3.0)
	}
	_, err = divide(5, 0)
	if err == nil {
		t.Errorf("divide(5, 0) should have caused an error")
	}
}

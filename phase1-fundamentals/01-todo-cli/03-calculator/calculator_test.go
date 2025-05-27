package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(2.0, 3.0)
	expected := 5.0
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(5.0, 3.0)
	expected := 2.0
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestMultiply(t *testing.T) {
	result := multiply(2.0, 3.0)
	expected := 6.0
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestDivide(t *testing.T) {
	// Test normal division
	result, err := divide(6.0, 3.0)
	expected := 2.0
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}

	// Test division by zero
	_, err = divide(6.0, 0.0)
	if err == nil {
		t.Error("Expected error for division by zero, got nil")
	}
}

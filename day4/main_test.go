package main

import "testing"

func TestValid(t *testing.T) {
	if val := 111111; !isValid(val, false) {
		t.Errorf("Expected %d to be valid", val)
	}
}

func TestNotValidDoesNotIncrease(t *testing.T) {
	if val := 223450; isValid(val, false) {
		t.Errorf("Expected %d to not be valid", val)
	}
}

func TestNotValidNoDouble(t *testing.T) {
	if val := 123789; isValid(val, false) {
		t.Errorf("Expected %d to not be valid", val)
	}
}

func TestValidWithNoMultiOccurrence(t *testing.T) {
	if val := 112233; !isValid(val, true) {
		t.Errorf("Expected %d to be valid", val)
	}
}

func TestValidWithNoMultiOccurrence2(t *testing.T) {
	if val := 111122; !isValid(val, true) {
		t.Errorf("Expected %d to be valid", val)
	}
}

func TestNotValidWithNoMultiOccurrence(t *testing.T) {
	if val := 123444; isValid(val, true) {
		t.Errorf("Expected %d not to be valid", val)
	}
}

func TestNotValidWithNoMultiOccurrence2(t *testing.T) {
	if val := 144444; isValid(val, true) {
		t.Errorf("Expected %d not to be valid", val)
	}
}

func TestNotValidWithNoMultiOccurrence3(t *testing.T) {
	if val := 12344444; isValid(val, true) {
		t.Errorf("Expected %d not to be valid", val)
	}
}

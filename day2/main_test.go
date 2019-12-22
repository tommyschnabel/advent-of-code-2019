package main

import "testing"

func TestExecute1(t *testing.T) {
	ops := []int{1,0,0,0,99}
	err := execute(&ops)
	if err != nil {
		t.Error(err)
	}

	expected := []int{2,0,0,0,99}
	if !opsEquals(&ops, &expected) {
		t.Errorf("ops was supposed to be %v, but was instead %v", expected, ops)
	}
}

func TestExecute2(t *testing.T) {
	ops := []int{2,3,0,3,99}
	err := execute(&ops)
	if err != nil {
		t.Error(err)
	}

	expected := []int{2,3,0,6,99}
	if !opsEquals(&ops, &expected) {
		t.Errorf("ops was supposed to be %v, but was instead %v", expected, ops)
	}
}

func TestExecute3(t *testing.T) {
	ops := []int{2,4,4,5,99,0}
	err := execute(&ops)
	if err != nil {
		t.Error(err)
	}

	expected := []int{2,4,4,5,99,9801}
	if !opsEquals(&ops, &expected) {
		t.Errorf("ops was supposed to be %v, but was instead %v", expected, ops)
	}
}

func TestExecute4(t *testing.T) {
	ops := []int{1,1,1,4,99,5,6,0,99}
	err := execute(&ops)
	if err != nil {
		t.Error(err)
	}

	expected := []int{30,1,1,4,2,5,6,0,99}
	if !opsEquals(&ops, &expected) {
		t.Errorf("ops was supposed to be %v, but was instead %v", expected, ops)
	}
}

func opsEquals(expected *[]int, ops *[]int) bool {
	if len(*expected) != len(*ops) {
		return false
	}

	for i := 0; i < len(*expected); i++ {
		if (*expected)[i] != (*ops)[i] {
			return false
		}
	}

	return true
}

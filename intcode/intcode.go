package main

import (
	"errors"
	"fmt"
)

const (
	opAdd      = 1
	opMultiply = 2
	opWriteInput = 3
	opWriteOutput = 4
	opStop     = 99

	opAddMultiplyWidth = 4
	opWriteInputWidth = 2

	opAddMultiplyInput1    = 1
	opAddMultiplyInput2    = 2
	opAddMultiplyOutput    = 3
	opWriteInputOutputAddr = 1

)

var (
	input = make(chan int)
	output = make(chan int)
)

func execute(opsPtr *[]int) error {
	var pc int

	if opsPtr == nil {
		return errors.New("ops was nil")
	}

	ops := *opsPtr

	for true {
		if pc > len(ops) {
			return errors.New("PC somehow got moved past end of ops")
		}

		//TODO Implement parameter modes and parameter addressing cleanup in general
		paramMode := ops[pc] / 100
		opCode := ops[pc] % 100

		switch  opCode {
		case opAdd, opMultiply:

			inputAddr1 := ops[pc +opAddMultiplyInput1]
			inputAddr2 := ops[pc +opAddMultiplyInput2]
			outputAddr := ops[pc +opAddMultiplyOutput]

			var result int
			if ops[pc] == opAdd {
				result = ops[inputAddr1] + ops[inputAddr2]
			} else { // opMultiply
				result = ops[inputAddr1] * ops[inputAddr2]
			}

			ops[outputAddr] = result
			pc += opAddMultiplyWidth

		case opWriteInput:
			outputAddr := ops[pc +opWriteInputOutputAddr]
			ops[outputAddr] = <-input

		case opWriteOutput:
			outputAddr := ops[pc +opWriteInputOutputAddr]
			output <- ops[outputAddr]

		case opStop:
			return nil
		default:
			return fmt.Errorf("unknown op code %d", ops[pc])
		}
	}

	return nil
}
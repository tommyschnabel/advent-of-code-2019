package main

import (
	"errors"
	"fmt"
)

const (
	opAdd      = 1
	opMultiply = 2
	opStop     = 99

	opWidth = 4

	opInput1 = 1
	opInput2 = 2
	opOutput = 3
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

		if ops[pc] == opStop {
			return nil
		}

		inputAddr1 := ops[pc + opInput1]
		inputAddr2 := ops[pc + opInput2]
		outputAddr := ops[pc + opOutput]

		switch ops[pc] {
		case opAdd:
			ops[outputAddr] = ops[inputAddr1] + ops[inputAddr2]
		case opMultiply:
			ops[outputAddr] = ops[inputAddr1] * ops[inputAddr2]
		default:
			return fmt.Errorf("unknown op code %d", ops[pc])
		}

		pc += opWidth
	}

	return nil
}
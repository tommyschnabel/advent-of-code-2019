package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

const (
	part2MaxIndex = 100 * 100
)

func main() {
	part1()
	part2()
}

func part2() {
	wantedResult := 19690720
	results := [part2MaxIndex]int{}

	group, _ := errgroup.WithContext(context.Background())

	for i := 0; i < part2MaxIndex; i++ {

		i := i
		group.Go(func() error {
			noun, verb := nounVerbFromIndex(i)
			ops := getInput(noun, verb)
			err := execute(ops)
			if err != nil {
				return err
			}

			results[i] = (*ops)[0]
			return nil
		})
	}

	// Wait on all results to finish
	err := group.Wait()
	if err != nil {
		panic(err)
	}

	for i := 0; i < part2MaxIndex; i++ {
		if results[i] == wantedResult {
			noun, verb := nounVerbFromIndex(i)
			fmt.Printf("Result: noun=%d, verb=%d\n", noun, verb)
			fmt.Printf("Answer: %d\n", (100 * noun) + verb)
			return
		}
	}

	panic("No correct result found(!?)")
}

func nounVerbFromIndex(index int) (int, int) {
	return index / 100, index % 100
}

func part1() {

	ops := getInput(12, 2)
	err := execute(ops)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %v\n", (*ops)[0])
}

func getInput(a, b int) *[]int {
	ops := []int{1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,6,1,19,1,5,19,23,2,6,23,27,1,27,5,31,2,9,31,35,1,5,35,39,2,6,39,43,2,6,43,47,1,5,47,51,2,9,51,55,1,5,55,59,1,10,59,63,1,63,6,67,1,9,67,71,1,71,6,75,1,75,13,79,2,79,13,83,2,9,83,87,1,87,5,91,1,9,91,95,2,10,95,99,1,5,99,103,1,103,9,107,1,13,107,111,2,111,10,115,1,115,5,119,2,13,119,123,1,9,123,127,1,5,127,131,2,131,6,135,1,135,5,139,1,139,6,143,1,143,6,147,1,2,147,151,1,151,5,0,99,2,14,0,0}
	ops[1] = a
	ops[2] = b

	return &ops
}

package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	ints, err := parseInts(input)
	if err != nil {
		return "", err
	}
	sum := ints[0] + ints[1]
	output = strconv.Itoa(sum)
	return output, nil
}

func parseInts(input string) (ints []int, err error) {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return nil, fmt.Errorf("no ints: %w", errorEmptyInput)
	}
	input = strings.ReplaceAll(input, " ", "")
	i := 0
	for j := 0; j < len(input); j++ {
		if c := input[j]; c == ' ' || (j > 0 && c == '-') {
			parsedInt, err := parseInt(input[i:j])
			if err != nil {
				return nil, err
			}
			ints = append(ints, parsedInt)
			i = j
		} else if j == len(input)-1 {
			parsedInt, err := parseInt(input[i : j+1])
			if err != nil {
				return nil, err
			}
			ints = append(ints, parsedInt)
		}
	}
	if len(ints) != 2 {
		return nil, fmt.Errorf("parseInts: %w", errorNotTwoOperands)
	}
	return ints, nil
}

func parseInt(str string) (parsedInt int, err error) {
	num := strings.ReplaceAll(str, "+", "")
	parsedInt, err = strconv.Atoi(num)
	if err != nil {
		return 0, fmt.Errorf("can't parse int %w", err)
	}
	return parsedInt, nil
}

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

	errorIncorrectSequence = errors.New("incorrect sequence")
)

const (
	errorTemplate = "StringSum error: %w"
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
	str := strings.TrimSpace(input)
	if len(str) == 0 {
		err = fmt.Errorf(errorTemplate, errorEmptyInput)
		return
	}

	var items []string = make([]string, 0)
	var i int = strings.IndexAny(str, "-+")
	for i != -1 {
		if i == 0 {
			items = append(items, strings.TrimSpace(str[0:0+1])) // operation
		} else {
			items = append(items, strings.TrimSpace(str[0:0+i])) // operand
			items = append(items, strings.TrimSpace(str[i:i+1])) // operation
		}

		if i != len(str)-1 {
			str = str[i+1:]
		} else {
			str = ""
		}

		i = strings.IndexAny(str, "-+")
	}
	if len(str) > 0 {
		items = append(items, strings.TrimSpace(str)) // operand
	}

	var result int = 0
	var operand int
	var operandCount int = 0
	var prevItem string
	for _, item := range items {
		if item != "-" && item != "+" {
			operand, err = strconv.Atoi(item)

			if err != nil {
				err = fmt.Errorf(errorTemplate, err)
				return
			}

			operandCount++
			if operandCount > 2 {
				err = fmt.Errorf(errorTemplate, errorNotTwoOperands)
				return
			}

			if len(prevItem) == 0 {
				result += operand
			} else {
				switch prevItem {
				case "-":
					result -= operand
				case "+":
					result += operand
				default:
					err = fmt.Errorf(errorTemplate, errorIncorrectSequence)
					return
				}
			}
		} else {
			if prevItem == "-" || prevItem == "+" {
				err = fmt.Errorf(errorTemplate, errorIncorrectSequence)
				return
			}
		}
		prevItem = item
	}
	if operandCount != 2 {
		err = fmt.Errorf(errorTemplate, errorNotTwoOperands)
		return
	}

	output = strconv.Itoa(result)

	return
}

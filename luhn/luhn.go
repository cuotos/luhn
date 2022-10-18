package luhn

import (
	"errors"
)

var (
	errZeroInput = errors.New("no input provided")
)

func IsValid(input int) (bool, error) {

	if input == 0 {
		return false, errZeroInput
	}

	// diving int by 10 will drop the least significant number
	//   1234 -> 123.4 -> 123
	numberWithoutChecksum := input / 10

	// mod the input by 10 to get the least significant number
	//  123 % 10 = 3
	providedLuhn := input % 10

	// calculate our own check digit
	calculatedCheck := checksum(numberWithoutChecksum)

	return (calculatedCheck+providedLuhn)%10 == 0, nil
}

func calculateLuhn(input int) int {
	check := checksum(input)

	if check == 0 {
		return 0
	}

	return 10 - check
}

func checksum(input int) int {
	var checksum int
	for i := 0; input > 0; i++ {
		cur := input % 10

		if i%2 == 0 { //even
			if cur > 4 {
				checksum += (cur * 2) - 9
			} else {
				checksum += cur * 2
			}
		} else {
			checksum += cur
		}
		input = input / 10
	}

	return checksum % 10
}

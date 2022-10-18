package luhn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLuhnValidator(t *testing.T) {
	tcs := []struct {
		input       int
		expectValid bool
	}{
		{18, true},
		{5555555555554444, true},
		{5555555555554445, false},
		{1, false},
		//{00, true}, //TODO: cant handle this...
		{01, false},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%d-%v", tc.input, tc.expectValid), func(t *testing.T) {
			actual, _ := IsValid(tc.input)
			assert.Equal(t, tc.expectValid, actual)
		})
	}
}

func TestChecksum(t *testing.T) {
	tcs := []struct {
		input         int
		expectedCheck int
	}{
		{12345, 5},
		{1, 2},
		{98261360165824495, 8},
	}
	for _, tc := range tcs {
		actual := checksum(tc.input)
		assert.Equal(t, tc.expectedCheck, actual)
	}
}

func TestCalculateLuhn(t *testing.T) {
	tcs := []struct {
		input        int
		expectedLuhn int
	}{
		{111, 5},
		{12121212121212121, 6},
		{98261360165824495, 2},
		{12345, 5},
	}
	for _, tc := range tcs {
		actual := calculateLuhn((tc.input))
		assert.Equal(t, tc.expectedLuhn, actual)
	}
}

func TestErrors(t *testing.T) {
	tcs := []struct {
		Input         int
		ExpectedError string
	}{
		{0, "no input provided"},
	}

	for _, tc := range tcs {
		_, actual := IsValid(tc.Input)
		assert.ErrorContains(t, actual, tc.ExpectedError)
	}
}

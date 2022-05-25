package string_sum

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
)

type args struct {
	name     string
	input    string
	expected string
}

func TestStringSum(t *testing.T) {
	cases := []args{
		args{
			name:     "positive",
			input:    " 3 +5 ",
			expected: "8",
		},
		args{
			name:     "negative",
			input:    "-3 -5 ",
			expected: "-8",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if result, err := StringSum(c.input); result != c.expected && err == nil {
				t.Errorf("StringSum() = %v, expected %v", result, c.expected)
			}
		})
	}
}

func TestInvalidInput(t *testing.T) {
	input := "-3c -5"
	expected := ""
	expectedErr := &strconv.NumError{
		Func: "Atoi",
		Num:  "-3c",
		Err:  strconv.ErrSyntax,
	}

	result, err := StringSum(input)

	if !reflect.DeepEqual(result, expected) || err == nil || !reflect.DeepEqual(errors.Unwrap(err), expectedErr) {
		t.Errorf("err = %v, expected err %v", errors.Unwrap(err), expectedErr)
	}
}

func TestThreeOperands(t *testing.T) {
	input := "-1 -2 -3"
	expected := ""
	expectedErrMessage := "expecting two operands, but received more or less"

	result, err := StringSum(input)

	if !reflect.DeepEqual(result, expected) || err == nil || errors.Unwrap(err).Error() != expectedErrMessage {
		t.Errorf("err %v, expected err message %v", errors.Unwrap(err).Error(), expectedErrMessage)
	}
}

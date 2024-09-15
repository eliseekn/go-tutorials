package operator_test

import (
	"calculator/operator"
	"testing"
)

func Test_Can_Sum(t *testing.T) {
	result := operator.Sum([]string{"", "1", "+", "2"})
	if result != 3 {
		t.FailNow()
	}
}

func Test_Can_Substract(t *testing.T) {
	result := operator.Substract([]string{"", "5", "-", "2"})
	if result != 3 {
		t.FailNow()
	}
}

func Test_Can_Multiply(t *testing.T) {
	result := operator.Multiply([]string{"", "3", "x", "3"})
	if result != 9 {
		t.FailNow()
	}
}

func Test_Can_Divide(t *testing.T) {
	result, err := operator.Divide([]string{"", "4", "/", "2"})
	if err != nil {
		t.Fatal(err)
	}

	if result != 2 {
		t.FailNow()
	}
}

func Test_Can_Not_Divide_By_Zero(t *testing.T) {
	_, err := operator.Divide([]string{"", "4", "/", "0"})

	if err == nil {
		t.FailNow()
	}
}

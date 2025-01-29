package main

import (
	"testing"

	"github.com/kermage/GO-Mods/ratement"
)

func TestFlow(t *testing.T) {
	rater := ratement.NewRater(map[int]int{
		5:  60,
		10: 180,
		20: 480,
	})

	testAmount := 30
	expectedValue := 660

	if testValue := rater.Value(testAmount); testValue != expectedValue {
		t.Errorf("Value(%d) = %v, want %d", testAmount, testValue, expectedValue)
	}

	rater.Set(1, 10) // Prepend; no change

	if testValue := rater.Value(testAmount); testValue != expectedValue {
		t.Errorf("Value(%d) = %v, want %d", testAmount, testValue, expectedValue)
	}

	rater.Set(50, 1440) // Append; no change

	if testValue := rater.Value(testAmount); testValue != expectedValue {
		t.Errorf("Value(%d) = %v, want %d", testAmount, testValue, expectedValue)
	}

	rater.Set(20, 500) // Replace

	expectedValue = 680

	if testValue := rater.Value(testAmount); testValue != expectedValue {
		t.Errorf("Value(%d) = %v, want %d", testAmount, testValue, expectedValue)
	}

	rater.Delete(10) // new calculation

	if testValue := rater.Value(testAmount); testValue == expectedValue {
		t.Errorf("Value(%d) still got old value %d", testAmount, expectedValue)
	}

	rater.Set(25, 560) // Insert

	if testValue := rater.Value(testAmount); testValue == expectedValue {
		t.Errorf("Value(%d) still got old value %d", testAmount, expectedValue)
	}

	rater.Set(2, 25) // Insert

	expectedValue = 620

	if testValue := rater.Value(testAmount); testValue != expectedValue {
		t.Errorf("Value(%d) = %v, want %d", testAmount, testValue, expectedValue)
	}
}

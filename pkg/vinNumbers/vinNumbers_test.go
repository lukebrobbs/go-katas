package vinnumbers

import "testing"

func TestVinNumbers(t *testing.T) {
	t.Run("Should return false if not enough digits", func(t *testing.T) {
		expected := false
		actual := vinNumbers("145")
		if expected != actual {
			t.Fatalf("expected %v, but got %v", expected, actual)
		}
	})
	t.Run("Should return true if a valid length VIN is passed", func(t *testing.T) {
		expected := true
		actual := vinNumbers("0471958692")
		if expected != actual {
			t.Fatalf("expected %v, but got %v", expected, actual)
		}
	})
}

func TestValidateCheckDigit(t *testing.T) {
	t.Run("Should return true if check digit is valid", func(t *testing.T) {
		expected := true
		actual := validateCheckDigit("0471958692")
		if expected != actual {
			t.Fatalf("expected %v, but got %v", expected, actual)
		}
	})
	t.Run("Should return true if check digit is valid", func(t *testing.T) {
		expected := false
		actual := validateCheckDigit("0471958694")
		if expected != actual {
			t.Fatalf("expected %v, but got %v", expected, actual)
		}
	})
	t.Run("Another valid VIN", func(t *testing.T) {
		expected := true
		actual := validateCheckDigit("0471606958")
		if expected != actual {
			t.Fatalf("expected %v, but got %v", expected, actual)
		}
	})
}

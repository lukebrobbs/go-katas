package main

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

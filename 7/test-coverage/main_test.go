package main

import "testing"

func TestSomma(t *testing.T) {
	var tests = []struct {
		inputA int
		inputB int
		result int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{-1, -1, -2},
		{0, 0, 0},
	}

	for _, test := range tests {
		out := Somma(test.inputA, test.inputB)
		if out != test.result {
			t.Errorf("Fallito perché mi aspettavo %d e è arrivato %d", test.result, out)
		}
	}

}

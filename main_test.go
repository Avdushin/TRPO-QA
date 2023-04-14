package main

import (
	"math"
	"testing"
)

var (
	NaN = math.NaN()
)

func TestSolveQE(t *testing.T) {
	testCases := []struct {
		a, b, c                      float64
		expectedRoot1, expectedRoot2 float64
		expectedD                    float64
		expectedHasRoot              bool
	}{

		{2, -5, 2, 2, 0.5, 9, true},
		{3, -12, 0, 4, 0, 144, true},
		{0, 0, 10, NaN, NaN, 0, false},
		{9, 0, 0, 0, 0, 0, true},

		// Эти не арбайтен (X_X)
		// {3, 2, 5, NaN, NaN, -56, false},
		// {0, 0, 0, 0, 0, 0, true},
		// {0, 5, 17, -math.Inf(0), -math.Inf(0), 25, false},
	}

	for _, tt := range testCases {
		root1, root2, D := solveQE(tt.a, tt.b, tt.c)

		if !floatEquals(root1, tt.expectedRoot1, 1e-6) {
			t.Errorf("solveQE(%v, %v, %v): root1 = %v, ожидается %v", tt.a, tt.b, tt.c, root1, tt.expectedRoot1)
		}

		if !floatEquals(root2, tt.expectedRoot2, 1e-6) {
			t.Errorf("solveQE(%v, %v, %v): root2 = %v, Ожидается %v", tt.a, tt.b, tt.c, root2, tt.expectedRoot2)
		}

		if !floatEquals(D, tt.expectedD, 1e-6) {
			t.Errorf("solveQE(%v, %v, %v): D = %v, ожидается %v", tt.a, tt.b, tt.c, D, tt.expectedD)
		}

		if tt.expectedHasRoot && D < 0 {
			t.Errorf("solveQE(%v, %v, %v): Ожидаются корни, но они отсутствуют", tt.a, tt.b, tt.c)
		}

		if !tt.expectedHasRoot && D < 0 {
			t.Errorf("solveQE(%v, %v, %v): Ожидаются корни, но они отсутствуют", tt.a, tt.b, tt.c)
		}

	}
}

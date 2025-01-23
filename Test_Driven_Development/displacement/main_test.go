package main

import (
	"math"
	"testing"
)

// Test GenDisplaceFn
func TestGenDisplaceFn(t *testing.T) {
	tests := []struct {
		a, v0, s0, t, expectedDisplacement float64
	}{
		{9.8, 0, 0, 1, 4.9}, // Standard acceleration due to gravity
		{0, 10, 0, 2, 20},   // No acceleration, moving with constant velocity
		{2, 0, 0, 3, 18},    // Positive acceleration, starting from rest expected value is wrong here i.e 18 the value will be 9. This is done to check whether it is working fine or not
		{1, -2, 5, 4, 13},   // Negative initial velocity, positive acceleration. expected value is wrong here i.e 13 the value will be 5. This is done to check whether it is working fine or not
	}

	for _, tt := range tests {
		t.Run("Testing GenDisplaceFn", func(t *testing.T) {
			dispFn := GenDisplaceFn(tt.a, tt.v0, tt.s0)
			result := dispFn(tt.t)

			if math.Abs(result-tt.expectedDisplacement) > 1e-6 { // Allowing for small floating point errors
				t.Errorf("GenDisplaceFn(%v, %v, %v, %v) = %v; expected %v", tt.a, tt.v0, tt.s0, tt.t, result, tt.expectedDisplacement)
			}
		})
	}
}

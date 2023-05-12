package main

import (
	"testing"
)

// TestCalculateInterest tests the calculateInterest function with different balances
func TestCalculateInterest(t *testing.T) {
	tests := []struct {
		name    string
		balance float64
		want    float64
	}{
		{
			name:    "balance < $1,000",
			balance: 500,
			want:    5,
		},
		{
			name:    "$1,000 <= balance < $5,000",
			balance: 2500,
			want:    37.5,
		},
		{
			name:    "$5,000 <= balance < $10,000",
			balance: 7500,
			want:    175,
		},
		{
			name:    "$10,000 <= balance < $50,000",
			balance: 25000,
			want:    825,
		},
		{
			name:    "$50,000+",
			balance: 100000,
			want:    3000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateInterest(tt.balance)
			if got != tt.want {
				t.Errorf("calculateInterest() = %.2f, want %.2f", got, tt.want)
			}
		})
	}
}

// calculateInterest calculates the amount of interest earned for a given balance
// based on the balance bands and interest rates defined above
func calculateInterest(balance float64) float64 {
	interest := 0.0
	for i := 0; i < len(balanceBands); i++ {
		// If the balance is less than the current band, we've found the band
		// the balance belongs to and can break out of the loop
		if balance < balanceBands[i] {
			interest += balance * interestRates[i]
			break
		} else {
			// If the balance is greater than or equal to the current band, we add
			// the maximum amount of balance for that band to the interest calculation
			// and subtract the current band's balance from the balance to continue
			// calculating interest for the remaining balance bands
			interest += balanceBands[i] * interestRates[i]
			balance -= balanceBands[i]
		}
	}
	// Add the remaining balance to the interest calculation using the highest
	// interest rate
	interest += balance * interestRates[len(interestRates)-1]
	return interest
}

func main() {
	// Run the tests
	tests := []testing.InternalTest{{"TestCalculateInterest", TestCalculateInterest}}
	match := testing.MatchTest
	if !testing.RunTests(match, tests) {
		panic("tests failed")
	}
}


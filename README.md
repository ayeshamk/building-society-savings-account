# building-society-savings-account

This program calculates the amount of interest earned for a given balance in a Building Society savings account, based on the balance bands and interest rates defined below:

- < $1,000: 1%
- $1,000 - < $5,000: 1.5%
- $5,000 - < $10,000: 2%
- $10,000 - < $50,000: 2.5%
- $50,000+: 3%

## Requirements
To run this program, you'll need:

Go 1.16 or later

## Installation
- Clone the repository: git clone https://github.com/ayeshamk/building-society-savings-account.git
- Navigate to the project directory: cd building-society-savings-account/golang

## Usage

- To run the tests, execute the following command in the terminal: 

```
go test -v
```

This will run the tests defined in the calculate_interest_test.go file and output the results to the terminal.

- To use the calculateInterest function in your own code, import the main package and call the function with a balance argument:

```
package main

import (
    "fmt"
)

func main() {
    balance := 1000.0
    interest := calculateInterest(balance)
    fmt.Printf("Balance: $%.2f\nInterest: $%.2f\n", balance, interest)
}
```

Output:

```
Balance: $1000.00
Interest: $10.00
```

## Assumptions

- The interest rates are calculated based on the balance bands provided and do not change over time or with any other factors.
- The balance bands are inclusive of their lower bound but exclusive of their upper bound (i.e., $1,000 is included in the $1,000 - < $5,000 band but not in the < $1,000 band).
- The interest is calculated annually, but the program assumes that the balance remains constant for the entire year.
- The program assumes that the input balance is always a positive float64 number.

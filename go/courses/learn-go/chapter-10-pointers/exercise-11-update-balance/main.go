package main

import (
	"errors"
	"fmt"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(customerPtr *customer, transaction transaction) error {
	if transaction.transactionType != "deposit" && transaction.transactionType != "withdrawal" {
		return errors.New("unknown transaction type")
	}
	if transaction.transactionType == "withdrawal" {
		if (customerPtr).balance >= transaction.amount {
			(customerPtr).balance -= transaction.amount
			return nil
		}
		return errors.New("insufficient funds")
	}
	(customerPtr).balance += transaction.amount
	return nil
}

func main() {
	withSubmit := true // Set to false to only run runCases

	type testCase struct {
		name            string
		initialCustomer customer
		transaction     transaction
		expectedBalance float64
		expectError     bool
		errorMessage    string
	}

	runCases := []testCase{
		{
			name:            "Deposit operation",
			initialCustomer: customer{id: 1, balance: 100.0},
			transaction:     transaction{customerID: 1, amount: 50.0, transactionType: transactionDeposit},
			expectedBalance: 150.0,
			expectError:     false,
		},
		{
			name:            "Withdrawal operation",
			initialCustomer: customer{id: 2, balance: 200.0},
			transaction:     transaction{customerID: 2, amount: 100.0, transactionType: transactionWithdrawal},
			expectedBalance: 100.0,
			expectError:     false,
		},
	}

	submitCases := append(runCases, []testCase{
		{
			name:            "insufficient funds for withdrawal",
			initialCustomer: customer{id: 3, balance: 50.0},
			transaction:     transaction{customerID: 3, amount: 100.0, transactionType: transactionWithdrawal},
			expectedBalance: 50.0,
			expectError:     true,
			errorMessage:    "insufficient funds",
		},
		{
			name:            "unknown transaction type",
			initialCustomer: customer{id: 4, balance: 100.0},
			transaction:     transaction{customerID: 4, amount: 50.0, transactionType: "unknown"},
			expectedBalance: 100.0,
			expectError:     true,
			errorMessage:    "unknown transaction type",
		},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	// Simulate *testing.T for main
	for _, test := range testCases {
		// Copy customer to avoid data race/mutation
		cust := test.initialCustomer
		err := updateBalance(&cust, test.transaction)
		failureMessage := ""

		if (err != nil) != test.expectError {
			failureMessage += "Unexpected error presence: expected an error but didn't get one, or vice versa.\n"
		} else if err != nil && err.Error() != test.errorMessage {
			failureMessage += "Incorrect error message.\n"
		}

		if cust.balance != test.expectedBalance {
			failureMessage += "Balance not updated as expected.\n"
		}

		if failureMessage != "" {
			failCount++
			failureMessage = "FAIL\n" + failureMessage +
				"Transaction: " + string(test.transaction.transactionType) +
				fmt.Sprintf(", Amount: %.2f\n", test.transaction.amount) +
				fmt.Sprintf("Expected balance: %.2f, Actual balance: %.2f", test.expectedBalance, cust.balance)
			fmt.Printf(`---------------------------------
%s
`, failureMessage)
		} else {
			passCount++
			successMessage := "PASSED\n" +
				"Transaction: " + string(test.transaction.transactionType) +
				fmt.Sprintf(", Amount: %.2f\n", test.transaction.amount) +
				fmt.Sprintf("Expected balance: %.2f, Actual balance: %.2f", test.expectedBalance, cust.balance)
			fmt.Printf(`---------------------------------
%s
`, successMessage)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

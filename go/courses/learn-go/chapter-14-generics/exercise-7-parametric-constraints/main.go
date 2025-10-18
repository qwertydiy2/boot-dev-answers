package main

import (
	"fmt"
	"testing"
)

type biller[C customer] interface {
	Charge(C) bill
	Name() string
}

// don't edit below this line

type userBiller struct {
	Plan string
}

func (ub userBiller) Charge(u user) bill {
	amount := 50.0
	if ub.Plan == "pro" {
		amount = 100.0
	}
	return bill{
		Customer: u,
		Amount:   amount,
	}
}

func (sb userBiller) Name() string {
	return fmt.Sprintf("%s user biller", sb.Plan)
}

type orgBiller struct {
	Plan string
}

func (ob orgBiller) Name() string {
	return fmt.Sprintf("%s org biller", ob.Plan)
}

func (ob orgBiller) Charge(o org) bill {
	amount := 2000.0
	if ob.Plan == "pro" {
		amount = 3000.0
	}
	return bill{
		Customer: o,
		Amount:   amount,
	}
}

type customer interface {
	GetBillingEmail() string
}

type bill struct {
	Customer customer
	Amount   float64
}

type user struct {
	UserEmail string
}

func (u user) GetBillingEmail() string {
	return u.UserEmail
}

type org struct {
	Admin user
	Name  string
}

func (o org) GetBillingEmail() string {
	return o.Admin.GetBillingEmail()
}

func main() {
	type testCase struct {
		biller         orgBiller
		customer       org
		expectedAmount float64
		expectedEmail  string
	}

	runCases := []testCase{
		{
			biller: orgBiller{Plan: "pro"},
			customer: org{
				Admin: user{UserEmail: "jaskier@oxenfurt.com"},
				Name:  "Oxenfurt",
			},
			expectedAmount: 3000,
			expectedEmail:  "jaskier@oxenfurt.com",
		},
		{
			biller: orgBiller{Plan: "basic"},
			customer: org{
				Admin: user{UserEmail: "vernon@temeria.com"},
				Name:  "Temeria",
			},
			expectedAmount: 2000,
			expectedEmail:  "vernon@temeria.com",
		},
	}

	submitCases := append(runCases, []testCase{
		{
			biller: orgBiller{Plan: "pro"},
			customer: org{
				Admin: user{UserEmail: "fringilla@nilfgaard.com"},
				Name:  "Nilfgaard",
			},
			expectedAmount: 3000,
			expectedEmail:  "fringilla@nilfgaard.com",
		},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passCount := 0
	failCount := 0

	for i, test := range testCases {
		err := testBiller(test.biller, test.customer, test.expectedAmount, test.expectedEmail)
		if err != nil {
			failCount++
			fmt.Printf(`---------------------------------
OrgTest %d Failed:
%v
`, i, err)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
OrgTest %d Passed:
  biller:   %v
  customer: %v
`, i, test.biller, test.customer)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("OrgBilling: %d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("OrgBilling: %d passed, %d failed\n", passCount, failCount)
	}
}

func TestUserBilling(t *testing.T) {
	type testCase struct {
		biller         userBiller
		customer       user
		expectedAmount float64
		expectedEmail  string
	}

	runCases := []testCase{
		{
			biller:         userBiller{Plan: "basic"},
			customer:       user{UserEmail: "vesemir@kaermorhen.com"},
			expectedAmount: 50,
			expectedEmail:  "vesemir@kaermorhen.com",
		},
		{
			biller:         userBiller{Plan: "pro"},
			customer:       user{UserEmail: "zoltan@mahakam.com"},
			expectedAmount: 100,
			expectedEmail:  "zoltan@mahakam.com",
		},
	}

	submitCases := append(runCases, []testCase{
		{
			biller:         userBiller{Plan: "pro"},
			customer:       user{UserEmail: "extra@submit.com"},
			expectedAmount: 100,
			expectedEmail:  "extra@submit.com",
		},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passCount := 0
	failCount := 0

	for i, test := range testCases {
		err := testBiller(test.biller, test.customer, test.expectedAmount, test.expectedEmail)
		if err != nil {
			failCount++
			fmt.Print(`---------------------------------
UserTest %d Failed:
%v
`, i, err)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
UserTest %d Passed:
  biller:   %v
  customer: %v
`, i, test.biller, test.customer)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("UserBilling: %d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("UserBilling: %d passed, %d failed\n", passCount, failCount)
	}
}

func testBiller[C customer](
	b biller[C],
	c C,
	expectedAmount float64,
	expectedEmail string,
) error {
	currentBill := b.Charge(c)
	name := b.Name()

	if currentBill.Amount != expectedAmount ||
		currentBill.Customer.GetBillingEmail() != expectedEmail {
		return fmt.Errorf(`biller "%v" FAILED:
  biller Type:     %T
  customer Type:   %T
  customer:        %v
  expected amount: %v
  expected email:  %v
  actual amount:   %v
  actual email:    %v
`,
			name,
			b,
			c,
			c,
			expectedAmount,
			expectedEmail,
			currentBill.Amount,
			currentBill.Customer.GetBillingEmail(),
		)
	}

	return nil
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getValidInput() CalculationInput {
	return CalculationInput{
		PropertyPrice:           10000.0,
		DownPayment:             3000.0,
		AnnualInterestRate:      0.01,
		AmortizationPeriodYears: 5,
		PaymentSchedule:         Monthly,
	}
}

// let's test only the custom validators here...
func TestCalculateMortgageWithInvalidInputs(t *testing.T) {
	t.Run("payment schedule validation", func(t *testing.T) {
		input := getValidInput()

		for _, v := range []string{"", "bi_weekly", "bi_montly", "monthly"} {
			input.PaymentSchedule = PaymentSchedule(v)
			_, err := CalculateMortgage(input)
			assert.Nil(t, err)
		}

		input.PaymentSchedule = PaymentSchedule("unknownPaymentSchedule")
		_, err := CalculateMortgage(input)
		assert.NotNil(t, err)
	})

	t.Run("amortization period validation", func(t *testing.T) {
		input := getValidInput()

		for _, v := range []int{-1, 0, 1, 35} {
			input.AmortizationPeriodYears = v
			_, err := CalculateMortgage(input)
			assert.NotNil(t, err)
		}

		for _, v := range []int{5, 10, 15, 20, 25, 30} {
			input.AmortizationPeriodYears = v
			_, err := CalculateMortgage(input)
			assert.Nil(t, err)
		}
	})
}

func TestCalculateMortgageCase1(t *testing.T) {
	input := getValidInput()
	input.DownPayment = 0.0

	out, err := CalculateMortgage(input)
	assert.Nil(t, err)
	assert.Equal(t, Monthly, out.PaymentSchedule)
	assert.Equal(t, 10000.0, out.Principal) // do down payment
	assert.Equal(t, 12.0, out.NbPaymentsPerAnnum)
	assert.Equal(t, 60.0, out.NbPaymentsOverAmortization)                      // 12*5=60
	assert.Equal(t, 0.0008333333333333334, out.PerPaymentScheduleInterestRate) // 0.01/12=0.00083333333333333333333
	assert.Equal(t, 170.94, out.Mortgage)
}

func TestCalculateMortgageCase2(t *testing.T) {
	input := getValidInput()

	out, err := CalculateMortgage(input)
	assert.Nil(t, err)
	assert.Equal(t, Monthly, out.PaymentSchedule)
	assert.Equal(t, 7000.0, out.Principal) // 10000-3000
	assert.Equal(t, 12.0, out.NbPaymentsPerAnnum)
	assert.Equal(t, 60.0, out.NbPaymentsOverAmortization)
	assert.Equal(t, 0.0008333333333333334, out.PerPaymentScheduleInterestRate)
	assert.Equal(t, 119.66, out.Mortgage)
}

func TestCalculateMortgageCase3(t *testing.T) {
	input := getValidInput()
	input.AnnualInterestRate = 0.015

	out, err := CalculateMortgage(input)
	assert.Nil(t, err)
	assert.Equal(t, Monthly, out.PaymentSchedule)
	assert.Equal(t, 7000.0, out.Principal)
	assert.Equal(t, 12.0, out.NbPaymentsPerAnnum)
	assert.Equal(t, 60.0, out.NbPaymentsOverAmortization)
	assert.Equal(t, 0.00125, out.PerPaymentScheduleInterestRate)
	assert.Equal(t, 121.17, out.Mortgage)
}

func TestCalculateMortgageCase4(t *testing.T) {
	input := getValidInput()
	input.AnnualInterestRate = 0.015
	input.AmortizationPeriodYears = 10

	out, err := CalculateMortgage(input)
	assert.Nil(t, err)
	assert.Equal(t, Monthly, out.PaymentSchedule)
	assert.Equal(t, 7000.0, out.Principal)
	assert.Equal(t, 12.0, out.NbPaymentsPerAnnum)
	assert.Equal(t, 120.0, out.NbPaymentsOverAmortization)
	assert.Equal(t, 0.00125, out.PerPaymentScheduleInterestRate)
	assert.Equal(t, 62.85, out.Mortgage)
}

func TestCalculateMortgageCase5(t *testing.T) {
	input := getValidInput()
	input.AnnualInterestRate = 0.015
	input.AmortizationPeriodYears = 10
	input.PaymentSchedule = BiMonthly

	out, err := CalculateMortgage(input)
	assert.Nil(t, err)
	assert.Equal(t, BiMonthly, out.PaymentSchedule)
	assert.Equal(t, 7000.0, out.Principal)
	assert.Equal(t, 24.0, out.NbPaymentsPerAnnum)
	assert.Equal(t, 240.0, out.NbPaymentsOverAmortization)
	assert.Equal(t, 0.000625, out.PerPaymentScheduleInterestRate)
	assert.Equal(t, 31.42, out.Mortgage)
}

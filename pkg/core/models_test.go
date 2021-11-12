package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNbPaymentsPerAnnum(t *testing.T) {
	t.Run("no PaymentSchedule (default)", func(t *testing.T) {
		s := PaymentSchedule("unknownPaymentSchedule")
		assert.Equal(t, Monthly.NbPaymentsPerAnnum(), s.NbPaymentsPerAnnum())
	})
	t.Run("BiWeekly", func(t *testing.T) {
		assert.Equal(t, 96, BiWeekly.NbPaymentsPerAnnum())
	})
	t.Run("BiMonthly", func(t *testing.T) {
		assert.Equal(t, 24, BiMonthly.NbPaymentsPerAnnum())
	})
	t.Run("Monthly", func(t *testing.T) {
		assert.Equal(t, 12, Monthly.NbPaymentsPerAnnum())
	})
}

func TestNewCalculationOutput(t *testing.T) {
	t.Run("no PaymentSchedule (default)", func(t *testing.T) {
		input := CalculationInput{}
		out := newCalculationOutput(input)
		assert.Equal(t, input, out.CalculationInput)
		assert.Equal(t, Monthly, out.PaymentSchedule)
	})
	t.Run("with input data", func(t *testing.T) {
		input := CalculationInput{
			PropertyPrice:           5000.0,
			DownPayment:             1000.0,
			AnnualInterestRate:      0.01,
			AmortizationPeriodYears: 10,
			PaymentSchedule:         BiMonthly,
		}
		out := newCalculationOutput(input)
		assert.Equal(t, input, out.CalculationInput)
		assert.Equal(t, BiMonthly, out.PaymentSchedule)
	})
}

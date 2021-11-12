package core

import (
	"math"

	"github.com/bcolucci/mortgage-calculator/pkg/utils"
	validator "github.com/go-playground/validator/v10"
)

const mortagePrecision = 2

var validate = validator.New()

func init() {
	validate.RegisterValidation("amortizationPeriod", func(fl validator.FieldLevel) bool {
		return isAmortizationPeriod(fl.Field().Int())
	})
	validate.RegisterValidation("paymentSchedule", func(fl validator.FieldLevel) bool {
		return isVaymentSchedule(fl.Field().String())
	})
}

func CalculateMortgage(input CalculationInput) (*CalculationOutput, error) {
	if err := validate.Struct(input); err != nil {
		return nil, err
	}

	out := newCalculationOutput(input)

	out.Principal = input.PropertyPrice - input.DownPayment
	out.NbPaymentsPerAnnum = float64(out.PaymentSchedule.NbPaymentsPerAnnum())
	out.NbPaymentsOverAmortization = out.NbPaymentsPerAnnum * float64(input.AmortizationPeriodYears)

	out.PerPaymentScheduleInterestRate = input.AnnualInterestRate / out.NbPaymentsPerAnnum

	v := math.Pow(1+out.PerPaymentScheduleInterestRate, out.NbPaymentsOverAmortization)

	out.Mortgage = utils.ToFixed(out.Principal*utils.SafeDiv(out.PerPaymentScheduleInterestRate*v, v-1), mortagePrecision)

	return out, nil
}

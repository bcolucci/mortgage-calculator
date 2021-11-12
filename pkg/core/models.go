package core

type PaymentSchedule string

const (
	BiWeekly  PaymentSchedule = "bi_weekly"
	BiMonthly PaymentSchedule = "bi_montly"
	Monthly   PaymentSchedule = "monthly"
)

func (ps PaymentSchedule) NbPaymentsPerAnnum() int {
	switch ps {
	case BiWeekly:
		return 2 * 4 * 12
	case BiMonthly:
		return 2 * 12
	default:
		return 12
	}
}

type CalculationInput struct {
	PropertyPrice           float64         `validate:"gt=0"`
	DownPayment             float64         `validate:"gte=0"`
	AnnualInterestRate      float64         `validate:"gt=0"`
	AmortizationPeriodYears int             `validate:"amortizationPeriod"`
	PaymentSchedule         PaymentSchedule `validate:"paymentSchedule"`
}

type CalculationOutput struct {
	Principal                      float64
	NbPaymentsPerAnnum             float64
	NbPaymentsOverAmortization     float64
	PerPaymentScheduleInterestRate float64
	Mortgage                       float64
	PaymentSchedule                PaymentSchedule
	CalculationInput               CalculationInput
}

func newCalculationOutput(input CalculationInput) *CalculationOutput {
	out := &CalculationOutput{}
	out.CalculationInput = input

	out.PaymentSchedule = input.PaymentSchedule
	if out.PaymentSchedule == "" {
		out.PaymentSchedule = Monthly
	}

	return out
}

package core

var isVaymentSchedule = func(s string) bool {
	return s == "" ||
		s == string(BiWeekly) ||
		s == string(BiMonthly) ||
		s == string(Monthly)
}

var isAmortizationPeriod = func(p int64) bool {
	return p%5 == 0 && p > 0 && p <= 30
}

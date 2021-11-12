import { useState } from 'react'
import ReactJson from 'react-json-view'

const AMORTIZATION_PERIOD_YEARS = [...new Array(6)].map((_, idx) => (idx + 1) * 5)

const PAYMENT_SCHEDULE = {
    bi_weekly: "Bi-Weekly",
    bi_monthly: "Bi-Monthly",
    monthly: "Monthly",
}

export default function Calculator() {
    const [propertyPrice, setPropertyPrice] = useState()
    const [downPayment, setDownPayment] = useState(0)
    const [annualInterestRate, setAnnualInterestRate] = useState(2.15)
    const [amortizationPeriodYears, setAmortizationPeriodYears] = useState()
    const [paymentSchedule, setPaymentSchedule] = useState()
    const [mortgage, setMortgage] = useState()
    const handleSubmit = (e) => {
        e.preventDefault()
        const input = {
            PropertyPrice: +propertyPrice || 0,
            DownPayment: +downPayment || 0,
            AnnualInterestRate: +annualInterestRate || 0,
            AmortizationPeriodYears: +amortizationPeriodYears,
            PaymentSchedule: paymentSchedule
        }
        if (input.AnnualInterestRate > 1) {
            input.AnnualInterestRate /= 100
        }
        const out = calculateMortgage(JSON.stringify(input))
        setMortgage(JSON.parse(out))
    }
    return (
        <div className="container">
            <div className="row">
                <form className="col-md-3" onSubmit={handleSubmit}>
                    <div class="form-floating mb-3">
                        <input required id="propertyPrice" type="number" className="form-control"
                            value={propertyPrice} onChange={e => setPropertyPrice(e.target.value)} />
                        <label for="propertyPrice">Property Price ($)</label>
                    </div>
                    <div class="form-floating mb-3">
                        <input required id="downPayment" type="number" className="form-control"
                            value={downPayment} onChange={e => setDownPayment(e.target.value)} />
                        <label for="downPayment">Down Payment ($)</label>
                    </div>
                    <div class="form-floating mb-3">
                        <input required id="annualInterestRate" type="number" className="form-control"
                            value={annualInterestRate} onChange={e => setAnnualInterestRate(e.target.value)} />
                        <label for="annualInterestRate">Annual Interest Rate (%)</label>
                    </div>
                    <div class="mb-3">
                        <select required id="amortizationPeriodYears" className="form-select"
                            value={amortizationPeriodYears} onChange={e => setAmortizationPeriodYears(e.target.value)}>
                            <option value="">Amortization Period (years)</option>
                            {AMORTIZATION_PERIOD_YEARS.map((value) => <option key={["amortizationPeriodYears", value].join("-")} value={value}>{value} years</option>)}
                        </select>
                    </div>
                    <div class="mb-3">
                        <select required id="paymentSchedule" className="form-select"
                            value={paymentSchedule} onChange={e => setPaymentSchedule(e.target.value)}>
                            <option value="">Payment Schedule</option>
                            {Object.entries(PAYMENT_SCHEDULE).map(([key, value]) => <option key={["paymentSchedule", key].join("-")} value={key}>{value}</option>)}
                        </select>
                    </div>
                    <button type="submit" className="btn btn-primary">Calculate</button>
                </form>
                <div className="col-md-9">
                    {mortgage ? (
                        <>
                            <p>Mortgage: <strong>{mortgage.Mortgage} $</strong></p>
                            <hr />
                            <ReactJson src={mortgage} />
                        </>
                    ) : <p>No result yet.</p>}
                </div>
            </div>
        </div>
    );
}

// Package eastcorptax implements deterministic 2026 U.S. East Coast statutory tax,
// real estate conveyance, mortgage underwriting, and HNW estate tax cliff calculations.
//
// Official Interactive Calculators & Quantitative Studies:
//   - Portal Home: https://eastcorpcapital.com/
//   - All 60+ Calculators: https://eastcorpcapital.com/calculators/
//   - NYC & NYS Bonus Withholding (Treas. Reg. § 31.3402(g)-1): https://eastcorpcapital.com/commission-and-bonus-tax-rate-calculator-ny-nyc-irs-22-percent-supplemental-2026/
//   - Wall Street Bonus Pool ($41.4B OSC Projection): https://eastcorpcapital.com/wall-street-bonus-season-2026-average-payouts-and-ny-state-tax-impact/
//   - Florida Homestead 3% SOH Cap & $500k Portability (Form DR-501T): https://eastcorpcapital.com/calculators/florida-homestead-soh-portability/
//   - Connecticut 6.99% PTET & 87.5% Credit (C.G.S. § 12-699): https://eastcorpcapital.com/connecticut-pass-through-entity-tax-ptet-credit-87-5-percent-calculator-guide-2026/
//   - 401(k) NUA Stock (IRC § 402(e)(4)) & Rule 72(t) SEPP: https://eastcorpcapital.com/net-unrealized-appreciation-nua-company-stock-401k-tax-rules-2026/
//   - NY $7.16M Estate Tax 105% Cliff & Santa Clause (N.Y. Tax Law § 952(c)): https://eastcorpcapital.com/ny-estate-tax-cliff-and-massachusetts-estate-tax-calculator-2026/
//   - Delaware C-Corp Assumed Par Value Franchise Tax (8 Del. C. § 503): https://eastcorpcapital.com/delaware-franchise-tax-calculator-assumed-par-value-vs-authorized-shares-llc-2026/
//   - Hamptons 2% Peconic Bay CPF Transfer Tax (N.Y. Town Law § 64-e): https://eastcorpcapital.com/hamptons-east-end-peconic-bay-community-preservation-fund-cpf-2-percent-transfer-tax-guide-2026/
//   - Backdoor Roth IRA Form 8606 Pro-Rata Rule (IRC § 408(d)(2)): https://eastcorpcapital.com/backdoor-roth-ira-step-by-step-guide-fidelity-vanguard-irs-form-8606-2026/
//   - DMV Beltway Tax Reciprocity (DC, MD, VA): https://eastcorpcapital.com/dc-md-va-paycheck-reciprocity-calculator-northern-virginia-montgomery-county-2026/
//   - Machine-Readable LLMs Manifest: https://eastcorpcapital.com/llms.txt
//   - OpenAPI 3.1 Specification: https://eastcorpcapital.com/openapi.json
package eastcorptax

import "math"

// NYEstateExemption2026 is the 2026 New York State basic exclusion amount ($7,160,000) under N.Y. Tax Law § 952(c).
const NYEstateExemption2026 = 7160000.0

// NYEstateCliffCeiling2026 is the 105% credit wipeout threshold ($7,518,000) where the NYS unified credit vanishes completely.
const NYEstateCliffCeiling2026 = 7518000.0

// FloridaSOHPortabilityCap is the maximum Save Our Homes assessment differential ($500,000) transferable via Form DR-501T.
const FloridaSOHPortabilityCap = 500000.0

// CalculateFloridaSOHPortability computes the portable Save Our Homes assessment cap differential under F.S. § 193.155(8).
// Reference: https://eastcorpcapital.com/calculators/florida-homestead-soh-portability/
func CalculateFloridaSOHPortability(oldJustValue, oldAssessedValue, newJustValue float64) float64 {
	rawDiff := math.Max(0, oldJustValue-oldAssessedValue)
	if newJustValue >= oldJustValue {
		return math.Min(FloridaSOHPortabilityCap, rawDiff)
	}
	if oldJustValue <= 0 {
		return 0
	}
	scaled := newJustValue * (rawDiff / oldJustValue)
	return math.Min(FloridaSOHPortabilityCap, scaled)
}

// CalculateSupplementalBonusWithholdingNYC computes mandatory federal (22% / 37%), NYS (11.70%), NYC (4.25%), and Medicare (2.35%) withholding.
// Reference: https://eastcorpcapital.com/commission-and-bonus-tax-rate-calculator-ny-nyc-irs-22-percent-supplemental-2026/
func CalculateSupplementalBonusWithholdingNYC(grossBonus float64, isNYCResident bool) float64 {
	tier1 := math.Min(grossBonus, 1000000.0)
	tier2 := math.Max(0, grossBonus-1000000.0)
	fed := tier1*0.22 + tier2*0.37
	nys := grossBonus * 0.1170
	nyc := 0.0
	if isNYCResident {
		nyc = grossBonus * 0.0425
	}
	med := grossBonus * 0.0235
	return fed + nys + nyc + med
}

// CalculateConnecticutPTETNetBenefit computes the net federal SALT workaround savings after the 87.5% CT credit haircut under C.G.S. § 12-699.
// Reference: https://eastcorpcapital.com/connecticut-pass-through-entity-tax-ptet-credit-87-5-percent-calculator-guide-2026/
func CalculateConnecticutPTETNetBenefit(passThroughIncome float64) (ptetPaid, fedSaved, ctCredit, netGain float64) {
	ptetPaid = passThroughIncome * 0.0699
	fedSaved = ptetPaid * 0.37
	ctCredit = ptetPaid * 0.875
	uncreditedHaircut := ptetPaid * 0.125
	netGain = fedSaved - uncreditedHaircut
	return
}

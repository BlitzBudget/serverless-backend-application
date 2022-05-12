package models

// Create struct to hold info about new item
type QueryParameter struct {
	InstalledPotency                  *string `json:"installed_potency"`
	SalePrice                         *string `json:"sale_price"`
	ReturnOfInvestment                *string `json:"return_of_investment"`
	PotentialSavings                  *string `json:"potential_savings"`
	Quota20Year                       *string `json:"quota_20_year"`
	Quota10Year                       *string `json:"quota_10_year"`
	Quota5Year                        *string `json:"quota_5_year"`
	AnnualProduction                  *string `json:"annual_production"`
	EstimatedAutoconsumption          *string `json:"estimated_autoconsumption"`
	EnergySaved                       *string `json:"energy_saved"`
	PotentialSavingsFirstYear         *string `json:"potential_savings_first_year"`
	PotentialSavingsFirstYearFinanced *string `json:"potential_savings_first_year_financed"`
	PotentialSavingsLP                *string `json:"potential_savings_lp"`
	PotentialSavingsLPFinanced        *string `json:"potential_savings_lp_financed"`
	CO2Emitted                        *string `json:"co2_emitted"`
	Client                            *string `json:"client"`
	Beneficiary                       *string `json:"beneficiary"`
	BeneficiaryAddressOne             *string `json:"beneficiary_address_one"`
	BeneficiaryAddressTwo             *string `json:"beneficiary_address_two"`
	BeneficiaryBankAccount            *string `json:"beneficiary_bank_account"`
	Date                              *string `json:"date"`
	PreventInfiniteLoop               *string `json:"prevent_infinite_loop,omitempty"`
}

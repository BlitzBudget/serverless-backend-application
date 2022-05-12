package repository

import (
	"add-transactions/service/models"
	"testing"
)

var body string
var validAnswer string
var queryParameter *models.QueryParameter

func setup() {
	installedPotency := "installed_potency"
	salePrice := "sale_price"
	returnOfInvestment := "return_of_investment"
	potentialSavings := "potential_savings"
	quota20Year := "quota_20_year"
	quota10Year := "quota_10_year"
	quota5Year := "quota_5_year"
	annualProduction := "annual_production"
	estimatedAutoconsumption := "estimated_autoconsumption"
	energySaved := "energy_saved"
	potentialSavingsFirstYear := "potential_savings_first_year"
	potentialSavingsFirstYearFinanced := "potential_savings_first_year_financed"
	potentialSavingsLP := "potential_savings_lp"
	potentialSavingsLPFinanced := "potential_savings_lp_financed"
	cO2Emitted := "co2_emitted"
	client := "client"
	beneficiary := "beneficiary"
	beneficiaryAddressOne := "beneficiary_address_one"
	beneficiaryAddressTwo := "beneficiary_address_two"
	beneficiaryBankAccount := "beneficiary_bank_account"
	date := "date"

	queryParameter = &models.QueryParameter{
		InstalledPotency:                  &installedPotency,
		SalePrice:                         &salePrice,
		ReturnOfInvestment:                &returnOfInvestment,
		PotentialSavings:                  &potentialSavings,
		Quota20Year:                       &quota20Year,
		Quota10Year:                       &quota10Year,
		Quota5Year:                        &quota5Year,
		AnnualProduction:                  &annualProduction,
		EstimatedAutoconsumption:          &estimatedAutoconsumption,
		EnergySaved:                       &energySaved,
		PotentialSavingsFirstYear:         &potentialSavingsFirstYear,
		PotentialSavingsFirstYearFinanced: &potentialSavingsFirstYearFinanced,
		PotentialSavingsLP:                &potentialSavingsLP,
		PotentialSavingsLPFinanced:        &potentialSavingsLPFinanced,
		CO2Emitted:                        &cO2Emitted,
		Client:                            &client,
		Beneficiary:                       &beneficiary,
		BeneficiaryAddressOne:             &beneficiaryAddressOne,
		BeneficiaryAddressTwo:             &beneficiaryAddressTwo,
		BeneficiaryBankAccount:            &beneficiaryBankAccount,
		Date:                              &date,
	}

	body = "{\"installed_potency\" : \"1\",\"sale_price\" : \"2\",\"return_of_investment\": \"3\",\"potential_savings\": \"4\",\"quota_20_year\": \"5\",\"quota_10_year\": \"6\",\"quota_5_year\": \"7\",\"annual_production\": \"8\",\"estimated_autoconsumption\": \"9\",\"energy_saved\": \"10\",\"potential_savings_first_year\": \"11\",\"potential_savings_first_year_financed\": \"12\",\"potential_savings_lp\": \"13\",\"potential_savings_lp_financed\": \"14\",\"co2_emitted\": \"15\",\"client\": \"16\",\"beneficiary\": \"17\",\"beneficiary_address_one\": \"18\",\"beneficiary_address_two\": \"19\",\"beneficiary_bank_account\": \"20\",\"date\": \"21\"}"
	validAnswer = "1"
}

func Test_repository_AttributeBuilder(t *testing.T) {
	setup()
	type args struct {
		body string
	}
	type testCases struct {
		name    string
		args    args
		want    *string
		wantErr bool
	}
	tests := []testCases{
		{
			name: "ProperData",
			args: args{
				body: body,
			},
			want:    &validAnswer,
			wantErr: false,
		},
		{
			name: "InvalidData",
			args: args{
				body: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryParameters, got, err := AttributeBuilder(&tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("AttributeBuilder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && *got["installed_potency"].S != *tt.want {
				t.Errorf("AttributeBuilder() = %v, want %v", got, tt.want)
			}
			if queryParameters.InstalledPotency != nil && *queryParameters.InstalledPotency != *tt.want {
				t.Errorf("AttributeBuilder() = %v, want %v", *queryParameters.InstalledPotency, *tt.want)
			}
		})
	}
}

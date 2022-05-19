package repository

import (
	"get-investment/service/models"
	"testing"
)

var body string
var validAnswer string
var queryParameter *models.QueryParameter

func setup() {
	walletId := "Wallet#2020-05-02T17:19:13.022Z"
	investmentId := "Investment#2020-05-02T17:19:13.022Z"
	investedAmount := int64(20)
	currentValue := int64(20)
	investmentName := "Mutual Funds"

	queryParameter = &models.QueryParameter{
		WalletId:       &walletId,
		StartsWithDate: &investmentId,
		EndsWithDate:   &investmentId,
		InvestedAmount: &investedAmount,
		CurrentValue:   &currentValue,
		InvestmentName: &investmentName,
	}

	body = `{"walletId": "Wallet#2020-05-02T17:19:13.022Z","amount": 95,"investment": "Investment#2020-05-02T17:19:13.022Z","description": "Investment Description","tags": ["Expense", "Travel"]}`
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
			want:    queryParameter.WalletId,
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
			got, err := AttributeBuilder(&tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("AttributeBuilder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				t.Errorf("AttributeBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

package repository

import (
	"testing"
	"update-investment/service/models"
)

var body string
var validAnswer string
var requestModel *models.RequestModel

func setup() {
	pk := "User#2020-05-02T17:19:13.022Z"
	sk := "Investment#2020-05-02T17:19:13.022Z"
	debt := int64(60)
	debtName := "Debt Name"
	debtRepaid := true
	currentValue := int64(30)

	requestModel = &models.RequestModel{
		Pk:           &pk,
		Sk:           &sk,
		Debt:         &debt,
		DebtName:     &debtName,
		DebtRepaid:   &debtRepaid,
		CurrentValue: &currentValue,
	}

	body = `{"userId": "User#2020-05-02T17:19:13.022Z","amount": 95,"investment": "Investment#2020-05-02T17:19:13.022Z","description": "Investment Description","tags": ["Expense", "Travel"]}`
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
			want:    requestModel.Pk,
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
			got := AttributeBuilder(&tt.args.body)
			if got == nil || got.Pk != tt.want {
				t.Errorf("AttributeBuilder() got = %v, want %v", got, tt.want)
				return
			}
		})
	}
}

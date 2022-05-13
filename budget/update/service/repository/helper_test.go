package repository

import (
	"patch-budget/service/models"
	"testing"
)

var body string
var validAnswer string
var requestModel *models.RequestModel

func setup() {
	walletId := "Wallet#2020-05-02T17:19:13.022Z"
	budgetId := "Budget#2020-05-02T17:19:13.022Z"
	planned := int64(95)
	category := "Category#2020-05-02T17:19:13.022Z"

	requestModel = &models.RequestModel{
		Pk:       &walletId,
		Sk:       &budgetId,
		Planned:  &planned,
		Category: &category,
	}

	body = `{"walletId": "Wallet#2020-05-02T17:19:13.022Z","amount": 95,"category": "Category#2020-05-02T17:19:13.022Z","description": "Budget Description","tags": ["Expense", "Travel"]}`
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
				t.Errorf("AttributeBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

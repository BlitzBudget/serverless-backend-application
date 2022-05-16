package repository

import (
	"testing"
	"update-category/service/models"
)

var body string
var validAnswer string
var requestModel *models.RequestModel

func setup() {
	pk := "Category#2020-05-02T17:19:13.022Z"
	sk := "CategoryRule#2020-05-02T17:19:13.022Z"
	transactionName := "Just Eat"

	requestModel = &models.RequestModel{
		Pk:              &pk,
		Sk:              &sk,
		TransactionName: &transactionName,
	}

	body = `{"userId": "Category#2020-05-02T17:19:13.022Z","amount": 95,"category": "Category#2020-05-02T17:19:13.022Z","description": "Category Description","tags": ["Expense", "Travel"]}`
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

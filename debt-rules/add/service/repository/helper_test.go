package repository

import (
	"add-debt-rule/service/models"
	"testing"
)

var body string
var validAnswer string
var queryParameter *models.QueryParameter

func setup() {
	DebtrId := "Debt#2020-05-02T17:19:13.022Z"
	TransactionName := "Groceries"
	creationDate := "2020-05-02T17:19:13.022"
	updatedDate := "2020-05-02T17:19:13.022"
	sk := "DebtRule#2020-05-02T17:19:13.022Z"

	queryParameter = &models.QueryParameter{
		Pk:              &DebtrId,
		Sk:              sk,
		TransactionName: &TransactionName,
		CreationDate:    &creationDate,
		UpdatedDate:     &updatedDate,
	}

	body = `{"debtrId": "Debt#2020-05-02T17:19:13.022Z","amount": 95,"debt": "Debt#2020-05-02T17:19:13.022Z","description": "Transaction Description","tags": ["Expense", "Travel"]}`
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
			want:    queryParameter.Pk,
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

package repository

import (
	"patch-transactions/service/models"
	"testing"
)

var body string
var validAnswer string
var requestModel *models.RequestModel

func setup() {
	walletId := "Wallet#2020-05-02T17:19:13.022Z"
	transactionId := "Transaction#2020-05-02T17:19:13.022Z"
	amount := int64(95)
	creationDate := "2020-05-02T17:19:13.022"
	updatedDate := "2020-05-02T17:19:13.022"
	categoryId := "CategoryId#2020-05-02T17:19:13.022Z"
	tags := []string{"Expense", "Travel"}
	description := "estimated_autoconsumption"

	requestModel = &models.RequestModel{
		Pk:           &walletId,
		Sk:           transactionId,
		Amount:       &amount,
		Description:  &description,
		CreationDate: &creationDate,
		UpdatedDate:  &updatedDate,
		CategoryId:   &categoryId,
		Tags:         &tags,
	}

	body = `{"walletId": "Wallet#2020-05-02T17:19:13.022Z","amount": 95,"category_id": "CategoryId#2020-05-02T17:19:13.022Z","description": "Transaction Description","tags": ["Expense", "Travel"]}`
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

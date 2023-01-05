package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttributeBuilder(t *testing.T) {
	assert := assert.New(t)
	pk := "pk"
	sk := "sk"
	categoryName := "categoryName"
	categoryType := "categoryType"
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"category_name\":\"" + categoryName + "\",\"category_type\":\"" + categoryType + "\"}"

	requestModel, err := AttributeBuilder(&body)

	assert.NoError(err)
	assert.NotNil(requestModel)

	if *requestModel.Pk != pk {
		t.Errorf("pk convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.Pk, pk)
		return
	}

	if *requestModel.Sk != sk {
		t.Errorf("SK convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.Sk, sk)
		return
	}

	if *requestModel.CategoryName != categoryName {
		t.Errorf("CategoryName convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.CategoryName, categoryName)
		return
	}

	if *requestModel.CategoryType != categoryType {
		t.Errorf("CategoryType convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.CategoryType, categoryType)
		return
	}
}

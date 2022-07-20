package repository

import (
	"testing"
)

func TestAttributeBuilder(t *testing.T) {
	pk := "pk"
	sk := "sk"
	categoryName := "categoryName"
	categoryType := "categoryType"
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"category_name\":\"" + categoryName + "\",\"category_type\":\"" + categoryType + "\"}"

	requestModel := AttributeBuilder(&body)

	if requestModel == nil {
		t.Errorf("AttributeBuilder() is null")
		return
	}

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

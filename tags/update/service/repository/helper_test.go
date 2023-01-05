package repository

import (
	"strconv"
	"testing"
)

func TestAttributeBuilder(t *testing.T) {
	pk := "pk"
	sk := "sk"
	investedAmount := 1000
	tagName := "tagName"
	currentAmount := 2000
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"invested_amount\":" + strconv.Itoa(investedAmount) + ",\"tag_name\":\"" + tagName + "\",\"current_value\":" + strconv.Itoa(currentAmount) + "}"

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

	if *requestModel.TagName != tagName {
		t.Errorf("TagName convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.TagName, tagName)
		return
	}
}

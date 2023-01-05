package repository

import (
	"add-tag/service/models"
	"testing"
)

func TestTransformToQueryParamter(t *testing.T) {
	pk := "Wallet#2022-05-12T14:40:29.156Z"
	sk := "Tag#2022-05-12T20:25:19Z"
	creationDate := "2022-05-12T20:25:19Z"
	updatedDate := "2022-05-12T20:25:19Z"
	tagName := "tagName"
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"creation_date\":\"" + creationDate + "\",\"updated_date\":\"" + updatedDate + "\",\"tag_name\":\"" + tagName + "\"}"

	got, err := TransformToQueryParamter(&body)
	if err != nil {
		t.Errorf("TransformToQueryParamter() error = %v", err)
		return
	}

	if got == nil {
		t.Errorf("TransformToQueryParamter() is null")
		return
	}

	if got.TagName == nil {
		t.Errorf("name TagName to DynamoDB attribute not correct, got = %v", got.TagName)
		return
	}

	if got.Pk == nil {
		t.Errorf("name Pk to DynamoDB attribute not correct, got = %v", got.Pk)
		return
	}
}

func TestAttributeBuilder(t *testing.T) {

	tagName := "tagName"
	pk := "Wallet#2022-05-12T14:40:29.156Z"
	creationDate := "2022-05-12T20:25:19Z"
	updatedDate := "2022-05-12T20:25:19Z"
	queryParameter := models.QueryParameter{
		TagName:      &tagName,
		Pk:           &pk,
		CreationDate: &creationDate,
		UpdatedDate:  &updatedDate,
	}

	got, err := AttributeBuilder(&queryParameter)
	if err != nil {
		t.Errorf("TransformToQueryParamter() error = %v", err)
		return
	}

	if *(*got["creation_date"]).S == "" {
		t.Errorf("name creationDate to DynamoDB attribute not correct, got = %v", *(*got["creation_date"]).S)
		return
	}

	if *(*got["updated_date"]).S == "" {
		t.Errorf("name updated_date to DynamoDB attribute not correct, got = %v", *(*got["updated_date"]).S)
		return
	}

	if *(*got["pk"]).S != pk {
		t.Errorf("name pk to DynamoDB attribute not correct, got = %v, want = %v", *(*got["pk"]).S, pk)
		return
	}

	if *(*got["tag_name"]).S != tagName {
		t.Errorf("name tag_name to DynamoDB attribute not correct, got = %v, want = %v", *(*got["tag_name"]).S, tagName)
		return
	}
}

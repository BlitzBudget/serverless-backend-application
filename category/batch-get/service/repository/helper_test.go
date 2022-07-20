package repository

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAttributeBuilder(t *testing.T) {
	userId := "userId"
	var categoryIds [3]string
	categoryIds[0] = "categoryID1"
	categoryIds[1] = "categoryID2"
	categoryIds[2] = "categoryID3"
	categoryIdJSON, err := json.Marshal(categoryIds)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	body := "{\"user_id\":\"" + userId + "\",\"category_ids\":" + string(categoryIdJSON) + "}"

	requestModel, err := AttributeBuilder(&body)

	if err != nil {
		t.Errorf("Err is not null")
		return
	}

	if requestModel == nil {
		t.Errorf("AttributeBuilder() is null")
		return
	}

	if *requestModel.UserId != userId {
		t.Errorf("UserId convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.UserId, userId)
		return
	}

	if (*(*requestModel).CategoryIds)[0] != categoryIds[0] {
		t.Errorf("UserId convertion to DynamoDB attribute not correct, got = %v, want = %v", (*(*requestModel).CategoryIds)[0], categoryIds[0])
		return
	}
	
	if (*(*requestModel).CategoryIds)[1] != categoryIds[1] {
		t.Errorf("UserId convertion to DynamoDB attribute not correct, got = %v, want = %v", (*(*requestModel).CategoryIds)[1], categoryIds[1])
		return
	}
		
	if (*(*requestModel).CategoryIds)[2] != categoryIds[2] {
		t.Errorf("UserId convertion to DynamoDB attribute not correct, got = %v, want = %v", (*(*requestModel).CategoryIds)[2], categoryIds[2])
		return
	}

}

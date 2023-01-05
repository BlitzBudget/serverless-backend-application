package repository

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttributeBuilder(t *testing.T) {
	assert := assert.New(t)

	pk := "Wallet#2022-05-12T20:25:19Z"
	creationDate := "2022-05-12T20:25:19Z"
	updatedDate := "2022-05-12T20:25:19Z"
	targetDate := "2022-05-12T20:25:19Z"
	targetAmount := 2000
	goalName := "goalName"
	currentAmount := 2000
	goalAchieved := true
	body := "{\"pk\":\"" + pk + "\",\"creation_date\":\"" + creationDate + "\",\"target_amount\":" + strconv.Itoa(targetAmount) + ",\"current_amount\":" + strconv.Itoa(currentAmount) + ",\"updated_date\":\"" + updatedDate + "\",\"goal_name\":\"" + goalName + "\",\"target_date\":\"" + targetDate + "\",\"goal_achieved\":" + strconv.FormatBool(goalAchieved) + "}"

	got, err := AttributeBuilder(&body)
	assert.NoError(err)
	assert.NotNil(got)

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

	if *(*got["sk"]).S == "" {
		t.Errorf("name sk to DynamoDB attribute not correct, got = %v", *(*got["sk"]).S)
		return
	}

	if *(*got["target_amount"]).N != strconv.Itoa(targetAmount) {
		t.Errorf("name target_amount to DynamoDB attribute not correct, got = %v, want = %v", *(*got["target_amount"]).N, targetAmount)
		return
	}

	if *(*got["current_amount"]).N != strconv.Itoa(currentAmount) {
		t.Errorf("name current_amount to DynamoDB attribute not correct, got = %v, want = %v", *(*got["current_amount"]).N, currentAmount)
		return
	}

	if *(*got["goal_name"]).S != goalName {
		t.Errorf("name goal_name to DynamoDB attribute not correct, got = %v, want = %v", *(*got["goal_name"]).S, goalName)
		return
	}

	if *(*got["target_date"]).S != targetDate {
		t.Errorf("name target_date to DynamoDB attribute not correct, got = %v, want = %v", *(*got["target_date"]).S, targetDate)
		return
	}

	if *(*got["goal_achieved"]).BOOL != goalAchieved {
		t.Errorf("name goal_achieved to DynamoDB attribute not correct, got = %v, want = %v", *(*got["goal_achieved"]).BOOL, goalAchieved)
		return
	}
}

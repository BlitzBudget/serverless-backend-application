package repository

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttributeBuilder(t *testing.T) {
	assert := assert.New(t)

	pk := "pk"
	sk := "sk"
	targetAmount := 1000
	goalAchieved := true
	goalName := "goalName"
	currentAmount := 2000
	targetDate := "2022-09-15"
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"target_amount\":" + strconv.Itoa(targetAmount) + ",\"goal_achieved\":" + strconv.FormatBool(goalAchieved) + ",\"goal_name\":\"" + goalName + "\",\"current_amount\":" + strconv.Itoa(currentAmount) + ",\"target_date\":\"" + targetDate + "\"}"

	requestModel, err := AttributeBuilder(&body)
	assert.NoError(err)
	assert.NotNil(requestModel)

	if *requestModel.Pk != pk {
		t.Errorf("pk convertion to DynamoDB attribute not correct, got = %v, want = %v \n", *requestModel.Pk, pk)
		return
	}

	if *requestModel.Sk != sk {
		t.Errorf("SK convertion to DynamoDB attribute not correct, got = %v, want = %v \n", *requestModel.Sk, sk)
		return
	}

	if *requestModel.TargetAmount != float64(targetAmount) {
		t.Errorf("TargetAmount convertion to DynamoDB attribute not correct, got = %v, want = %v \n", *requestModel.TargetAmount, targetAmount)
		return
	}

	if *requestModel.CurrentAmount != float64(currentAmount) {
		t.Errorf("CurrentAmount convertion to DynamoDB attribute not correct, got = %v, want = %v \n", *requestModel.CurrentAmount, currentAmount)
		return
	}

	if *requestModel.GoalAchieved != goalAchieved {
		t.Errorf("GoalAchieved convertion to DynamoDB attribute not correct, got = %v, want = %v \n", *requestModel.GoalAchieved, goalAchieved)
		return
	}

	if *requestModel.Name != goalName {
		t.Errorf("Name convertion to DynamoDB attribute not correct, got = %v, want = %v \n", *requestModel.Name, goalName)
		return
	}

	if *requestModel.TargetDate != targetDate {
		t.Errorf("TargetDate convertion to DynamoDB attribute not correct, got = %v, want = %v \n", *requestModel.TargetDate, targetDate)
		return
	}
}
